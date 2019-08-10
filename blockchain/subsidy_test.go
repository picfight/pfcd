package blockchain

import (
	"github.com/picfight/pfcd/chaincfg"
	"testing"
)

// TestTotalSubsidy ensures the total subsidy produced matches the expected
// value.
func TestTotalSubsidy(t *testing.T) {
	// Locals for convenience.
	netParams := &chaincfg.MainNetParams

	YEARS := 132
	N := int32(60 / 10 * 24 * 365 * YEARS)

	// Calculate the total possible subsidy.
	totalSubsidy := int64(0)
	for blockNum := int32(0); blockNum <= N; blockNum++ {
		totalSubsidy += CalcBlockSubsidy(blockNum, netParams)
	}

	// Ensure the total calculated subsidy is the expected value.
	const expectedTotalSubsidy = 2099999997690000
	if totalSubsidy != expectedTotalSubsidy {
		t.Fatalf("mismatched total subsidy -- got %d, want %d", totalSubsidy,
			expectedTotalSubsidy)
	}
}
