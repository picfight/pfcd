// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2015-2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

// BlockOneLedgerMainNet is the block one output ledger for the main
// network.
var BlockOneLedgerMainNet = []*TokenPayout{
	{"DscjupVVEevuGDaRpsfn5Ur2ATkwkF3abCB", int64(7777 * 1e8)},
}

// BlockOneLedgerTestNet is the block one output ledger for the test
// network.
var BlockOneLedgerTestNet = []*TokenPayout{
	{"TsmWaPM77WSyA3aiQ2Q1KnwGDVWvEkhipBc", int64(100000 * 1e8)},
}

// BlockOneLedgerTestNet2 is the block one output ledger for the 2nd test
// network.
var BlockOneLedgerTestNet2 = []*TokenPayout{
	{"TsT5rhHYqJF7sXouh9jHtwQEn5YJ9KKc5L9", int64(100000 * 1e8)},
}

// BlockOneLedgerSimNet is the block one output ledger for the simulation
// network. See under "PicFight organization related parameters" in params.go
// for information on how to spend these outputs.
var BlockOneLedgerSimNet = []*TokenPayout{
	{"Sshw6S86G2bV6W32cbc7EhtFy8f93rU6pae", int64(100000 * 1e8)},
	{"SsjXRK6Xz6CFuBt6PugBvrkdAa4xGbcZ18w", int64(100000 * 1e8)},
	{"SsfXiYkYkCoo31CuVQw428N6wWKus2ZEw5X", int64(100000 * 1e8)},
}
