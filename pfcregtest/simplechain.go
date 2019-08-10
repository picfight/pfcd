// Copyright (c) 2018 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package pfcregtest

import (
	"fmt"
	"github.com/jfixby/coinharness"
	"github.com/jfixby/pin"
	"github.com/picfight/pfcharness"
	"path/filepath"

	"github.com/picfight/pfcd/chaincfg"
)

// ChainWithMatureOutputsSpawner initializes the primary mining node
// with a test chain of desired height, providing numMatureOutputs coinbases
// to allow spending from for testing purposes.
type ChainWithMatureOutputsSpawner struct {
	// Each harness will be provided with a dedicated
	// folder inside the WorkingDir
	WorkingDir string

	// DebugNodeOutput, set true to print out node output to console
	DebugNodeOutput bool

	// DebugWalletOutput, set true to print out wallet output to console
	DebugWalletOutput bool

	// NumMatureOutputs sets requirement for the generated test chain
	NumMatureOutputs uint32

	NodeFactory   coinharness.TestNodeFactory
	WalletFactory coinharness.TestWalletFactory

	ActiveNet *chaincfg.Params

	NetPortManager coinharness.NetPortManager

	NodeStartExtraArguments   map[string]interface{}
	WalletStartExtraArguments map[string]interface{}
}

// NewInstance does the following:
//   1. Starts a new DcrdTestServer process with a fresh SimNet chain.
//   2. Creates a new temporary WalletTestServer connected to the running DcrdTestServer.
//   3. Gets a new address from the WalletTestServer for mining subsidy.
//   4. Restarts the DcrdTestServer with the new mining address.
//   5. Generates a number of blocks so that testing starts with a spendable
//      balance.
func (testSetup *ChainWithMatureOutputsSpawner) NewInstance(harnessName string) pin.Spawnable {
	harnessFolderName := "harness-" + harnessName
	pin.AssertNotNil("ConsoleNodeFactory", testSetup.NodeFactory)
	pin.AssertNotNil("ActiveNet", testSetup.ActiveNet)
	pin.AssertNotNil("WalletFactory", testSetup.WalletFactory)

	// This allows to specify custom walled seed salt by adding the dot
	// in the harness name.
	// Example: "harness.65" will create harness wallet seed equal to the 65
	seedSalt := extractSeedSaltFromHarnessName(harnessName)

	harnessFolder := filepath.Join(testSetup.WorkingDir, harnessFolderName)

	p2p := testSetup.NetPortManager.ObtainPort()
	nodeRPC := testSetup.NetPortManager.ObtainPort()
	walletRPC := testSetup.NetPortManager.ObtainPort()

	localhost := "127.0.0.1"

	nodeConfig := &coinharness.TestNodeConfig{
		P2PHost: localhost,
		P2PPort: p2p,

		NodeRPCHost: localhost,
		NodeRPCPort: nodeRPC,

		ActiveNet: testSetup.ActiveNet,

		WorkingDir: harnessFolder,
	}

	walletConfig := &coinharness.TestWalletConfig{
		Seed:          pfcharness.NewTestSeed(seedSalt),
		WalletRPCHost: localhost,
		WalletRPCPort: walletRPC,
		ActiveNet:     testSetup.ActiveNet,
	}

	harness := &coinharness.Harness{
		Name:       harnessName,
		Node:       testSetup.NodeFactory.NewNode(nodeConfig),
		Wallet:     testSetup.WalletFactory.NewWallet(walletConfig),
		WorkingDir: harnessFolder,
	}

	nodeNet := harness.Node.Network()
	walletNet := harness.Wallet.Network()

	pin.AssertTrue(
		fmt.Sprintf(
			"Wallet net<%v> is the same as Node net<%v>", walletNet, nodeNet),
		walletNet == nodeNet)

	DeploySimpleChain(testSetup, harness)

	return harness
}

// Dispose harness. This includes removing
// all temporary directories, and shutting down any created processes.
func (testSetup *ChainWithMatureOutputsSpawner) Dispose(s pin.Spawnable) error {
	h := s.(*coinharness.Harness)
	if h == nil {
		return nil
	}
	h.Wallet.Dispose()
	h.Node.Dispose()
	return h.DeleteWorkingDir()
}

// NameForTag defines policy for mapping input tags to harness names
func (testSetup *ChainWithMatureOutputsSpawner) NameForTag(tag string) string {
	harnessName := tag
	return harnessName
}
