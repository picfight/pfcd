package chaincfg

import (
	"fmt"
	"github.com/jfixby/coin"
	"github.com/picfight/picfightcoin"
	"testing"
)

func TestLockPremine(t *testing.T) {
	{ // LOCK PROJECT SUBSIDY
		projectPremine := picfightcoin.Premine()["JsKFRL5ivSH7CnYaTtaBT4M9fZG878g49Fg"]
		expected := coin.FromFloat(20291)
		if projectPremine.AtomsValue != expected.AtomsValue {
			t.Errorf("Premine mismatch: got %v expected %v ",
				projectPremine,
				expected,
			)
			t.Fail()
		}
	}
	{ // LOCK PROJECT POS SUBSIDY
		projectPosPremine := picfightcoin.Premine()["JsRjbYZ448FxZQ5kQAc15NcwUQ1oqYydVEG"]
		expected := coin.FromFloat(60)
		if projectPosPremine.AtomsValue != expected.AtomsValue {
			t.Errorf("Premine mismatch: got %v expected %v ",
				projectPosPremine,
				expected,
			)
			t.Fail()
		}
	}
}

func TestPremine(t *testing.T) {
	premine := BlockOneLedgerPicfightCoin()
	totalAtoms := int64(0)
	expected := picfightcoin.PremineTotal.AtomsValue
	for _, e := range premine {
		totalAtoms = totalAtoms + e.Amount
		fmt.Println(fmt.Sprintf("premine: %v -> %v",
			e.Address,
			e.Amount,
		),
		)
	}
	if totalAtoms != expected {
		t.Errorf("Premine mismatch: got %v expected %v ",
			totalAtoms,
			expected,
		)
		t.Fail()
	}
	expected = coin.FromFloat(20351).AtomsValue
	if totalAtoms != expected {
		t.Errorf("Premine mismatch: got %v expected %v ",
			totalAtoms,
			expected,
		)
		t.Fail()
	}

}

func TestDecredBlock1(t *testing.T) {
	params := DecredNetParams

	block1_subsidy := params.SubsidyCalculator().CalcBlockSubsidy(1)

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

func TestPicFightCoinBlock1(t *testing.T) {
	params := PicFightCoinNetParams

	block1_subsidy := params.SubsidyCalculator().CalcBlockSubsidy(1)

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
