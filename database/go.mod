module github.com/picfight/pfcd/database

require (
	github.com/btcsuite/goleveldb v1.0.0
	github.com/picfight/pfcd/chaincfg v1.1.1
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/pfcutil v1.1.1
	github.com/picfight/pfcd/wire v1.1.0
	github.com/decred/slog v1.0.0
	github.com/jessevdk/go-flags v1.4.0
	golang.org/x/net v0.0.0-20180808004115-f9ce57c11b24 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
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
