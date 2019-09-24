module github.com/picfight/dcrd/blockchain

require (
	github.com/picfight/dcrd/blockchain/stake v1.1.0
	github.com/picfight/dcrd/chaincfg v1.2.0
	github.com/picfight/dcrd/chaincfg/chainhash v1.0.1
	github.com/picfight/dcrd/database v1.0.3
	github.com/picfight/dcrd/dcrec v0.0.0-20180801202239-0761de129164
	github.com/picfight/dcrd/dcrec/secp256k1 v1.0.1
	github.com/picfight/dcrd/dcrutil v1.2.0
	github.com/picfight/dcrd/gcs v1.0.1
	github.com/picfight/dcrd/txscript v1.0.2
	github.com/picfight/dcrd/wire v1.2.0
	github.com/decred/slog v1.0.0
)

replace (
	github.com/picfight/dcrd/blockchain/stake => ./stake
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
