package chaincfg

import (
	"github.com/picfight/picfightcoin"
	"testing"
)

func TestPremine(t *testing.T) {
	premine := BlockOneLedgerPicfightCoin()
	totalAtoms := int64(0)
	expected := picfightcoin.PremineTotal.AtomsValue
	for _, e := range premine {
		totalAtoms = totalAtoms + e.Amount
		if totalAtoms != expected {
			t.Errorf("Premine mismatch: got %v expected %v ",
				totalAtoms,
				expected,
			)
			t.Fail()
		}
	}
}
func TestBlock1(t *testing.T) {
	//-----------
	params := PicFightCoinNetParams

	block1_subsidy := params.SubsidyCalculator.CalcBlockSubsidy(1)

	// Block height 1 subsidy is 'special' and used to
	// distribute initial tokens, if any.
	block1_spremined := params.BlockOneSubsidy()

	if block1_subsidy != block1_spremined {
		t.Errorf("Premine mismatch: SubsidyCalculator(1) %v BlockOneSubsidy() %v ",
			block1_subsidy,
			block1_spremined,
		)
		t.Fail()
	}

}
