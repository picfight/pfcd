module github.com/picfight/dcrd/addrmgr

require (
	github.com/picfight/dcrd/chaincfg/chainhash v1.0.1
	github.com/picfight/dcrd/wire v1.1.0
	github.com/decred/slog v1.0.0
)

replace (
	github.com/picfight/dcrd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/dcrd/wire => ../wire
)
