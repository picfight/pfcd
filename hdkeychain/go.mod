module github.com/picfight/pfcd/hdkeychain

require (
	github.com/decred/base58 v1.0.0
	github.com/picfight/pfcd/chaincfg v1.2.0
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/dcrec v0.0.0-20180721005914-d26200ec716b
	github.com/picfight/pfcd/dcrec/secp256k1 v1.0.1
	github.com/picfight/pfcd/dcrutil v1.1.1
)

replace (
	github.com/picfight/pfcd/chaincfg => ../chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/dcrec => ../dcrec
	github.com/picfight/pfcd/dcrec/edwards => ../dcrec/edwards
	github.com/picfight/pfcd/dcrec/secp256k1 => ../dcrec/secp256k1
	github.com/picfight/pfcd/dcrutil => ../dcrutil
	github.com/picfight/pfcd/wire => ../wire
)
