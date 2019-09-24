module github.com/picfight/dcrd/connmgr

require (
	github.com/picfight/dcrd/chaincfg v1.1.1
	github.com/picfight/dcrd/wire v1.2.0
	github.com/decred/slog v1.0.0
)

replace (
	github.com/picfight/dcrd/chaincfg => ../chaincfg
	github.com/picfight/dcrd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/dcrd/dcrec/secp256k1 => ../dcrec/secp256k1
	github.com/picfight/dcrd/wire => ../wire
)
