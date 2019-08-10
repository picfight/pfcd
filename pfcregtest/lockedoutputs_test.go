// Copyright (c) 2018 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package pfcregtest

import (
	"github.com/jfixby/coinharness"
	"github.com/picfight/pfcutil"
	"testing"

	"github.com/picfight/pfcd/txscript"
	"github.com/picfight/pfcd/wire"
)

func TestMemWalletLockedOutputs(t *testing.T) {
	// Skip tests when running with -short
	//if testing.Short() {
	//	t.Skip("Skipping RPC h tests in short mode")
	//}
	r := ObtainHarness(mainHarnessName)
	// Obtain the initial balance of the wallet at this point.
	startingBalance := r.Wallet.ConfirmedBalance().(pfcutil.Amount)

	// First, create a signed transaction spending some outputs.
	addr, err := r.Wallet.NewAddress(nil)
	if err != nil {
		t.Fatalf("unable to generate new address: %v", err)
	}
	pkScript, err := txscript.PayToAddrScript(addr.(pfcutil.Address))
	if err != nil {
		t.Fatalf("unable to create script: %v", err)
	}
	outputAmt := pfcutil.Amount(50 * pfcutil.SatoshiPerPicfightcoin)
	output := wire.NewTxOut(int64(outputAmt), pkScript)
	ctargs := &coinharness.CreateTransactionArgs{
		Outputs: []coinharness.OutputTx{output},
		FeeRate: 10,
		Change:  true,
	}
	tx, err := r.Wallet.CreateTransaction(ctargs)
	if err != nil {
		t.Fatalf("unable to create transaction: %v", err)
	}

	// The current wallet balance should now be at least 50 PFC less
	// (accounting for fees) than the period balance
	currentBalance := r.Wallet.ConfirmedBalance().(pfcutil.Amount)
	if !(currentBalance <= startingBalance-outputAmt) {
		t.Fatalf("spent outputs not locked: previous balance %v, "+
			"current balance %v", startingBalance, currentBalance)
	}

	// Now unlocked all the spent inputs within the unbroadcast signed
	// transaction. The current balance should now be exactly that of the
	// starting balance.
	txin := tx.(*wire.MsgTx).TxIn
	inpts := make([]coinharness.InputTx, len(txin))
	for i, j := range txin {
		inpts[i] = j
	}
	r.Wallet.UnlockOutputs(inpts)
	currentBalance = r.Wallet.ConfirmedBalance().(pfcutil.Amount)
	if currentBalance != startingBalance {
		t.Fatalf("current and starting balance should now match: "+
			"expected %v, got %v", startingBalance, currentBalance)
	}
}
