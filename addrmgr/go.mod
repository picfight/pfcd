module github.com/picfight/pfcd/addrmgr

require (
	github.com/decred/slog v1.0.0
)

replace (
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/wire => ../wire
)
