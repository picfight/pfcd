module github.com/picfight/pfcd/pfcutil

require (
	github.com/davecgh/go-spew v1.1.0
	github.com/decred/base58 v1.0.0
	github.com/picfight/pfcd/chaincfg v1.2.0
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/pfcec v0.0.0-20180721005212-59fe2b293f69
	github.com/picfight/pfcd/pfcec/edwards v0.0.0-20181208004914-a0816cf4301f
	github.com/picfight/pfcd/pfcec/secp256k1 v1.0.1
	github.com/picfight/pfcd/wire v1.2.0
	golang.org/x/crypto v0.0.0-20180718160520-a2144134853f
)

replace (
	github.com/picfight/pfcd/chaincfg => ../chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/pfcec => ../pfcec
	github.com/picfight/pfcd/pfcec/edwards => ../pfcec/edwards
	github.com/picfight/pfcd/pfcec/secp256k1 => ../pfcec/secp256k1
	github.com/picfight/pfcd/wire => ../wire
)
