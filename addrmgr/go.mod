module github.com/picfight/pfcd/addrmgr

require (
	github.com/decred/slog v1.0.0
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/wire v1.1.0
)

replace (
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/wire => ../wire
)
