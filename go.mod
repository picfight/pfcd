module github.com/picfight/pfcd

require (
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd
	github.com/btcsuite/goleveldb v1.0.0
	github.com/btcsuite/snappy-go v1.0.0
	github.com/btcsuite/winsvc v1.0.0
	github.com/davecgh/go-spew v1.1.0
	github.com/dchest/blake256 v1.0.0
	github.com/dchest/siphash v1.2.1
	github.com/decred/base58 v1.0.0
	github.com/picfight/pfcd/addrmgr v1.0.2
	github.com/picfight/pfcd/blockchain v1.1.1
	github.com/picfight/pfcd/blockchain/stake v1.1.0
	github.com/picfight/pfcd/certgen v1.0.2
	github.com/picfight/pfcd/chaincfg v1.3.0
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/connmgr v1.0.2
	github.com/picfight/pfcd/database v1.0.3
	github.com/picfight/pfcd/pfcec v0.0.0-20180801202239-0761de129164
	github.com/picfight/pfcd/pfcec/secp256k1 v1.0.1
	github.com/picfight/pfcd/pfcjson v1.1.0
	github.com/picfight/pfcd/pfcutil v1.2.0
	github.com/picfight/pfcd/fees v1.0.0
	github.com/picfight/pfcd/gcs v1.0.2
	github.com/picfight/pfcd/hdkeychain v1.1.1
	github.com/picfight/pfcd/mempool v1.1.0
	github.com/picfight/pfcd/mining v1.1.0
	github.com/picfight/pfcd/peer v1.1.0
	github.com/picfight/pfcd/rpcclient v1.1.0
	github.com/picfight/pfcd/txscript v1.0.2
	github.com/picfight/pfcd/wire v1.2.0
	github.com/decred/slog v1.0.0
	github.com/gorilla/websocket v1.2.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/jrick/bitset v1.0.0
	github.com/jrick/logrotate v1.0.0
	golang.org/x/crypto v0.0.0-20180718160520-a2144134853f
	golang.org/x/sys v0.0.0-20181211161752-7da8ea5c8182
)

replace (
	github.com/picfight/pfcd/addrmgr => ./addrmgr
	github.com/picfight/pfcd/blockchain => ./blockchain
	github.com/picfight/pfcd/blockchain/stake => ./blockchain/stake
	github.com/picfight/pfcd/certgen => ./certgen
	github.com/picfight/pfcd/chaincfg => ./chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ./chaincfg/chainhash
	github.com/picfight/pfcd/connmgr => ./connmgr
	github.com/picfight/pfcd/database => ./database
	github.com/picfight/pfcd/pfcec => ./pfcec
	github.com/picfight/pfcd/pfcec/edwards => ./pfcec/edwards
	github.com/picfight/pfcd/pfcec/secp256k1 => ./pfcec/secp256k1
	github.com/picfight/pfcd/pfcjson => ./pfcjson
	github.com/picfight/pfcd/pfcutil => ./pfcutil
	github.com/picfight/pfcd/fees => ./fees
	github.com/picfight/pfcd/gcs => ./gcs
	github.com/picfight/pfcd/hdkeychain => ./hdkeychain
	github.com/picfight/pfcd/limits => ./limits
	github.com/picfight/pfcd/mempool => ./mempool
	github.com/picfight/pfcd/mining => ./mining
	github.com/picfight/pfcd/peer => ./peer
	github.com/picfight/pfcd/rpcclient => ./rpcclient
	github.com/picfight/pfcd/txscript => ./txscript
	github.com/picfight/pfcd/wire => ./wire
)
