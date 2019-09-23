fees
=======


[![Build Status](http://img.shields.io/travis/picfight/pfcd.svg)](https://travis-ci.org/picfight/pfcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/picfight/pfcd/fees)

Package fees provides picfight-specific methods for tracking and estimating fee
rates for new transactions to be mined into the network. Fee rate estimation has
two main goals:

- Ensuring transactions are mined within a target _confirmation range_
  (expressed in blocks);
- Attempting to minimize fees while maintaining be above restriction.

This package was started in order to resolve issue picfight/pfcd#1412 and related.
See that issue for discussion of the selected approach.

This package was developed for pfcd, a full-node implementation of Picfight which
is under active development.  Although it was primarily written for
pfcd, this package has intentionally been designed so it can be used as a
standalone package for any projects needing the functionality provided.

## Installation and Updating

```bash
$ go get -u github.com/picfight/pfcd/fees
```

## License

Package pfcutil is licensed under the [copyfree](http://copyfree.org) ISC
License.
