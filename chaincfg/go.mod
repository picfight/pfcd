module github.com/picfight/pfcd/chaincfg

require (
	github.com/davecgh/go-spew v1.1.0
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/pfcec/edwards v0.0.0-20181208004914-a0816cf4301f
	github.com/picfight/pfcd/pfcec/secp256k1 v1.0.1
	github.com/picfight/pfcd/wire v1.2.0
)

replace (
	github.com/picfight/pfcd/chaincfg/chainhash => ./chainhash
	github.com/picfight/pfcd/pfcec/edwards => ../pfcec/edwards
	github.com/picfight/pfcd/pfcec/secp256k1 => ../pfcec/secp256k1
	github.com/picfight/pfcd/wire => ../wire
)
