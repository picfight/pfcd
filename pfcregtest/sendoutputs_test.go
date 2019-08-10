// Copyright (c) 2018 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package pfcregtest

import (
	"github.com/jfixby/coinharness"
	"github.com/picfight/pfcd/rpcclient"
	"testing"

	"github.com/picfight/pfcd/chaincfg/chainhash"
	"github.com/picfight/pfcd/txscript"
	"github.com/picfight/pfcd/wire"
	"github.com/picfight/pfcutil"
)

func TestBallance(t *testing.T) {
	// Skip tests when running with -short
	if testing.Short() {
		t.Skip("Skipping RPC harness tests in short mode")
	}
	r := ObtainHarness(t.Name() + ".8")

	expectedBalance := pfcutil.Amount(
		int64(testSetup.Regnet25.NumMatureOutputs) * 50 * pfcutil.SatoshiPerPicfightcoin)
	actualBalance := r.Wallet.ConfirmedBalance()

	if actualBalance != expectedBalance {
		t.Fatalf("expected wallet balance of %v instead have %v",
			expectedBalance, actualBalance)
	}
}

func TestSendOutputs(t *testing.T) {
	// Skip tests when running with -short
	if testing.Short() {
		t.Skip("Skipping RPC harness tests in short mode")
	}
	r := ObtainHarness(t.Name() + ".7")

	genSpend := func(amt pfcutil.Amount) *chainhash.Hash {
		// Grab a fresh address from the wallet.
		addr, err := r.Wallet.NewAddress(nil)
		if err != nil {
			t.Fatalf("unable to get new address: %v", err)
		}

		// Next, send amt PFC to this address, spending from one of our mature
		// coinbase outputs.
		addrScript, err := txscript.PayToAddrScript(addr.(pfcutil.Address))
		if err != nil {
			t.Fatalf("unable to generate pkscript to addr: %v", err)
		}
		output := wire.NewTxOut(int64(amt), addrScript)
		sendArgs := &coinharness.SendOutputsArgs{
			Outputs: []coinharness.OutputTx{output},
			FeeRate: 10,
		}
		txid, err := r.Wallet.SendOutputs(sendArgs)
		if err != nil {
			t.Fatalf("coinbase spend failed: %v", err)
		}
		return txid.(*chainhash.Hash)
	}

	assertTxMined := func(txid *chainhash.Hash, blockHash *chainhash.Hash) {
		block, err := r.NodeRPCClient().(*rpcclient.Client).GetBlock(blockHash)
		if err != nil {
			t.Fatalf("unable to get block: %v", err)
		}

		numBlockTxns := len(block.Transactions)
		if numBlockTxns < 2 {
			t.Fatalf("crafted transaction wasn't mined, block should have "+
				"at least %v transactions instead has %v", 2, numBlockTxns)
		}

		minedTx := block.Transactions[1]
		txHash := minedTx.TxHash()
		if txHash != *txid {
			t.Fatalf("txid's don't match, %v vs %v", txHash, txid)
		}
	}

	// First, generate a small spend which will require only a single
	// input.
	txid := genSpend(pfcutil.Amount(5 * pfcutil.SatoshiPerPicfightcoin))

	// Generate a single block, the transaction the wallet created should
	// be found in this block.
	blockHashes, err := r.NodeRPCClient().(*rpcclient.Client).Generate(1)
	if err != nil {
		t.Fatalf("unable to generate single block: %v", err)
	}
	assertTxMined(txid, blockHashes[0])

	// Next, generate a spend much greater than the block reward. This
	// transaction should also have been mined properly.
	txid = genSpend(pfcutil.Amount(500 * pfcutil.SatoshiPerPicfightcoin))
	blockHashes, err = r.NodeRPCClient().(*rpcclient.Client).Generate(1)
	if err != nil {
		t.Fatalf("unable to generate single block: %v", err)
	}
	assertTxMined(txid, blockHashes[0])
}
