package picfightcoin

import (
	"fmt"
	"github.com/decred/base58"
	"github.com/decred/dcrd/chaincfg"
	"github.com/decred/dcrd/dcrec"
	"github.com/decred/dcrd/dcrutil"
)

// DecodeAddress decodes the string encoding of an address and returns
// the Address if addr is a valid encoding for a known address type
func DecodeAddress(addr string) (dcrutil.Address, error) {
	// Switch on decoded length to determine the type.
	decoded, netID, err := base58.CheckDecode(addr)
	if err != nil {
		if err == base58.ErrChecksum {
			return nil, dcrutil.ErrChecksumMismatch
		}
		return nil, fmt.Errorf("decoded address is of unknown format: %v",
			err.Error())
	}

	net, err := detectNetworkForAddress(addr)
	if err != nil {
		return nil, dcrutil.ErrUnknownAddressType
	}

	switch netID {
	case net.PubKeyAddrID:
		return dcrutil.NewAddressPubKey(decoded, net)

	case net.PubKeyHashAddrID:
		return dcrutil.NewAddressPubKeyHash(decoded, net, dcrec.STEcdsaSecp256k1)

	case net.PKHEdwardsAddrID:
		return dcrutil.NewAddressPubKeyHash(decoded, net, dcrec.STEd25519)

	case net.PKHSchnorrAddrID:
		return dcrutil.NewAddressPubKeyHash(decoded, net, dcrec.STSchnorrSecp256k1)

	case net.ScriptHashAddrID:
		return dcrutil.NewAddressScriptHashFromHash(decoded, net)

	default:
		return nil, dcrutil.ErrUnknownAddressType
	}
}

// detectNetworkForAddress pops the first character from a string encoded
// address and detects what network type it is for.
func detectNetworkForAddress(addr string) (*chaincfg.Params, error) {
	if len(addr) < 1 {
		return nil, fmt.Errorf("empty string given for network detection")
	}

	networkChar := addr[0:1]
	switch networkChar {
	case PicFightCoinNetParams.NetworkAddressPrefix:
		return &PicFightCoinNetParams, nil
	case chaincfg.TestNet3Params.NetworkAddressPrefix:
		return &chaincfg.TestNet3Params, nil
	case chaincfg.SimNetParams.NetworkAddressPrefix:
		return &chaincfg.SimNetParams, nil
	case chaincfg.RegNetParams.NetworkAddressPrefix:
		return &chaincfg.RegNetParams, nil
	}

	return nil, fmt.Errorf("unknown network type in string encoded address")
}
