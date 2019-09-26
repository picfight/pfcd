package pfcregtest

import (
	"fmt"
	"github.com/jfixby/coinharness"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/commandline"
	"github.com/jfixby/pin/fileops"
	"path/filepath"
	"testing"
)

func TestCertCreationWithHosts(t *testing.T) {

	fileops.EngageDeleteSafeLock(true)
	fileops.CopyFolderContentToFolder(
		filepath.Join("1000-blocks", "nodeA"),
		filepath.Join("temp", "harness-A", "node"),
		func(string) bool { return true },
		fileops.ALL_CHILDREN,
	)

	var harnessA *coinharness.Harness
	var harnessB *coinharness.Harness
	defer Dispose(harnessA)
	defer Dispose(harnessB)
	{
		ts := testSetup.PicFightCoinNetA
		ts.WorkingDir = fileops.Abs(filepath.Join("temp"))
		harnessA = NewInstance(ts, "A").(*coinharness.Harness)
	}
	{
		ts := testSetup.PicFightCoinNetB
		ts.WorkingDir = fileops.Abs(filepath.Join("temp"))
		harnessB = NewInstance(ts, "B").(*coinharness.Harness)
	}
	commits := map[string]interface{}{
		"generate": commandline.NoArgumentValue,
		"addpeer":  harnessB.Node.P2PAddress(),
	}

	sargs := &coinharness.StartNodeArgs{
		DebugOutput:    true,
		MiningAddress:  harnessA.MiningAddress,
		ExtraArguments: commits,
	}

	harnessA.Node.Stop()
	harnessA.Node.Start(sargs)
	for {
		n, err := harnessA.NodeRPCClient().GetBlockCount()
		pin.CheckTestSetupMalfunction(err)
		if n > 2000 {
			break
		}
	}

	fileops.Delete("temp")

}

// NewInstance does the following:
//   1. Starts a new NodeTestServer process with a fresh SimNet chain.
//   2. Creates a new temporary WalletTestServer connected to the running NodeTestServer.
//   3. Gets a new address from the WalletTestServer for mining subsidy.
//   4. Restarts the NodeTestServer with the new mining address.
//   5. Generates a number of blocks so that testing starts with a spendable
//      balance.
func NewInstance(testSetup *ChainWithMatureOutputsSpawner, harnessName string) pin.Spawnable {
	harnessFolderName := "harness-" + harnessName
	pin.AssertNotNil("ConsoleNodeFactory", testSetup.NodeFactory)
	pin.AssertNotNil("ActiveNet", testSetup.ActiveNet)
	pin.AssertNotNil("WalletFactory", testSetup.WalletFactory)

	// This allows to specify custom walled seed salt by adding the dot
	// in the harness name.
	// Example: "harness.65" will create harness wallet seed equal to the 65
	seedSalt := coinharness.ExtractSeedSaltFromHarnessName(harnessName)

	harnessFolder := filepath.Join(testSetup.WorkingDir, harnessFolderName)
	walletFolder := filepath.Join(harnessFolder, "wallet")
	nodeFolder := filepath.Join(harnessFolder, "node")

	p2p := testSetup.NetPortManager.ObtainPort()
	nodeRPC := testSetup.NetPortManager.ObtainPort()
	walletRPC := testSetup.NetPortManager.ObtainPort()

	localhost := "127.0.0.1"

	nodeConfig := &coinharness.TestNodeConfig{
		P2PHost: localhost,
		P2PPort: p2p,

		NodeRPCHost: localhost,
		NodeRPCPort: nodeRPC,

		NodeUser:     "node.user",
		NodePassword: "node.pass",

		ActiveNet: testSetup.ActiveNet,

		WorkingDir: nodeFolder,
	}

	walletConfig := &coinharness.TestWalletConfig{
		Seed:        testSetup.NewTestSeed(seedSalt),
		NodeRPCHost: localhost,
		NodeRPCPort: nodeRPC,

		WalletRPCHost: localhost,
		WalletRPCPort: walletRPC,

		NodeUser:     "node.user",
		NodePassword: "node.pass",

		WalletUser:     "wallet.user",
		WalletPassword: "wallet.pass",

		ActiveNet:  testSetup.ActiveNet,
		WorkingDir: walletFolder,
	}

	harness := &coinharness.Harness{
		Name:       harnessName,
		Node:       testSetup.NodeFactory.NewNode(nodeConfig),
		Wallet:     testSetup.WalletFactory.NewWallet(walletConfig),
		WorkingDir: harnessFolder,
	}

	pin.AssertTrue("Networks match", harness.Node.Network() == harness.Wallet.Network())

	nodeNet := harness.Node.Network()
	walletNet := harness.Wallet.Network()

	pin.AssertTrue(
		fmt.Sprintf(
			"Wallet Net<%v> is the same as Node Net<%v>", walletNet, nodeNet),
		walletNet == nodeNet)

	DeploySimpleChain(testSetup, harness)

	return harness
}

// Dispose harness. This includes removing
// all temporary directories, and shutting down any created processes.
func Dispose(s pin.Spawnable) error {
	h := s.(*coinharness.Harness)
	if h == nil {
		return nil
	}
	h.Wallet.Dispose()
	h.Node.Dispose()
	return h.DeleteWorkingDir()
}
