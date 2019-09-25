// Copyright (c) 2013-2016 The btcsuite developers
// Copyright (c) 2015-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"github.com/picfight/pfcd/chaincfg"
)

// activeNetParams is a pointer to the parameters specific to the
// currently active Decred network.
var activeNetParams = &pfcNetParams

// params is used to group parameters for various networks such as the main
// network and test networks.
type params struct {
	*chaincfg.Params
	rpcPort string
}

// testNet3Params contains parameters specific to the test network (version 3)
// (wire.TestNet3).
var testNet3Params = params{
	Params:  &chaincfg.TestNet3Params,
	rpcPort: "19109",
}

// simNetParams contains parameters specific to the simulation test network
// (wire.SimNet).
var simNetParams = params{
	Params:  &chaincfg.SimNetParams,
	rpcPort: "19556",
}

// regNetParams contains parameters specific to the regression test
// network (wire.RegNet).
var regNetParams = params{
	Params:  &chaincfg.RegNetParams,
	rpcPort: "18656",
}

var pfcNetParams = params{
	Params:  &chaincfg.PicFightCoinNetParams,
	rpcPort: "9109",
}
