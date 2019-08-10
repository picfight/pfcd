package blockchain

import (
	"github.com/picfight/pfcd/chaincfg"
	"testing"
	"time"
)

// TestTotalSubsidy ensures the total subsidy produced matches the expected
// value.
func TestTotalSubsidy(t *testing.T) {
	// Locals for convenience.
	netParams := &chaincfg.MainNetParams

	//YEARS := 132
	minPblock := int(netParams.TargetTimePerBlock / time.Minute)
	//N := int32(60 / minPblock * 24 * 365 * YEARS)
	N := int32(60 / minPblock)

	// Calculate the total possible subsidy.
	totalSubsidy := int64(0)
	for blockNum := int32(1); blockNum <= N; blockNum++ {
		totalSubsidy += CalcBlockSubsidy(blockNum, netParams)
	}

	// Ensure the total calculated subsidy is the expected value.
	const expectedTotalSubsidy = int64(60 * 1e8)

	if totalSubsidy != expectedTotalSubsidy {
		t.Fatalf("mismatched total subsidy -- got %d, want %d", totalSubsidy,
			expectedTotalSubsidy)
	}
}
