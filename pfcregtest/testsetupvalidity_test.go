package pfcregtest

import (
	"fmt"
	"github.com/picfight/pfcutil"
	"testing"
)

func TestSetupValidity(t *testing.T) {
	//pin.D("btcdEXE", btcdEXE)
	coins50 := pfcutil.Amount(50 /*PFC*/ * 1e8)
	//pin.D("pfcutil.Amount(coinbaseTx.Value)", coins50)
	stringVal := fmt.Sprintf("%v", coins50)
	expectedStringVal := "50 PFC"
	//pin.D("stringVal", stringVal)
	if expectedStringVal != stringVal {
		t.Fatalf("Incorrect coin: "+
			"expected %v, got %v", expectedStringVal, stringVal)
	}
}
