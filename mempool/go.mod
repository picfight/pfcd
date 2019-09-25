module github.com/picfight/pfcd/mempool

require (
	github.com/picfight/pfcd/blockchain v1.0.1
	github.com/picfight/pfcd/blockchain/stake v1.1.0
	github.com/picfight/pfcd/chaincfg v1.2.0
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/dcrec v0.0.0-20180801202239-0761de129164
	github.com/picfight/pfcd/dcrec/secp256k1 v1.0.1
	github.com/picfight/pfcd/dcrjson v1.0.0
	github.com/picfight/pfcd/dcrutil v1.2.0
	github.com/picfight/pfcd/mining v1.1.0
	github.com/picfight/pfcd/txscript v1.0.2
	github.com/picfight/pfcd/wire v1.2.0
	github.com/decred/slog v1.0.0
)

replace (
	github.com/picfight/pfcd/blockchain => ../blockchain
	github.com/picfight/pfcd/blockchain/stake => ../blockchain/stake
	github.com/picfight/pfcd/chaincfg => ../chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/database => ../database
	github.com/picfight/pfcd/dcrec/secp256k1 => ../dcrec/secp256k1
	github.com/picfight/pfcd/dcrjson => ../dcrjson
	github.com/picfight/pfcd/dcrutil => ../dcrutil
	github.com/picfight/pfcd/gcs => ../gcs
	github.com/picfight/pfcd/mining => ../mining
	github.com/picfight/pfcd/txscript => ../txscript
	github.com/picfight/pfcd/wire => ../wire
)
