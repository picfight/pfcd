wire
====

[![Build Status](http://img.shields.io/travis/picfight/pfcd.svg)](https://travis-ci.org/picfight/pfcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/picfight/pfcd/wire)

Package wire implements the Picfight wire protocol.  A comprehensive suite of
tests with 100% test coverage is provided to ensure proper functionality.

This package has intentionally been designed so it can be used as a standalone
package for any projects needing to interface with Picfight peers at the wire
protocol level.

## Installation and Updating

```bash
$ go get -u github.com/picfight/pfcd/wire
```

## Picfight Message Overview

The Picfight protocol consists of exchanging messages between peers. Each message
is preceded by a header which identifies information about it such as which
picfight network it is a part of, its type, how big it is, and a checksum to
verify validity. All encoding and decoding of message headers is handled by this
package.

To accomplish this, there is a generic interface for Picfight messages named
`Message` which allows messages of any type to be read, written, or passed
around through channels, functions, etc. In addition, concrete implementations
of most of the currently supported Picfight messages are provided. For these
supported messages, all of the details of marshalling and unmarshalling to and
from the wire using Picfight encoding are handled so the caller doesn't have to
concern themselves with the specifics.

## Reading Messages Example

In order to unmarshal Picfight messages from the wire, use the `ReadMessage`
function. It accepts any `io.Reader`, but typically this will be a `net.Conn`
to a remote node running a Picfight peer.  Example syntax is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main Picfight network.
	pver := wire.ProtocolVersion
	pfcnet := wire.MainNet

	// Reads and validates the next Picfight message from conn using the
	// protocol version pver and the Picfight network pfcnet.  The returns
	// are a wire.Message, a []byte which contains the unmarshalled
	// raw payload, and a possible error.
	msg, rawPayload, err := wire.ReadMessage(conn, pver, pfcnet)
	if err != nil {
		// Log and handle the error
	}
```

See the package documentation for details on determining the message type.

## Writing Messages Example

In order to marshal Picfight messages to the wire, use the `WriteMessage`
function. It accepts any `io.Writer`, but typically this will be a `net.Conn`
to a remote node running a Picfight peer. Example syntax to request addresses
from a remote peer is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main Picfight network.
	pver := wire.ProtocolVersion
	pfcnet := wire.MainNet

	// Create a new getaddr Picfight message.
	msg := wire.NewMsgGetAddr()

	// Writes a Picfight message msg to conn using the protocol version
	// pver, and the Picfight network pfcnet.  The return is a possible
	// error.
	err := wire.WriteMessage(conn, msg, pver, pfcnet)
	if err != nil {
		// Log and handle the error
	}
```

## License

Package wire is licensed under the [copyfree](http://copyfree.org) ISC
License.
