package pfcregtest

import (
	"fmt"
	"github.com/jfixby/coinharness"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/commandline"
	"github.com/jfixby/pin/gobuilder"
	"github.com/picfight/pfcharness"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/picfight/pfcd/chaincfg"
)

// Default harness name
const mainHarnessName = "main"

// SimpleTestSetup harbours:
type SimpleTestSetup struct {
	PicFightCoinNetA *ChainWithMatureOutputsSpawner
	PicFightCoinNetB *ChainWithMatureOutputsSpawner

	harnessPool *pin.Pool

	// WorkingDir defines test setup working dir
	WorkingDir *pin.TempDirHandler
}

// TearDown all harnesses in test Pool.
// This includes removing all temporary directories,
// and shutting down any created processes.
func (setup *SimpleTestSetup) TearDown() {
	setup.harnessPool.DisposeAll()
	setup.WorkingDir.Dispose()
}

// Setup deploys this test setup
func Setup() *SimpleTestSetup {
	setup := &SimpleTestSetup{
		WorkingDir: pin.NewTempDir(setupWorkingDir(), "pfc-testing").MakeDir(),
	}

	memWalletFactory := &pfcharness.InMemoryWalletFactory{}

	//wEXE := &commandline.ExplicitExecutablePathString{
	//	PathString: "pfcwallet",
	//}
	//consoleWalletFactory := &pfcharness.ConsoleWalletFactory{
	//	WalletExecutablePathProvider: wEXE,
	//}

	walletFactory := memWalletFactory

	nodeFactoryA := &pfcharness.ConsoleNodeFactory{
		NodeExecutablePathProvider: &commandline.ExplicitExecutablePathString{
			PathString: "pfcd_25",
		},
	}

	nodeFactoryB := &pfcharness.ConsoleNodeFactory{
		NodeExecutablePathProvider: &commandline.ExplicitExecutablePathString{
			PathString: "pfcd_drop",
		},
	}

	portManager := &coinharness.LazyPortManager{
		BasePort: 30000,
	}

	testSeed := pfcharness.NewTestSeed

	setup.PicFightCoinNetA = &ChainWithMatureOutputsSpawner{
		WorkingDir:        setup.WorkingDir.Path(),
		DebugNodeOutput:   true,
		DebugWalletOutput: true,
		NetPortManager:    portManager,
		WalletFactory:     walletFactory,
		NodeFactory:       nodeFactoryA,
		ActiveNet:         &pfcharness.Network{&chaincfg.PicFightCoinNetParams},
		CreateTempWallet:  true,
		NewTestSeed:       testSeed,
	}

	setup.PicFightCoinNetB = &ChainWithMatureOutputsSpawner{
		WorkingDir:        setup.WorkingDir.Path(),
		DebugNodeOutput:   true,
		DebugWalletOutput: true,
		NetPortManager:    portManager,
		WalletFactory:     walletFactory,
		NodeFactory:       nodeFactoryB,
		ActiveNet:         &pfcharness.Network{&chaincfg.PicFightCoinNetParams},
		CreateTempWallet:  true,
		NewTestSeed:       testSeed,
	}

	setup.harnessPool = pin.NewPool(setup.PicFightCoinNetA)

	return setup
}

func setupWorkingDir() string {
	testWorkingDir, err := ioutil.TempDir("", "integrationtest")
	if err != nil {
		fmt.Println("Unable to create working dir: ", err)
		os.Exit(-1)
	}
	return testWorkingDir
}

func setupBuild(buildName string, workingDir string, nodeProjectGoPath string) *gobuilder.GoBuider {

	tempBinDir := filepath.Join(workingDir, "bin")
	pin.MakeDirs(tempBinDir)

	nodeGoBuilder := &gobuilder.GoBuider{
		GoProjectPath:    nodeProjectGoPath,
		OutputFolderPath: tempBinDir,
		BuildFileName:    buildName,
	}
	return nodeGoBuilder
}
