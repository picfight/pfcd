module github.com/picfight/dcrd/hdkeychain

require (
	github.com/decred/base58 v1.0.0
	github.com/picfight/dcrd/chaincfg v1.2.0
	github.com/picfight/dcrd/chaincfg/chainhash v1.0.1
	github.com/picfight/dcrd/dcrec v0.0.0-20180721005914-d26200ec716b
	github.com/picfight/dcrd/dcrec/secp256k1 v1.0.1
	github.com/picfight/dcrd/dcrutil v1.1.1
)

replace (
	github.com/picfight/dcrd/chaincfg => ../chaincfg
	github.com/picfight/dcrd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/dcrd/dcrec => ../dcrec
	github.com/picfight/dcrd/dcrec/edwards => ../dcrec/edwards
	github.com/picfight/dcrd/dcrec/secp256k1 => ../dcrec/secp256k1
	github.com/picfight/dcrd/dcrutil => ../dcrutil
	github.com/picfight/dcrd/wire => ../wire
)
