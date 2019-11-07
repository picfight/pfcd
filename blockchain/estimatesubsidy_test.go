package blockchain

import (
	"fmt"
	"github.com/jfixby/pin"
	"github.com/picfight/pfcd/chaincfg"
	"github.com/picfight/picfightcoin"
	"testing"
)

func TestEstimateSupplyRefactored(t *testing.T) {
	net := &chaincfg.DecredNetParams
	for height, _ := range decredExpectedSubsidyEstimateValues {
		oldValue := estimateSupplyV1(net, height)
		newValue := estimateSupplyV2(net, height)
		if oldValue != newValue {
			t.Errorf("EstimateSupply mismatch at block[%v]: %v != %v",
				height,
				oldValue,
				newValue,
			)
		}

		pin.D(fmt.Sprintf("[%v]", height), oldValue)
	}

}

func TestEstimateSupplyPicFightCoin(t *testing.T) {
	calc := picfightcoin.PicFightCoinSubsidy()
	expectedValues := map[int64]int64{
		0:                                   0,
		1:                                   calc.PreminedCoins().AtomsValue,
		2:                                   2035109104832,
		4:                                   2035127314475,
		40:                                  2035455083102,
		400:                                 2038732254894,
		4000:                                2071452524241,
		40000:                               2393510359606,
		400000:                              5099602903191,
		1200000:                             7764076009766,
		2400000:                             7777699369282,
		3600000:                             7777699369282,
		4000000:                             7777699369282,
		4625285:                             7777699369282,
		calc.NumberOfGeneratingBlocks() + 5: 7777699369282,
	}

	for height, expected := range expectedValues {
		est := calc.EstimateSupply(height)
		if est != expected {
			t.Errorf("EstimateSupply mismatch at block[%v]: %v != %v",
				height,
				est,
				expected,
			)
			t.FailNow()
		}
		t.Logf("[%v] %v", height, est)
	}

}

var decredExpectedSubsidyEstimateValues = map[int64]int64{
	0:            0,
	1:            picfightcoin.DecredMainNetSubsidy().PreminedCoins().AtomsValue,
	2:            168003119582664,
	4:            168009358747992,
	40:           168121663723896,
	400:          169244713482936,
	4000:         180475211073336,
	40000:        289403792216190,
	400000:       1091015157627320,
	1200000:      1826592807729333,
	2400000:      2064127743332932,
	3600000:      2098145810684668,
	4000000:      2100856975988926,
	4625285:      2102751212506332,
	9250570:      2103831023303515,
	11003904:     2103831507350641,
	11003904 * 2: 2103831507356784,
}

func TestEstimateSupplyDecred(t *testing.T) {
	calc := picfightcoin.DecredMainNetSubsidy()
	for height, expected := range decredExpectedSubsidyEstimateValues {
		est := calc.EstimateSupply(height)
		if est != expected {
			t.Errorf("EstimateSupply mismatch at block[%v]: %v != %v",
				height,
				est,
				expected,
			)
			t.FailNow()
		}
		t.Logf("[%v] %v", height, est)
	}

}
