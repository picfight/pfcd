package blockchain

import (
	"fmt"
	"github.com/jfixby/pin"
	"github.com/picfight/pfcd/chaincfg"
	"testing"
)

func TestEstimateSupplyRefactored(t *testing.T) {
	net := &chaincfg.DecredNetParams
	indexes := []int64{0, 1, 2, 3, 4, 5, 6, 7, 200, 300, 400, 500, 600, 1000, 10000, 100000, 1000000, 10000000, 1000000000, 100000000000}
	for _, e := range indexes {
		oldValue := estimateSupply(net, e)
		newValue := estimateSupplyV2(net, e)
		if oldValue != newValue {
			t.Errorf("EstimateSupply mismatch at block[%v]: %v != %v",
				e,
				oldValue,
				newValue,
			)
		}

		pin.D(fmt.Sprintf("[%v]", e), oldValue)
	}

}
