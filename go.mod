module github.com/picfight/dcrd

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
	github.com/picfight/dcrd/addrmgr v1.0.2
	github.com/picfight/dcrd/blockchain v1.1.1
	github.com/picfight/dcrd/blockchain/stake v1.1.0
	github.com/picfight/dcrd/certgen v1.0.2
	github.com/picfight/dcrd/chaincfg v1.3.0
	github.com/picfight/dcrd/chaincfg/chainhash v1.0.1
	github.com/picfight/dcrd/connmgr v1.0.2
	github.com/picfight/dcrd/database v1.0.3
	github.com/picfight/dcrd/dcrec v0.0.0-20180801202239-0761de129164
	github.com/picfight/dcrd/dcrec/secp256k1 v1.0.1
	github.com/picfight/dcrd/dcrjson v1.1.0
	github.com/picfight/dcrd/dcrutil v1.2.0
	github.com/picfight/dcrd/fees v1.0.0
	github.com/picfight/dcrd/gcs v1.0.2
	github.com/picfight/dcrd/hdkeychain v1.1.1
	github.com/picfight/dcrd/mempool v1.1.0
	github.com/picfight/dcrd/mining v1.1.0
	github.com/picfight/dcrd/peer v1.1.0
	github.com/picfight/dcrd/rpcclient v1.1.0
	github.com/picfight/dcrd/txscript v1.0.2
	github.com/picfight/dcrd/wire v1.2.0
	github.com/decred/slog v1.0.0
	github.com/gorilla/websocket v1.2.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/jrick/bitset v1.0.0
	github.com/jrick/logrotate v1.0.0
	golang.org/x/crypto v0.0.0-20180718160520-a2144134853f
	golang.org/x/sys v0.0.0-20181211161752-7da8ea5c8182
)

replace (
	github.com/picfight/dcrd/addrmgr => ./addrmgr
	github.com/picfight/dcrd/blockchain => ./blockchain
	github.com/picfight/dcrd/blockchain/stake => ./blockchain/stake
	github.com/picfight/dcrd/certgen => ./certgen
	github.com/picfight/dcrd/chaincfg => ./chaincfg
	github.com/picfight/dcrd/chaincfg/chainhash => ./chaincfg/chainhash
	github.com/picfight/dcrd/connmgr => ./connmgr
	github.com/picfight/dcrd/database => ./database
	github.com/picfight/dcrd/dcrec => ./dcrec
	github.com/picfight/dcrd/dcrec/edwards => ./dcrec/edwards
	github.com/picfight/dcrd/dcrec/secp256k1 => ./dcrec/secp256k1
	github.com/picfight/dcrd/dcrjson => ./dcrjson
	github.com/picfight/dcrd/dcrutil => ./dcrutil
	github.com/picfight/dcrd/fees => ./fees
	github.com/picfight/dcrd/gcs => ./gcs
	github.com/picfight/dcrd/hdkeychain => ./hdkeychain
	github.com/picfight/dcrd/limits => ./limits
	github.com/picfight/dcrd/mempool => ./mempool
	github.com/picfight/dcrd/mining => ./mining
	github.com/picfight/dcrd/peer => ./peer
	github.com/picfight/dcrd/rpcclient => ./rpcclient
	github.com/picfight/dcrd/txscript => ./txscript
	github.com/picfight/dcrd/wire => ./wire
)
