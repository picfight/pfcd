package pfcregtest

import (
	"fmt"
	"github.com/jfixby/coinharness"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/commandline"
	"path/filepath"
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

	NodeFactory   coinharness.TestNodeFactory
	WalletFactory coinharness.TestWalletFactory

	ActiveNet coinharness.Network

	NewTestSeed func(u uint32) coinharness.Seed

	NetPortManager coinharness.NetPortManager

	NodeStartExtraArguments   map[string]interface{}
	WalletStartExtraArguments map[string]interface{}
	CreateTempWallet          bool
}

// NewInstance does the following:
//   1. Starts a new NodeTestServer process with a fresh SimNet chain.
//   2. Creates a new temporary WalletTestServer connected to the running NodeTestServer.
//   3. Gets a new address from the WalletTestServer for mining subsidy.
//   4. Restarts the NodeTestServer with the new mining address.
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

// DeploySimpleChain defines harness setup sequence for this package:
// 1. obtains a new mining wallet address
// 2. restart harness node and wallet with the new mining address
// 3. builds a new chain with the target number of mature outputs
// receiving the mining reward to the test wallet
// 4. syncs wallet to the tip of the chain
func DeploySimpleChain(testSetup *ChainWithMatureOutputsSpawner, h *coinharness.Harness) {
	pin.AssertNotEmpty("harness name", h.Name)
	fmt.Println("Deploying Harness[" + h.Name + "]")
	createFlag := testSetup.CreateTempWallet
	// launch a fresh h (assumes h working dir is empty)
	{
		args := &launchArguments{
			DebugNodeOutput:    testSetup.DebugNodeOutput,
			DebugWalletOutput:  testSetup.DebugWalletOutput,
			NodeExtraArguments: testSetup.NodeStartExtraArguments,
		}
		if createFlag {
			args.WalletExtraArguments = make(map[string]interface{})
			args.WalletExtraArguments["createtemp"] = commandline.NoArgumentValue
		}
		launchHarnessSequence(h, args)
	}

	// Get a new address from the WalletTestServer
	// to be set with node --miningaddr
	var address coinharness.Address
	var err error
	{
		for {
			address, err = h.Wallet.NewAddress(coinharness.DefaultAccountName)
			if err != nil {
				pin.D("address", address)
				pin.D("error", err)
				pin.Sleep(1000)
			} else {
				break
			}
		}

		//pin.CheckTestSetupMalfunction(err)
		h.MiningAddress = address

		pin.AssertNotNil("MiningAddress", h.MiningAddress)
		pin.AssertNotEmpty("MiningAddress", h.MiningAddress.String())

		fmt.Println("Mining address: " + h.MiningAddress.String())
	}

	// restart the h with the new argument
	{
		shutdownHarnessSequence(h)

		args := &launchArguments{
			DebugNodeOutput:    testSetup.DebugNodeOutput,
			DebugWalletOutput:  testSetup.DebugWalletOutput,
			NodeExtraArguments: testSetup.NodeStartExtraArguments,
		}
		if createFlag {
			args.WalletExtraArguments = make(map[string]interface{})
			args.WalletExtraArguments["createtemp"] = commandline.NoArgumentValue
		}
		launchHarnessSequence(h, args)
	}

	fmt.Println("Harness[" + h.Name + "] is ready")
}

// local struct to bundle launchHarnessSequence function arguments
type launchArguments struct {
	DebugNodeOutput      bool
	DebugWalletOutput    bool
	MiningAddress        coinharness.Address
	NodeExtraArguments   map[string]interface{}
	WalletExtraArguments map[string]interface{}
}

// launchHarnessSequence
func launchHarnessSequence(h *coinharness.Harness, args *launchArguments) {
	node := h.Node
	wallet := h.Wallet

	sargs := &coinharness.StartNodeArgs{
		DebugOutput:    args.DebugNodeOutput,
		MiningAddress:  h.MiningAddress,
		ExtraArguments: args.NodeExtraArguments,
	}
	node.Start(sargs)

	rpcConfig := node.RPCConnectionConfig()

	walletLaunchArguments := &coinharness.TestWalletStartArgs{
		NodeRPCCertFile:          node.CertFile(),
		DebugOutput:              args.DebugWalletOutput,
		MaxSecondsToWaitOnLaunch: 90,
		NodeRPCConfig:            rpcConfig,
		ExtraArguments:           args.WalletExtraArguments,
	}

	// wait for the WalletTestServer to sync up to the current height
	_, _, e := h.NodeRPCClient().GetBestBlock()
	pin.CheckTestSetupMalfunction(e)

	wallet.Start(walletLaunchArguments)

}

// shutdownHarnessSequence reverses the launchHarnessSequence
func shutdownHarnessSequence(harness *coinharness.Harness) {
	harness.Wallet.Stop()
	harness.Node.Stop()
}
