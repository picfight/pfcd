module github.com/picfight/pfcd/hdkeychain

require (
	github.com/decred/base58 v1.0.0
	github.com/picfight/pfcd/chaincfg v1.1.1
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/pfcec v0.0.0-20180721005914-d26200ec716b
	github.com/picfight/pfcd/pfcec/secp256k1 v1.0.0
	github.com/picfight/pfcd/pfcutil v1.1.1
)

replace (
	github.com/picfight/pfcd/chaincfg => ../chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/pfcec => ../pfcec
	github.com/picfight/pfcd/pfcec/edwards => ../pfcec/edwards
	github.com/picfight/pfcd/pfcec/secp256k1 => ../pfcec/secp256k1
	github.com/picfight/pfcd/pfcutil => ../pfcutil
	github.com/picfight/pfcd/wire => ../wire
)
