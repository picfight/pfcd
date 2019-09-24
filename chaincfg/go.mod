module github.com/picfight/dcrd/chaincfg

require (
	github.com/davecgh/go-spew v1.1.0
	github.com/picfight/dcrd/chaincfg/chainhash v1.0.1
	github.com/picfight/dcrd/dcrec/edwards v0.0.0-20181208004914-a0816cf4301f
	github.com/picfight/dcrd/dcrec/secp256k1 v1.0.1
	github.com/picfight/dcrd/wire v1.2.0
)

replace (
	github.com/picfight/dcrd/chaincfg/chainhash => ./chainhash
	github.com/picfight/dcrd/dcrec/edwards => ../dcrec/edwards
	github.com/picfight/dcrd/dcrec/secp256k1 => ../dcrec/secp256k1
	github.com/picfight/dcrd/wire => ../wire
)
