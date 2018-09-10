module github.com/picfight/pfcd/chaincfg

require (
	github.com/davecgh/go-spew v1.1.0
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/wire v1.1.0
)

replace (
	github.com/picfight/pfcd/chaincfg/chainhash => ./chainhash
	github.com/picfight/pfcd/wire => ../wire
)
