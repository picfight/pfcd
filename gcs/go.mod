module github.com/picfight/pfcd/gcs

require (
	github.com/dchest/blake256 v1.0.0
	github.com/dchest/siphash v1.2.0
	github.com/picfight/pfcd/blockchain/stake v1.0.1
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/txscript v1.0.1
	github.com/picfight/pfcd/wire v1.2.0
)

replace (
	github.com/picfight/pfcd/blockchain/stake => ../blockchain/stake
	github.com/picfight/pfcd/chaincfg => ../chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/database => ../database
	github.com/picfight/pfcd/pfcec => ../pfcec
	github.com/picfight/pfcd/pfcec/edwards => ../pfcec/edwards
	github.com/picfight/pfcd/pfcec/secp256k1 => ../pfcec/secp256k1
	github.com/picfight/pfcd/pfcutil => ../pfcutil
	github.com/picfight/pfcd/txscript => ../txscript
	github.com/picfight/pfcd/wire => ../wire
)
