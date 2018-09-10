module github.com/picfight/pfcd/connmgr

require (
	github.com/picfight/pfcd/chaincfg v1.1.1
	github.com/picfight/pfcd/wire v1.1.0
	github.com/decred/slog v1.0.0
)

replace (
	github.com/picfight/pfcd/chaincfg => ../chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/wire => ../wire
)
