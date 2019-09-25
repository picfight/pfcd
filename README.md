pfcd
====

[![Build Status](https://travis-ci.org/picfight/pfcd.png?branch=master)](https://travis-ci.org/picfight/pfcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/picfight/pfcd)
[![Go Report Card](https://goreportcard.com/badge/github.com/picfight/pfcd)](https://goreportcard.com/report/github.com/picfight/pfcd)

## PicFight coin overview

PicFight coin is a [Decred](https://decred.org)-based cryptocurrency.
It utilizes a hybrid proof-of-work and proof-of-stake mining system.
A unit of the currency is called a `picfight coin` (PFC).

https://picfight.org

## What is pfcd?

pfcd is a full node implementation of PicFight coin written in Go (golang).

It acts as a fully-validating chain daemon.  pfcd maintains the entire past
transactional ledger of PicFight coin and allows relaying of transactions
to other PicFight coin nodes around the world.

## What is a full node?

The term 'full node' is short for 'fully-validating node' and refers to software
that fully validates all transactions and blocks, as opposed to trusting a 3rd
party. In addition to validating transactions and blocks, nearly all full nodes
also participate in relaying transactions and blocks to other full nodes around
the world, thus forming the peer-to-peer network.

## Getting Started

So, you've decided to help the network by running a full node.  Great!  Running
pfcd is simple.  All you need to do is install pfcd on a machine that is
connected to the internet and launch it.

Also, make sure your firewall is configured to allow inbound connections to port
9108.

## Installing and updating

### Build from source (all platforms)

Building or updating from source requires the following build dependencies:

- **Go 1.13**

  Installation instructions can be found here: https://golang.org/doc/install.
  It is recommended to add `$GOPATH/bin` to your `PATH` at this point.

- **Git**

  Installation instructions can be found at https://git-scm.com or
  https://gitforwindows.org.

To build and install from a checked-out repo, run `go install . ./cmd/...` in
the repo's root directory.  Some notes:

* Set the `GO111MODULE=on` environment variable if using Go 1.11 and building
  from within `GOPATH`.

* The `pfcd` executable will be installed to `$GOPATH/bin`.  `GOPATH`
  defaults to `$HOME/go` (or `%USERPROFILE%\go` on Windows) if unset.

### Example of obtaining and building from source on Windows:

```bash
git clone https://github.com/picfight/pfcd ~/src/pfcd
cd ~/src/pfcd
set GO111MODULE=on
go build ./...
go install . .\cmd\...
```

### Running Tests

To run the tests locally:

```bash
set GO111MODULE=on
go build ./...
go clean -testcache
go test ./...
```

## Example run commands

Launch mining node:
```bash
pfcd
     --generate
     --miningaddr "JsCVh5SVDQovpW1dswaZNan2mfNWy6uRpPx"
     --listen=127.0.0.1:30000
     --rpclisten=127.0.0.1:30001
     --datadir=nodeA
     --rpccert=nodeA\rpc.cert
     --rpckey=nodeA\rpc.key     
     --txindex
     --addrindex
     --rpcuser=node.user
     --rpcpass=node.pass
```

Launch second node and connect to it the first one for syncing:
```bash
pfcd
     --listen=127.0.0.1:30002
     --rpclisten=127.0.0.1:30003
     --addpeer=127.0.0.1:30000
     --datadir=nodeB
     --rpccert=nodeB\rpc.cert
     --rpckey=nodeB\rpc.key
     --txindex
     --addrindex
     --rpcuser=node.user
     --rpcpass=node.pass
     
```

## Contact

If you have any further questions you can find us at the
[integrated github issue tracker](https://github.com/picfight/pfcd/issues).

## Documentation

Since pfcd is a fork of Decred, the documentation for pfcd is located in the
[Decred-docs](https://github.com/decred/dcrd/tree/master/docs) folder.

## License

pfcd is licensed under the [copyfree](http://copyfree.org) ISC License.
