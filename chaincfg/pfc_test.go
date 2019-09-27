package chaincfg

import (
	"testing"
)

func TestPremine(t *testing.T) {
	premine := BlockOneLedgerPicfightCoin()
	totalAtoms := int64(0)
	expected := int64(0)
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
