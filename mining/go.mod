module github.com/picfight/dcrd/mining

require (
	github.com/picfight/dcrd/blockchain v1.0.1
	github.com/picfight/dcrd/blockchain/stake v1.0.1
	github.com/picfight/dcrd/chaincfg/chainhash v1.0.1
	github.com/picfight/dcrd/dcrutil v1.1.1
	github.com/picfight/dcrd/wire v1.2.0
)

replace (
	github.com/picfight/dcrd/blockchain => ../blockchain
	github.com/picfight/dcrd/blockchain/stake => ../blockchain/stake
	github.com/picfight/dcrd/chaincfg => ../chaincfg
	github.com/picfight/dcrd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/dcrd/database => ../database
	github.com/picfight/dcrd/dcrec => ../dcrec
	github.com/picfight/dcrd/dcrec/edwards => ../dcrec/edwards
	github.com/picfight/dcrd/dcrec/secp256k1 => ../dcrec/secp256k1
	github.com/picfight/dcrd/dcrutil => ../dcrutil
	github.com/picfight/dcrd/gcs => ../gcs
	github.com/picfight/dcrd/txscript => ../txscript
	github.com/picfight/dcrd/wire => ../wire
)
