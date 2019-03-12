module github.com/picfight/pfcd/pfcutil

require (
	github.com/davecgh/go-spew v1.1.0
	github.com/decred/base58 v1.0.0
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
