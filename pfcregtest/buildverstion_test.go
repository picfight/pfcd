package pfcregtest

import (
	"github.com/jfixby/coinharness"
	"github.com/picfight/pfcd/rpcclient"
	"testing"
)

func TestBuildVerstion(t *testing.T) {
	//if testing.Short() {
	//	t.Skip("Skipping RPC harness tests in short mode")
	//}
	pool := testSetup.Mainnet0
	r := pool.NewInstance(t.Name()).(*coinharness.Harness)
	defer pool.Dispose(r)
	// Create a new block connecting to the current tip.
	version, err := r.NodeRPCClient().(*rpcclient.Client).GetBuildVersion()
	if err != nil {
		t.Fatalf("Unable to get build vesion: %v", err)
	}
	EXPECTED := "build-00002.mainnet"
	if version.VersionString != EXPECTED {
		t.Fatalf("Wrong build vesion: <%v>, expected <%v>", version.VersionString, EXPECTED)
	}

}
