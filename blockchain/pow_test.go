package blockchain

import (
	"fmt"
	//"github.com/jfixby/pin"
	"github.com/picfight/pfcd/chaincfg"
	"math/big"
	"reflect"
	"testing"
)

func TestPowLimits(t *testing.T) {

	checkParams(t, chaincfg.MainNetParams)
	checkParams(t, chaincfg.TestNet3Params)
	checkParams(t, chaincfg.RegressionNetParams)
	checkParams(t, chaincfg.SimNetParams)
}

func checkParams(t *testing.T, params chaincfg.Params) {
	//lim := params.PowLimit
	//limBits := params.PowLimitBits
	//big := CompactToBig(limBits)
	//compact := BigToCompact(lim)
	//pin.D("net", params.Name)
	//
	//pin.D("limBits", hexPrint(limBits))
	//pin.D("compact", hexPrint(compact))
	//
	//pin.D("    lim", hexPrint(lim))
	//pin.D("    big", hexPrint(big))
}

func hexPrint(u interface{}) string {
	hex := fmt.Sprintf("%x", u)
	t := reflect.TypeOf(u)
	I := reflect.TypeOf(new(big.Int))
	if t == I {
		hex = fmt.Sprintf("%064v", hex)
	} else {

	}
	return fmt.Sprintf("(%v) %v", len(hex), hex)
}
