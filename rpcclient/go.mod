module github.com/picfight/dcrd/rpcclient

require (
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd
	github.com/davecgh/go-spew v1.1.0
	github.com/picfight/dcrd/chaincfg/chainhash v1.0.1
	github.com/picfight/dcrd/dcrjson v1.0.0
	github.com/picfight/dcrd/dcrutil v1.1.1
	github.com/picfight/dcrd/gcs v1.0.1
	github.com/picfight/dcrd/wire v1.2.0
	github.com/decred/slog v1.0.0
	github.com/gorilla/websocket v1.2.0
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
	github.com/picfight/dcrd/dcrjson => ../dcrjson
	github.com/picfight/dcrd/dcrutil => ../dcrutil
	github.com/picfight/dcrd/gcs => ../gcs
	github.com/picfight/dcrd/txscript => ../txscript
	github.com/picfight/dcrd/wire => ../wire
)
