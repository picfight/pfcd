// Copyright (c) 2014-2017 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package pfcjson_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/picfight/pfcd/pfcjson"
)

// TestChainSvrWsCmds tests all of the chain server websocket-specific commands
// marshal and unmarshal into valid results include handling of optional fields
// being omitted in the marshalled command, while optional fields with defaults
// have the default assigned on unmarshalled commands.
func TestChainSvrWsCmds(t *testing.T) {
	t.Parallel()

	testID := int(1)
	tests := []struct {
		name         string
		newCmd       func() (interface{}, error)
		staticCmd    func() interface{}
		marshalled   string
		unmarshalled interface{}
	}{
		{
			name: "authenticate",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("authenticate", "user", "pass")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewAuthenticateCmd("user", "pass")
			},
			marshalled:   `{"jsonrpc":"1.0","method":"authenticate","params":["user","pass"],"id":1}`,
			unmarshalled: &pfcjson.AuthenticateCmd{Username: "user", Passphrase: "pass"},
		},
		{
			name: "notifyblocks",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("notifyblocks")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewNotifyBlocksCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"notifyblocks","params":[],"id":1}`,
			unmarshalled: &pfcjson.NotifyBlocksCmd{},
		},
		{
			name: "stopnotifyblocks",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("stopnotifyblocks")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewStopNotifyBlocksCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"stopnotifyblocks","params":[],"id":1}`,
			unmarshalled: &pfcjson.StopNotifyBlocksCmd{},
		},
		{
			name: "notifynewtransactions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("notifynewtransactions")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewNotifyNewTransactionsCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"notifynewtransactions","params":[],"id":1}`,
			unmarshalled: &pfcjson.NotifyNewTransactionsCmd{
				Verbose: pfcjson.Bool(false),
			},
		},
		{
			name: "notifynewtransactions optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("notifynewtransactions", true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewNotifyNewTransactionsCmd(pfcjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"notifynewtransactions","params":[true],"id":1}`,
			unmarshalled: &pfcjson.NotifyNewTransactionsCmd{
				Verbose: pfcjson.Bool(true),
			},
		},
		{
			name: "stopnotifynewtransactions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("stopnotifynewtransactions")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewStopNotifyNewTransactionsCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"stopnotifynewtransactions","params":[],"id":1}`,
			unmarshalled: &pfcjson.StopNotifyNewTransactionsCmd{},
		},
		{
			name: "notifyreceived",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("notifyreceived", []string{"1Address"})
			},
			staticCmd: func() interface{} {
				return pfcjson.NewNotifyReceivedCmd([]string{"1Address"})
			},
			marshalled: `{"jsonrpc":"1.0","method":"notifyreceived","params":[["1Address"]],"id":1}`,
			unmarshalled: &pfcjson.NotifyReceivedCmd{
				Addresses: []string{"1Address"},
			},
		},
		{
			name: "stopnotifyreceived",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("stopnotifyreceived", []string{"1Address"})
			},
			staticCmd: func() interface{} {
				return pfcjson.NewStopNotifyReceivedCmd([]string{"1Address"})
			},
			marshalled: `{"jsonrpc":"1.0","method":"stopnotifyreceived","params":[["1Address"]],"id":1}`,
			unmarshalled: &pfcjson.StopNotifyReceivedCmd{
				Addresses: []string{"1Address"},
			},
		},
		{
			name: "notifyspent",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("notifyspent", `[{"hash":"123","index":0}]`)
			},
			staticCmd: func() interface{} {
				ops := []pfcjson.OutPoint{{Hash: "123", Index: 0}}
				return pfcjson.NewNotifySpentCmd(ops)
			},
			marshalled: `{"jsonrpc":"1.0","method":"notifyspent","params":[[{"hash":"123","index":0}]],"id":1}`,
			unmarshalled: &pfcjson.NotifySpentCmd{
				OutPoints: []pfcjson.OutPoint{{Hash: "123", Index: 0}},
			},
		},
		{
			name: "stopnotifyspent",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("stopnotifyspent", `[{"hash":"123","index":0}]`)
			},
			staticCmd: func() interface{} {
				ops := []pfcjson.OutPoint{{Hash: "123", Index: 0}}
				return pfcjson.NewStopNotifySpentCmd(ops)
			},
			marshalled: `{"jsonrpc":"1.0","method":"stopnotifyspent","params":[[{"hash":"123","index":0}]],"id":1}`,
			unmarshalled: &pfcjson.StopNotifySpentCmd{
				OutPoints: []pfcjson.OutPoint{{Hash: "123", Index: 0}},
			},
		},
		{
			name: "rescan",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("rescan", "123", `["1Address"]`, `[{"hash":"0000000000000000000000000000000000000000000000000000000000000123","index":0}]`)
			},
			staticCmd: func() interface{} {
				addrs := []string{"1Address"}
				ops := []pfcjson.OutPoint{{
					Hash:  "0000000000000000000000000000000000000000000000000000000000000123",
					Index: 0,
				}}
				return pfcjson.NewRescanCmd("123", addrs, ops, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"rescan","params":["123",["1Address"],[{"hash":"0000000000000000000000000000000000000000000000000000000000000123","index":0}]],"id":1}`,
			unmarshalled: &pfcjson.RescanCmd{
				BeginBlock: "123",
				Addresses:  []string{"1Address"},
				OutPoints:  []pfcjson.OutPoint{{Hash: "0000000000000000000000000000000000000000000000000000000000000123", Index: 0}},
				EndBlock:   nil,
			},
		},
		{
			name: "rescan optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("rescan", "123", `["1Address"]`, `[{"hash":"123","index":0}]`, "456")
			},
			staticCmd: func() interface{} {
				addrs := []string{"1Address"}
				ops := []pfcjson.OutPoint{{Hash: "123", Index: 0}}
				return pfcjson.NewRescanCmd("123", addrs, ops, pfcjson.String("456"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"rescan","params":["123",["1Address"],[{"hash":"123","index":0}],"456"],"id":1}`,
			unmarshalled: &pfcjson.RescanCmd{
				BeginBlock: "123",
				Addresses:  []string{"1Address"},
				OutPoints:  []pfcjson.OutPoint{{Hash: "123", Index: 0}},
				EndBlock:   pfcjson.String("456"),
			},
		},
		{
			name: "loadtxfilter",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("loadtxfilter", false, `["1Address"]`, `[{"hash":"0000000000000000000000000000000000000000000000000000000000000123","index":0}]`)
			},
			staticCmd: func() interface{} {
				addrs := []string{"1Address"}
				ops := []pfcjson.OutPoint{{
					Hash:  "0000000000000000000000000000000000000000000000000000000000000123",
					Index: 0,
				}}
				return pfcjson.NewLoadTxFilterCmd(false, addrs, ops)
			},
			marshalled: `{"jsonrpc":"1.0","method":"loadtxfilter","params":[false,["1Address"],[{"hash":"0000000000000000000000000000000000000000000000000000000000000123","index":0}]],"id":1}`,
			unmarshalled: &pfcjson.LoadTxFilterCmd{
				Reload:    false,
				Addresses: []string{"1Address"},
				OutPoints: []pfcjson.OutPoint{{Hash: "0000000000000000000000000000000000000000000000000000000000000123", Index: 0}},
			},
		},
		{
			name: "rescanblocks",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("rescanblocks", `["0000000000000000000000000000000000000000000000000000000000000123"]`)
			},
			staticCmd: func() interface{} {
				blockhashes := []string{"0000000000000000000000000000000000000000000000000000000000000123"}
				return pfcjson.NewRescanBlocksCmd(blockhashes)
			},
			marshalled: `{"jsonrpc":"1.0","method":"rescanblocks","params":[["0000000000000000000000000000000000000000000000000000000000000123"]],"id":1}`,
			unmarshalled: &pfcjson.RescanBlocksCmd{
				BlockHashes: []string{"0000000000000000000000000000000000000000000000000000000000000123"},
			},
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Marshal the command as created by the new static command
		// creation function.
		marshalled, err := pfcjson.MarshalCmd(testID, test.staticCmd())
		if err != nil {
			t.Errorf("MarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !bytes.Equal(marshalled, []byte(test.marshalled)) {
			t.Errorf("Test #%d (%s) unexpected marshalled data - "+
				"got %s, want %s", i, test.name, marshalled,
				test.marshalled)
			continue
		}

		// Ensure the command is created without error via the generic
		// new command creation function.
		cmd, err := test.newCmd()
		if err != nil {
			t.Errorf("Test #%d (%s) unexpected NewCmd error: %v ",
				i, test.name, err)
		}

		// Marshal the command as created by the generic new command
		// creation function.
		marshalled, err = pfcjson.MarshalCmd(testID, cmd)
		if err != nil {
			t.Errorf("MarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !bytes.Equal(marshalled, []byte(test.marshalled)) {
			t.Errorf("Test #%d (%s) unexpected marshalled data - "+
				"got %s, want %s", i, test.name, marshalled,
				test.marshalled)
			continue
		}

		var request pfcjson.Request
		if err := json.Unmarshal(marshalled, &request); err != nil {
			t.Errorf("Test #%d (%s) unexpected error while "+
				"unmarshalling JSON-RPC request: %v", i,
				test.name, err)
			continue
		}

		cmd, err = pfcjson.UnmarshalCmd(&request)
		if err != nil {
			t.Errorf("UnmarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !reflect.DeepEqual(cmd, test.unmarshalled) {
			t.Errorf("Test #%d (%s) unexpected unmarshalled command "+
				"- got %s, want %s", i, test.name,
				fmt.Sprintf("(%T) %+[1]v", cmd),
				fmt.Sprintf("(%T) %+[1]v\n", test.unmarshalled))
			continue
		}
	}
}
