package blockchain

import (
	"fmt"
	"github.com/jfixby/bignum"
	"github.com/picfight/pfcd/chaincfg"
	"testing"
)

// TestTotalSubsidyFloat64 ensures the total subsidy produced matches the expected
// value.
func TestTotalSubsidyFloat64(t *testing.T) {
	netParams := &chaincfg.MainNetParams
	testChainSubsidy(t, bignum.Float64Engine{}, netParams)
}

// TestTotalSubsidyCalcAlgo ensures the calcSubsidy() function works correctly.
func TestTotalSubsidyCalcAlgo(t *testing.T) {
	t.Skip()
	netParams := &chaincfg.MainNetParams
	// we need the exact sum value, so use BigIntEngine
	testChainSubsidy(t, bignum.BigIntEngine{}, netParams)
}

func testChainSubsidy(t *testing.T, engine bignum.BigNumEngine, netParams *chaincfg.Params) {
	N := int32(netParams.SubsidyProductionPeriod / netParams.TargetTimePerBlock)
	subsidyBlocksNumber := int64(N)
	targetTotalSubsidy := float64(netParams.TargetTotalSubsidy)
	satoshiPerCoin := engine.NewBigNum(chaincfg.SatoshiPerPicfightcoin)

	result := testCalcSubsidy(t, engine, subsidyBlocksNumber, targetTotalSubsidy, 60*24*365)

	resultSatoshi := result.Mul(result, satoshiPerCoin).ToInt64()
	expectedSatoshi := netParams.TargetTotalSubsidy * chaincfg.SatoshiPerPicfightcoin
	if resultSatoshi != expectedSatoshi {
		floatComputedExpectedSatoshi := int64(expectedSatoshi - 1)
		if resultSatoshi != floatComputedExpectedSatoshi {
			t.Fatalf("mismatched total satoshi subsidy -- \n got %v, \nwant %v", expectedSatoshi,
				resultSatoshi)
		}
	}
}

func testCalcSubsidy(t *testing.T, engine bignum.BigNumEngine, subsidyBlocksNumber int64, targetTotalSubsidy float64, printIterations int64) bignum.BigNum {
	testHeight := subsidyBlocksNumber
	totalSubsidy := engine.NewBigNum(0)
	for blockNum := int64(0); blockNum <= testHeight; blockNum++ {
		sub := calcSubsidy(engine, subsidyBlocksNumber, blockNum, targetTotalSubsidy)
		//totalSubsidy += sub
		totalSubsidy = totalSubsidy.Add(totalSubsidy, sub)
		if blockNum%printIterations == 0 {
			blockNumPad := fmt.Sprintf("%15v", blockNum)
			subPad := fmt.Sprintf("%-30v", sub.ToFloat64())
			totalSubsidyPad := fmt.Sprintf("%-30v", totalSubsidy.ToFloat64())
			t.Log(fmt.Sprintf("[%v] %v coins %v total", blockNumPad, subPad, totalSubsidyPad))
		}
	}
	t.Log(fmt.Sprintf("totalSubsidy: %16v", totalSubsidy.ToFloat64()))
	return totalSubsidy
}

func TestFloatEngine(t *testing.T) {
	t.Parallel()
	subsidyBlocksNumber := int64(3)
	targetTotalSubsidy := float64(1)
	resultFloat64 := testCalcSubsidy(t, bignum.Float64Engine{}, subsidyBlocksNumber, targetTotalSubsidy, 1).ToFloat64()
	resultBigFloat := testCalcSubsidy(t, bignum.BigIntEngine{}, subsidyBlocksNumber, targetTotalSubsidy, 1).ToFloat64()

	if resultFloat64 != (resultBigFloat) {
		t.Fatalf("mismatched total subsidy -- \n got %v, \nwant %v", resultFloat64, resultBigFloat)
	}
}
