module github.com/picfight/pfcd/mining

require (
	github.com/picfight/pfcd/blockchain v1.0.1
	github.com/picfight/pfcd/blockchain/stake v1.0.1
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/dcrutil v1.1.1
	github.com/picfight/pfcd/wire v1.2.0
)

replace (
	github.com/picfight/pfcd/blockchain => ../blockchain
	github.com/picfight/pfcd/blockchain/stake => ../blockchain/stake
	github.com/picfight/pfcd/chaincfg => ../chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/database => ../database
	github.com/picfight/pfcd/dcrec => ../dcrec
	github.com/picfight/pfcd/dcrec/edwards => ../dcrec/edwards
	github.com/picfight/pfcd/dcrec/secp256k1 => ../dcrec/secp256k1
	github.com/picfight/pfcd/dcrutil => ../dcrutil
	github.com/picfight/pfcd/gcs => ../gcs
	github.com/picfight/pfcd/txscript => ../txscript
	github.com/picfight/pfcd/wire => ../wire
)
