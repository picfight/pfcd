package pfcregtest

import (
	"github.com/picfight/pfcd/rpcclient"
	"testing"
)

func TestBuildVerstion(t *testing.T) {
	//if testing.Short() {
	//	t.Skip("Skipping RPC harness tests in short mode")
	//}
	r := ObtainHarness(mainHarnessName)
	// Create a new block connecting to the current tip.
	versionString, err := r.NodeRPCClient().(*rpcclient.Client).GetBuildVersion()
	if err != nil {
		t.Fatalf("Unable to get build vesion: %v", err)
	}
	EXPECTED := "build-00001"
	if *versionString != EXPECTED {
		t.Fatalf("Wrong build vesion: <%v>, expected <%v>", versionString, EXPECTED)
	}

}
