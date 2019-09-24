pfcd
====

[![Build Status](https://travis-ci.org/picfight/pfcd.png?branch=master)](https://travis-ci.org/picfight/pfcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/picfight/pfcd)
[![Go Report Card](https://goreportcard.com/badge/github.com/picfight/pfcd)](https://goreportcard.com/report/github.com/picfight/pfcd)

## PicFight coin Overview

PicFight coin is based on Decred (https://decred.org).
It utilizes a hybrid proof-of-work and proof-of-stake mining system.
A unit of the currency is called a `picfight coin` (PFC).

https://picfight.org

## Latest Downloads

https://picfight.org/downloads

## What is pfcd?

pfcd is a full node implementation of PicFight coin written in Go (golang).

It acts as a fully-validating chain daemon for the PicFight coin cryptocurrency.  pfcd
maintains the entire past transactional ledger of PicFight coin and allows relaying of
transactions to other PicFight coin nodes around the world.

This software is currently under active development.  It is extremely stable and
has been in production use since February 2016.

## What is a full node?

The term 'full node' is short for 'fully-validating node' and refers to software
that fully validates all transactions and blocks, as opposed to trusting a 3rd
party.  In addition to validating transactions and blocks, nearly all full nodes
also participate in relaying transactions and blocks to other full nodes around
the world, thus forming the peer-to-peer network that is the backbone of the
PicFight coin cryptocurrency.

The full node distinction is important, since full nodes are not the only type
of software participating in the PicFight coin peer network. For instance, there are
'lightweight nodes' which rely on full nodes to serve the transactions, blocks,
and cryptographic proofs they require to function, as well as relay their
transactions to the rest of the global network.

## Getting Started

So, you've decided to help the network by running a full node.  Great!  Running
pfcd is simple.  All you need to do is install pfcd on a machine that is
connected to the internet and meets the minimum recommended specifications, and
launch it.

Also, make sure your firewall is configured to allow inbound connections to port
9108.

<a name="Installation" />

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

* Replace `go` with `vgo` when using Go 1.10.

* The `pfcd` executable will be installed to `$GOPATH/bin`.  `GOPATH`
  defaults to `$HOME/go` (or `%USERPROFILE%\go` on Windows) if unset.


### Example of obtaining and building from source on Windows 10 with Go 1.11:

```PowerShell
PS> git clone https://github.com/picfight/pfcd $env:USERPROFILE\src\pfcd
PS> cd $env:USERPROFILE\src\pfcd
PS> go install . .\cmd\...
PS> & "$(go env GOPATH)\bin\pfcd" -V

```

### Example of obtaining and building from source:

```bash
$ git clone https://github.com/picfight/pfcd ~/src/pfcd
$ cd ~/src/pfcd
set GO111MODULE=on
go build ./...
```


### Running Tests

To run the tests locally:

```
set GO111MODULE=on
go build ./...
go clean -testcache
go test ./...
```

## Contact

If you have any further questions you can find us at the
[integrated github issue tracker](https://github.com/picfight/pfcd/issues).


## Documentation

Since pfcd is a code fork of Decred, the documentation for pfcd is located in the
[docs](https://github.com/decred/dcrd/tree/master/docs) folder.

## License

pfcd is licensed under the [copyfree](http://copyfree.org) ISC License.
