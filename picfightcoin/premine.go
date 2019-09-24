package picfightcoin

import (
	"encoding/hex"
	"github.com/decred/dcrd/chaincfg"
)

var BlockOneLedgerPicfightCoin = []*chaincfg.TokenPayout{
	{"JsCVh5SVDQovpW1dswaZNan2mfNWy6uRpPx", 9000000 * 1e8},
}

func hexDecode(hexStr string) []byte {
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}
	return b
}
