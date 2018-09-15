// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2015-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

// BlockOneLedgerMainNet is the block one output ledger for the main
// network.
var BlockOneLedgerMainNet = []*TokenPayout{
	//{"DsZZX6MSkR1RWThuZkmtrcWs5exTHiVUrE3", 282.63795424 * 1e8},
}

// BlockOneLedgerTestNet3 is the block one output ledger for testnet version 3.
var BlockOneLedgerTestNet3 = []*TokenPayout{}

// BlockOneLedgerSimNet is the block one output ledger for the simulation
// network. See under "PicFight organization related parameters" in params.go
// for information on how to spend these outputs.
var BlockOneLedgerSimNet = []*TokenPayout{}
