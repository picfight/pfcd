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
		2:                                   400000172962501,
		4:                                   400000518887391,
		40:                                  400006745509838,
		400:                                 400069009068787,
		4000:                                400691378106416,
		40000:                               406888413296876,
		400000:                              466193246619086,
		1200000:                             580630465698828,
		2400000:                             707412244109675,
		3600000:                             780345162270002,
		4000000:                             792689721601108,
		4625285:                             799999997687360,
		calc.NumberOfGeneratingBlocks() + 5: 799999997687360,
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
