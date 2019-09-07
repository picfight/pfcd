// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2015-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

// BlockOneLedgerMainNet is the block one output ledger for the main
// network.
var BlockOneLedgerMainNet = []*TokenPayout{}

// BlockOneLedgerTestNet3 is the block one output ledger for testnet version 3.
var BlockOneLedgerTestNet3 = []*TokenPayout{}

// BlockOneLedgerSimNet is the block one output ledger for the simulation
// network.  See "PicFight organization related parameters" in simnetparams.go for
// information on how to spend these outputs.
var BlockOneLedgerSimNet = []*TokenPayout{}

// BlockOneLedgerRegNet is the block one output ledger for the regression test
// network.  See "PicFight organization related parameters" in regnetparams.go for
// information on how to spend these outputs.
var BlockOneLedgerRegNet = []*TokenPayout{}
