rpcclient
=========

[![Build Status](http://img.shields.io/travis/picfight/pfcd.svg)](https://travis-ci.org/picfight/pfcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/picfight/pfcd/rpcclient)

rpcclient implements a Websocket-enabled Picfightcoin JSON-RPC client package written
in [Go](http://golang.org/).  It provides a robust and easy to use client for
interfacing with a Picfightcoin RPC server that uses a btcd/bitcoin core compatible
Picfightcoin JSON-RPC API.

## Status

This package is currently under active development.  It is already stable and
the infrastructure is complete.  However, there are still several RPCs left to
implement and the API is not stable yet.

## Documentation

* [API Reference](http://godoc.org/github.com/picfight/pfcd/rpcclient)
* [pfcd Websockets Example](https://github.com/picfight/pfcd/tree/master/rpcclient/examples/btcdwebsockets)
  Connects to a pfcd RPC server using TLS-secured websockets, registers for
  block connected and block disconnected notifications, and gets the current
  block count
* [pfcwallet Websockets Example](https://github.com/picfight/pfcd/tree/master/rpcclient/examples/pfcwalletwebsockets)
  Connects to a pfcwallet RPC server using TLS-secured websockets, registers for
  notifications about changes to account balances, and gets a list of unspent
  transaction outputs (utxos) the wallet can sign
* [Picfightcoin Core HTTP POST Example](https://github.com/picfight/pfcd/tree/master/rpcclient/examples/bitcoincorehttp)
  Connects to a picfightcoin core RPC server using HTTP POST mode with TLS disabled
  and gets the current block count

## Major Features

* Supports Websockets (pfcd/pfcwallet) and HTTP POST mode (bitcoin core)
* Provides callback and registration functions for pfcd/pfcwallet notifications
* Supports pfcd extensions
* Translates to and from higher-level and easier to use Go types
* Offers a synchronous (blocking) and asynchronous API
* When running in Websockets mode (the default):
  * Automatic reconnect handling (can be disabled)
  * Outstanding commands are automatically reissued
  * Registered notifications are automatically reregistered
  * Back-off support on reconnect attempts

## Installation

```bash
$ go get -u github.com/picfight/pfcd/rpcclient
```

## License

Package rpcclient is licensed under the [copyfree](http://copyfree.org) ISC
License.
