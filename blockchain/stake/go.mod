module github.com/picfight/pfcd/blockchain/stake

require (
	github.com/decred/slog v1.0.0
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f // indirect
	golang.org/x/sys v0.0.0-20181211161752-7da8ea5c8182 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

replace (
	github.com/picfight/pfcd/chaincfg => ../../chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ../../chaincfg/chainhash
	github.com/picfight/pfcd/database => ../../database
	github.com/picfight/pfcd/pfcec => ../../pfcec
	github.com/picfight/pfcd/pfcec/edwards => ../../pfcec/edwards
	github.com/picfight/pfcd/pfcec/secp256k1 => ../../pfcec/secp256k1
	github.com/picfight/pfcd/pfcutil => ../../pfcutil
	github.com/picfight/pfcd/txscript => ../../txscript
	github.com/picfight/pfcd/wire => ../../wire
)
