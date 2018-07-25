// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2015-2018 The Decred developers
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

// TestWalletSvrCmds tests all of the wallet server commands marshal and
// unmarshal into valid results include handling of optional fields being
// omitted in the marshalled command, while optional fields with defaults have
// the default assigned on unmarshalled commands.
func TestWalletSvrCmds(t *testing.T) {
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
			name: "addmultisigaddress",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("addmultisigaddress", 2, []string{"031234", "035678"})
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return pfcjson.NewAddMultisigAddressCmd(2, keys, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"addmultisigaddress","params":[2,["031234","035678"]],"id":1}`,
			unmarshalled: &pfcjson.AddMultisigAddressCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
				Account:   nil,
			},
		},
		{
			name: "addmultisigaddress optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("addmultisigaddress", 2, []string{"031234", "035678"}, "test")
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return pfcjson.NewAddMultisigAddressCmd(2, keys, pfcjson.String("test"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"addmultisigaddress","params":[2,["031234","035678"],"test"],"id":1}`,
			unmarshalled: &pfcjson.AddMultisigAddressCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
				Account:   pfcjson.String("test"),
			},
		},
		{
			name: "createmultisig",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("createmultisig", 2, []string{"031234", "035678"})
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return pfcjson.NewCreateMultisigCmd(2, keys)
			},
			marshalled: `{"jsonrpc":"1.0","method":"createmultisig","params":[2,["031234","035678"]],"id":1}`,
			unmarshalled: &pfcjson.CreateMultisigCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
			},
		},
		{
			name: "createnewaccount",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("createnewaccount", "acct")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewCreateNewAccountCmd("acct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"createnewaccount","params":["acct"],"id":1}`,
			unmarshalled: &pfcjson.CreateNewAccountCmd{
				Account: "acct",
			},
		},
		{
			name: "dumpprivkey",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("dumpprivkey", "1Address")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewDumpPrivKeyCmd("1Address")
			},
			marshalled: `{"jsonrpc":"1.0","method":"dumpprivkey","params":["1Address"],"id":1}`,
			unmarshalled: &pfcjson.DumpPrivKeyCmd{
				Address: "1Address",
			},
		},
		{
			name: "estimatefee",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("estimatefee", 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewEstimateFeeCmd(6)
			},
			marshalled: `{"jsonrpc":"1.0","method":"estimatefee","params":[6],"id":1}`,
			unmarshalled: &pfcjson.EstimateFeeCmd{
				NumBlocks: 6,
			},
		},
		{
			name: "estimatepriority",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("estimatepriority", 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewEstimatePriorityCmd(6)
			},
			marshalled: `{"jsonrpc":"1.0","method":"estimatepriority","params":[6],"id":1}`,
			unmarshalled: &pfcjson.EstimatePriorityCmd{
				NumBlocks: 6,
			},
		},
		{
			name: "getaccount",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getaccount", "1Address")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetAccountCmd("1Address")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getaccount","params":["1Address"],"id":1}`,
			unmarshalled: &pfcjson.GetAccountCmd{
				Address: "1Address",
			},
		},
		{
			name: "getaccountaddress",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getaccountaddress", "acct")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetAccountAddressCmd("acct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getaccountaddress","params":["acct"],"id":1}`,
			unmarshalled: &pfcjson.GetAccountAddressCmd{
				Account: "acct",
			},
		},
		{
			name: "getaddressesbyaccount",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getaddressesbyaccount", "acct")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetAddressesByAccountCmd("acct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getaddressesbyaccount","params":["acct"],"id":1}`,
			unmarshalled: &pfcjson.GetAddressesByAccountCmd{
				Account: "acct",
			},
		},
		{
			name: "getbalance",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getbalance")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBalanceCmd(nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetBalanceCmd{
				Account: nil,
				MinConf: pfcjson.Int(1),
			},
		},
		{
			name: "getbalance optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getbalance", "acct")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBalanceCmd(pfcjson.String("acct"), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":["acct"],"id":1}`,
			unmarshalled: &pfcjson.GetBalanceCmd{
				Account: pfcjson.String("acct"),
				MinConf: pfcjson.Int(1),
			},
		},
		{
			name: "getbalance optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getbalance", "acct", 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBalanceCmd(pfcjson.String("acct"), pfcjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":["acct",6],"id":1}`,
			unmarshalled: &pfcjson.GetBalanceCmd{
				Account: pfcjson.String("acct"),
				MinConf: pfcjson.Int(6),
			},
		},
		{
			name: "getnewaddress",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getnewaddress")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetNewAddressCmd(nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getnewaddress","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetNewAddressCmd{
				Account:   nil,
				GapPolicy: nil,
			},
		},
		{
			name: "getnewaddress optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getnewaddress", "acct", "ignore")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetNewAddressCmd(pfcjson.String("acct"), pfcjson.String("ignore"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getnewaddress","params":["acct","ignore"],"id":1}`,
			unmarshalled: &pfcjson.GetNewAddressCmd{
				Account:   pfcjson.String("acct"),
				GapPolicy: pfcjson.String("ignore"),
			},
		},
		{
			name: "getrawchangeaddress",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getrawchangeaddress")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetRawChangeAddressCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawchangeaddress","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetRawChangeAddressCmd{
				Account: nil,
			},
		},
		{
			name: "getrawchangeaddress optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getrawchangeaddress", "acct")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetRawChangeAddressCmd(pfcjson.String("acct"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawchangeaddress","params":["acct"],"id":1}`,
			unmarshalled: &pfcjson.GetRawChangeAddressCmd{
				Account: pfcjson.String("acct"),
			},
		},
		{
			name: "getreceivedbyaccount",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getreceivedbyaccount", "acct")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetReceivedByAccountCmd("acct", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaccount","params":["acct"],"id":1}`,
			unmarshalled: &pfcjson.GetReceivedByAccountCmd{
				Account: "acct",
				MinConf: pfcjson.Int(1),
			},
		},
		{
			name: "getreceivedbyaccount optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getreceivedbyaccount", "acct", 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetReceivedByAccountCmd("acct", pfcjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaccount","params":["acct",6],"id":1}`,
			unmarshalled: &pfcjson.GetReceivedByAccountCmd{
				Account: "acct",
				MinConf: pfcjson.Int(6),
			},
		},
		{
			name: "getreceivedbyaddress",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getreceivedbyaddress", "1Address")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetReceivedByAddressCmd("1Address", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaddress","params":["1Address"],"id":1}`,
			unmarshalled: &pfcjson.GetReceivedByAddressCmd{
				Address: "1Address",
				MinConf: pfcjson.Int(1),
			},
		},
		{
			name: "getreceivedbyaddress optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getreceivedbyaddress", "1Address", 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetReceivedByAddressCmd("1Address", pfcjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaddress","params":["1Address",6],"id":1}`,
			unmarshalled: &pfcjson.GetReceivedByAddressCmd{
				Address: "1Address",
				MinConf: pfcjson.Int(6),
			},
		},
		{
			name: "gettransaction",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("gettransaction", "123")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetTransactionCmd("123", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"gettransaction","params":["123"],"id":1}`,
			unmarshalled: &pfcjson.GetTransactionCmd{
				Txid:             "123",
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "gettransaction optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("gettransaction", "123", true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetTransactionCmd("123", pfcjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"gettransaction","params":["123",true],"id":1}`,
			unmarshalled: &pfcjson.GetTransactionCmd{
				Txid:             "123",
				IncludeWatchOnly: pfcjson.Bool(true),
			},
		},
		{
			name: "importaddress",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("importaddress", "1Address")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewImportAddressCmd("1Address", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importaddress","params":["1Address"],"id":1}`,
			unmarshalled: &pfcjson.ImportAddressCmd{
				Address: "1Address",
				Rescan:  pfcjson.Bool(true),
			},
		},
		{
			name: "importaddress optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("importaddress", "1Address", false)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewImportAddressCmd("1Address", pfcjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"importaddress","params":["1Address",false],"id":1}`,
			unmarshalled: &pfcjson.ImportAddressCmd{
				Address: "1Address",
				Rescan:  pfcjson.Bool(false),
			},
		},
		{
			name: "importprivkey",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("importprivkey", "abc")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewImportPrivKeyCmd("abc", nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc"],"id":1}`,
			unmarshalled: &pfcjson.ImportPrivKeyCmd{
				PrivKey: "abc",
				Label:   nil,
				Rescan:  pfcjson.Bool(true),
			},
		},
		{
			name: "importprivkey optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("importprivkey", "abc", "label")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewImportPrivKeyCmd("abc", pfcjson.String("label"), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc","label"],"id":1}`,
			unmarshalled: &pfcjson.ImportPrivKeyCmd{
				PrivKey: "abc",
				Label:   pfcjson.String("label"),
				Rescan:  pfcjson.Bool(true),
			},
		},
		{
			name: "importprivkey optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("importprivkey", "abc", "label", false)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewImportPrivKeyCmd("abc", pfcjson.String("label"), pfcjson.Bool(false), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc","label",false],"id":1}`,
			unmarshalled: &pfcjson.ImportPrivKeyCmd{
				PrivKey: "abc",
				Label:   pfcjson.String("label"),
				Rescan:  pfcjson.Bool(false),
			},
		},
		{
			name: "importprivkey optional3",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("importprivkey", "abc", "label", false, 12345)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewImportPrivKeyCmd("abc", pfcjson.String("label"), pfcjson.Bool(false), pfcjson.Int(12345))
			},
			marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc","label",false,12345],"id":1}`,
			unmarshalled: &pfcjson.ImportPrivKeyCmd{
				PrivKey:  "abc",
				Label:    pfcjson.String("label"),
				Rescan:   pfcjson.Bool(false),
				ScanFrom: pfcjson.Int(12345),
			},
		},
		{
			name: "importpubkey",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("importpubkey", "031234")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewImportPubKeyCmd("031234", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importpubkey","params":["031234"],"id":1}`,
			unmarshalled: &pfcjson.ImportPubKeyCmd{
				PubKey: "031234",
				Rescan: pfcjson.Bool(true),
			},
		},
		{
			name: "importpubkey optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("importpubkey", "031234", false)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewImportPubKeyCmd("031234", pfcjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"importpubkey","params":["031234",false],"id":1}`,
			unmarshalled: &pfcjson.ImportPubKeyCmd{
				PubKey: "031234",
				Rescan: pfcjson.Bool(false),
			},
		},
		{
			name: "keypoolrefill",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("keypoolrefill")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewKeyPoolRefillCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"keypoolrefill","params":[],"id":1}`,
			unmarshalled: &pfcjson.KeyPoolRefillCmd{
				NewSize: pfcjson.Uint(100),
			},
		},
		{
			name: "keypoolrefill optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("keypoolrefill", 200)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewKeyPoolRefillCmd(pfcjson.Uint(200))
			},
			marshalled: `{"jsonrpc":"1.0","method":"keypoolrefill","params":[200],"id":1}`,
			unmarshalled: &pfcjson.KeyPoolRefillCmd{
				NewSize: pfcjson.Uint(200),
			},
		},
		{
			name: "listaccounts",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listaccounts")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListAccountsCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listaccounts","params":[],"id":1}`,
			unmarshalled: &pfcjson.ListAccountsCmd{
				MinConf: pfcjson.Int(1),
			},
		},
		{
			name: "listaccounts optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listaccounts", 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListAccountsCmd(pfcjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listaccounts","params":[6],"id":1}`,
			unmarshalled: &pfcjson.ListAccountsCmd{
				MinConf: pfcjson.Int(6),
			},
		},
		{
			name: "listlockunspent",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listlockunspent")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListLockUnspentCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"listlockunspent","params":[],"id":1}`,
			unmarshalled: &pfcjson.ListLockUnspentCmd{},
		},
		{
			name: "listreceivedbyaccount",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listreceivedbyaccount")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListReceivedByAccountCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[],"id":1}`,
			unmarshalled: &pfcjson.ListReceivedByAccountCmd{
				MinConf:          pfcjson.Int(1),
				IncludeEmpty:     pfcjson.Bool(false),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaccount optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listreceivedbyaccount", 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListReceivedByAccountCmd(pfcjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6],"id":1}`,
			unmarshalled: &pfcjson.ListReceivedByAccountCmd{
				MinConf:          pfcjson.Int(6),
				IncludeEmpty:     pfcjson.Bool(false),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaccount optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listreceivedbyaccount", 6, true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListReceivedByAccountCmd(pfcjson.Int(6), pfcjson.Bool(true), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6,true],"id":1}`,
			unmarshalled: &pfcjson.ListReceivedByAccountCmd{
				MinConf:          pfcjson.Int(6),
				IncludeEmpty:     pfcjson.Bool(true),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaccount optional3",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listreceivedbyaccount", 6, true, false)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListReceivedByAccountCmd(pfcjson.Int(6), pfcjson.Bool(true), pfcjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6,true,false],"id":1}`,
			unmarshalled: &pfcjson.ListReceivedByAccountCmd{
				MinConf:          pfcjson.Int(6),
				IncludeEmpty:     pfcjson.Bool(true),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listreceivedbyaddress")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListReceivedByAddressCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[],"id":1}`,
			unmarshalled: &pfcjson.ListReceivedByAddressCmd{
				MinConf:          pfcjson.Int(1),
				IncludeEmpty:     pfcjson.Bool(false),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listreceivedbyaddress", 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListReceivedByAddressCmd(pfcjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6],"id":1}`,
			unmarshalled: &pfcjson.ListReceivedByAddressCmd{
				MinConf:          pfcjson.Int(6),
				IncludeEmpty:     pfcjson.Bool(false),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listreceivedbyaddress", 6, true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListReceivedByAddressCmd(pfcjson.Int(6), pfcjson.Bool(true), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6,true],"id":1}`,
			unmarshalled: &pfcjson.ListReceivedByAddressCmd{
				MinConf:          pfcjson.Int(6),
				IncludeEmpty:     pfcjson.Bool(true),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress optional3",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listreceivedbyaddress", 6, true, false)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListReceivedByAddressCmd(pfcjson.Int(6), pfcjson.Bool(true), pfcjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6,true,false],"id":1}`,
			unmarshalled: &pfcjson.ListReceivedByAddressCmd{
				MinConf:          pfcjson.Int(6),
				IncludeEmpty:     pfcjson.Bool(true),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listsinceblock",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listsinceblock")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListSinceBlockCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":[],"id":1}`,
			unmarshalled: &pfcjson.ListSinceBlockCmd{
				BlockHash:           nil,
				TargetConfirmations: pfcjson.Int(1),
				IncludeWatchOnly:    pfcjson.Bool(false),
			},
		},
		{
			name: "listsinceblock optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listsinceblock", "123")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListSinceBlockCmd(pfcjson.String("123"), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123"],"id":1}`,
			unmarshalled: &pfcjson.ListSinceBlockCmd{
				BlockHash:           pfcjson.String("123"),
				TargetConfirmations: pfcjson.Int(1),
				IncludeWatchOnly:    pfcjson.Bool(false),
			},
		},
		{
			name: "listsinceblock optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listsinceblock", "123", 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListSinceBlockCmd(pfcjson.String("123"), pfcjson.Int(6), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123",6],"id":1}`,
			unmarshalled: &pfcjson.ListSinceBlockCmd{
				BlockHash:           pfcjson.String("123"),
				TargetConfirmations: pfcjson.Int(6),
				IncludeWatchOnly:    pfcjson.Bool(false),
			},
		},
		{
			name: "listsinceblock optional3",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listsinceblock", "123", 6, true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListSinceBlockCmd(pfcjson.String("123"), pfcjson.Int(6), pfcjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123",6,true],"id":1}`,
			unmarshalled: &pfcjson.ListSinceBlockCmd{
				BlockHash:           pfcjson.String("123"),
				TargetConfirmations: pfcjson.Int(6),
				IncludeWatchOnly:    pfcjson.Bool(true),
			},
		},
		{
			name: "listtransactions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listtransactions")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListTransactionsCmd(nil, nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":[],"id":1}`,
			unmarshalled: &pfcjson.ListTransactionsCmd{
				Account:          nil,
				Count:            pfcjson.Int(10),
				From:             pfcjson.Int(0),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listtransactions", "acct")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListTransactionsCmd(pfcjson.String("acct"), nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct"],"id":1}`,
			unmarshalled: &pfcjson.ListTransactionsCmd{
				Account:          pfcjson.String("acct"),
				Count:            pfcjson.Int(10),
				From:             pfcjson.Int(0),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listtransactions", "acct", 20)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListTransactionsCmd(pfcjson.String("acct"), pfcjson.Int(20), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20],"id":1}`,
			unmarshalled: &pfcjson.ListTransactionsCmd{
				Account:          pfcjson.String("acct"),
				Count:            pfcjson.Int(20),
				From:             pfcjson.Int(0),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional3",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listtransactions", "acct", 20, 1)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListTransactionsCmd(pfcjson.String("acct"), pfcjson.Int(20),
					pfcjson.Int(1), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20,1],"id":1}`,
			unmarshalled: &pfcjson.ListTransactionsCmd{
				Account:          pfcjson.String("acct"),
				Count:            pfcjson.Int(20),
				From:             pfcjson.Int(1),
				IncludeWatchOnly: pfcjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional4",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listtransactions", "acct", 20, 1, true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListTransactionsCmd(pfcjson.String("acct"), pfcjson.Int(20),
					pfcjson.Int(1), pfcjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20,1,true],"id":1}`,
			unmarshalled: &pfcjson.ListTransactionsCmd{
				Account:          pfcjson.String("acct"),
				Count:            pfcjson.Int(20),
				From:             pfcjson.Int(1),
				IncludeWatchOnly: pfcjson.Bool(true),
			},
		},
		{
			name: "listunspent",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listunspent")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListUnspentCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[],"id":1}`,
			unmarshalled: &pfcjson.ListUnspentCmd{
				MinConf:   pfcjson.Int(1),
				MaxConf:   pfcjson.Int(9999999),
				Addresses: nil,
			},
		},
		{
			name: "listunspent optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listunspent", 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListUnspentCmd(pfcjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6],"id":1}`,
			unmarshalled: &pfcjson.ListUnspentCmd{
				MinConf:   pfcjson.Int(6),
				MaxConf:   pfcjson.Int(9999999),
				Addresses: nil,
			},
		},
		{
			name: "listunspent optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listunspent", 6, 100)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListUnspentCmd(pfcjson.Int(6), pfcjson.Int(100), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6,100],"id":1}`,
			unmarshalled: &pfcjson.ListUnspentCmd{
				MinConf:   pfcjson.Int(6),
				MaxConf:   pfcjson.Int(100),
				Addresses: nil,
			},
		},
		{
			name: "listunspent optional3",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("listunspent", 6, 100, []string{"1Address", "1Address2"})
			},
			staticCmd: func() interface{} {
				return pfcjson.NewListUnspentCmd(pfcjson.Int(6), pfcjson.Int(100),
					&[]string{"1Address", "1Address2"})
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6,100,["1Address","1Address2"]],"id":1}`,
			unmarshalled: &pfcjson.ListUnspentCmd{
				MinConf:   pfcjson.Int(6),
				MaxConf:   pfcjson.Int(100),
				Addresses: &[]string{"1Address", "1Address2"},
			},
		},
		{
			name: "lockunspent",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("lockunspent", true, `[{"txid":"123","vout":1}]`)
			},
			staticCmd: func() interface{} {
				txInputs := []pfcjson.TransactionInput{
					{Txid: "123", Vout: 1},
				}
				return pfcjson.NewLockUnspentCmd(true, txInputs)
			},
			marshalled: `{"jsonrpc":"1.0","method":"lockunspent","params":[true,[{"txid":"123","vout":1,"tree":0}]],"id":1}`,
			unmarshalled: &pfcjson.LockUnspentCmd{
				Unlock: true,
				Transactions: []pfcjson.TransactionInput{
					{Txid: "123", Vout: 1},
				},
			},
		},
		{
			name: "renameaccount",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("renameaccount", "oldacct", "newacct")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewRenameAccountCmd("oldacct", "newacct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"renameaccount","params":["oldacct","newacct"],"id":1}`,
			unmarshalled: &pfcjson.RenameAccountCmd{
				OldAccount: "oldacct",
				NewAccount: "newacct",
			},
		},
		{
			name: "sendfrom",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendfrom", "from", "1Address", 0.5)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSendFromCmd("from", "1Address", 0.5, nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5],"id":1}`,
			unmarshalled: &pfcjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     pfcjson.Int(1),
				Comment:     nil,
				CommentTo:   nil,
			},
		},
		{
			name: "sendfrom optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSendFromCmd("from", "1Address", 0.5, pfcjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6],"id":1}`,
			unmarshalled: &pfcjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     pfcjson.Int(6),
				Comment:     nil,
				CommentTo:   nil,
			},
		},
		{
			name: "sendfrom optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6, "comment")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSendFromCmd("from", "1Address", 0.5, pfcjson.Int(6),
					pfcjson.String("comment"), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6,"comment"],"id":1}`,
			unmarshalled: &pfcjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     pfcjson.Int(6),
				Comment:     pfcjson.String("comment"),
				CommentTo:   nil,
			},
		},
		{
			name: "sendfrom optional3",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6, "comment", "commentto")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSendFromCmd("from", "1Address", 0.5, pfcjson.Int(6),
					pfcjson.String("comment"), pfcjson.String("commentto"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6,"comment","commentto"],"id":1}`,
			unmarshalled: &pfcjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     pfcjson.Int(6),
				Comment:     pfcjson.String("comment"),
				CommentTo:   pfcjson.String("commentto"),
			},
		},
		{
			name: "sendmany",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendmany", "from", `{"1Address":0.5}`)
			},
			staticCmd: func() interface{} {
				amounts := map[string]float64{"1Address": 0.5}
				return pfcjson.NewSendManyCmd("from", amounts, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5}],"id":1}`,
			unmarshalled: &pfcjson.SendManyCmd{
				FromAccount: "from",
				Amounts:     map[string]float64{"1Address": 0.5},
				MinConf:     pfcjson.Int(1),
				Comment:     nil,
			},
		},
		{
			name: "sendmany optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendmany", "from", `{"1Address":0.5}`, 6)
			},
			staticCmd: func() interface{} {
				amounts := map[string]float64{"1Address": 0.5}
				return pfcjson.NewSendManyCmd("from", amounts, pfcjson.Int(6), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5},6],"id":1}`,
			unmarshalled: &pfcjson.SendManyCmd{
				FromAccount: "from",
				Amounts:     map[string]float64{"1Address": 0.5},
				MinConf:     pfcjson.Int(6),
				Comment:     nil,
			},
		},
		{
			name: "sendmany optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendmany", "from", `{"1Address":0.5}`, 6, "comment")
			},
			staticCmd: func() interface{} {
				amounts := map[string]float64{"1Address": 0.5}
				return pfcjson.NewSendManyCmd("from", amounts, pfcjson.Int(6), pfcjson.String("comment"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5},6,"comment"],"id":1}`,
			unmarshalled: &pfcjson.SendManyCmd{
				FromAccount: "from",
				Amounts:     map[string]float64{"1Address": 0.5},
				MinConf:     pfcjson.Int(6),
				Comment:     pfcjson.String("comment"),
			},
		},
		{
			name: "sendtoaddress",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendtoaddress", "1Address", 0.5)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSendToAddressCmd("1Address", 0.5, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendtoaddress","params":["1Address",0.5],"id":1}`,
			unmarshalled: &pfcjson.SendToAddressCmd{
				Address:   "1Address",
				Amount:    0.5,
				Comment:   nil,
				CommentTo: nil,
			},
		},
		{
			name: "sendtoaddress optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendtoaddress", "1Address", 0.5, "comment", "commentto")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSendToAddressCmd("1Address", 0.5, pfcjson.String("comment"),
					pfcjson.String("commentto"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendtoaddress","params":["1Address",0.5,"comment","commentto"],"id":1}`,
			unmarshalled: &pfcjson.SendToAddressCmd{
				Address:   "1Address",
				Amount:    0.5,
				Comment:   pfcjson.String("comment"),
				CommentTo: pfcjson.String("commentto"),
			},
		},
		{
			name: "settxfee",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("settxfee", 0.0001)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSetTxFeeCmd(0.0001)
			},
			marshalled: `{"jsonrpc":"1.0","method":"settxfee","params":[0.0001],"id":1}`,
			unmarshalled: &pfcjson.SetTxFeeCmd{
				Amount: 0.0001,
			},
		},
		{
			name: "signmessage",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("signmessage", "1Address", "message")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSignMessageCmd("1Address", "message")
			},
			marshalled: `{"jsonrpc":"1.0","method":"signmessage","params":["1Address","message"],"id":1}`,
			unmarshalled: &pfcjson.SignMessageCmd{
				Address: "1Address",
				Message: "message",
			},
		},
		{
			name: "signrawtransaction",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("signrawtransaction", "001122")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSignRawTransactionCmd("001122", nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122"],"id":1}`,
			unmarshalled: &pfcjson.SignRawTransactionCmd{
				RawTx:    "001122",
				Inputs:   nil,
				PrivKeys: nil,
				Flags:    pfcjson.String("ALL"),
			},
		},
		{
			name: "signrawtransaction optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("signrawtransaction", "001122", `[{"txid":"123","vout":1,"tree":0,"scriptPubKey":"00","redeemScript":"01"}]`)
			},
			staticCmd: func() interface{} {
				txInputs := []pfcjson.RawTxInput{
					{
						Txid:         "123",
						Vout:         1,
						ScriptPubKey: "00",
						RedeemScript: "01",
					},
				}

				return pfcjson.NewSignRawTransactionCmd("001122", &txInputs, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[{"txid":"123","vout":1,"tree":0,"scriptPubKey":"00","redeemScript":"01"}]],"id":1}`,
			unmarshalled: &pfcjson.SignRawTransactionCmd{
				RawTx: "001122",
				Inputs: &[]pfcjson.RawTxInput{
					{
						Txid:         "123",
						Vout:         1,
						ScriptPubKey: "00",
						RedeemScript: "01",
					},
				},
				PrivKeys: nil,
				Flags:    pfcjson.String("ALL"),
			},
		},
		{
			name: "signrawtransaction optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("signrawtransaction", "001122", `[]`, `["abc"]`)
			},
			staticCmd: func() interface{} {
				txInputs := []pfcjson.RawTxInput{}
				privKeys := []string{"abc"}
				return pfcjson.NewSignRawTransactionCmd("001122", &txInputs, &privKeys, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[],["abc"]],"id":1}`,
			unmarshalled: &pfcjson.SignRawTransactionCmd{
				RawTx:    "001122",
				Inputs:   &[]pfcjson.RawTxInput{},
				PrivKeys: &[]string{"abc"},
				Flags:    pfcjson.String("ALL"),
			},
		},
		{
			name: "signrawtransaction optional3",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("signrawtransaction", "001122", `[]`, `[]`, "ALL")
			},
			staticCmd: func() interface{} {
				txInputs := []pfcjson.RawTxInput{}
				privKeys := []string{}
				return pfcjson.NewSignRawTransactionCmd("001122", &txInputs, &privKeys,
					pfcjson.String("ALL"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[],[],"ALL"],"id":1}`,
			unmarshalled: &pfcjson.SignRawTransactionCmd{
				RawTx:    "001122",
				Inputs:   &[]pfcjson.RawTxInput{},
				PrivKeys: &[]string{},
				Flags:    pfcjson.String("ALL"),
			},
		},
		{
			name: "sweepaccount - optionals provided",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sweepaccount", "default", "DsUZxxoHJSty8DCfwfartwTYbuhmVct7tJu", 6, 0.05)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSweepAccountCmd("default", "DsUZxxoHJSty8DCfwfartwTYbuhmVct7tJu",
					func(i uint32) *uint32 { return &i }(6),
					func(i float64) *float64 { return &i }(0.05))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sweepaccount","params":["default","DsUZxxoHJSty8DCfwfartwTYbuhmVct7tJu",6,0.05],"id":1}`,
			unmarshalled: &pfcjson.SweepAccountCmd{
				SourceAccount:         "default",
				DestinationAddress:    "DsUZxxoHJSty8DCfwfartwTYbuhmVct7tJu",
				RequiredConfirmations: func(i uint32) *uint32 { return &i }(6),
				FeePerKb:              func(i float64) *float64 { return &i }(0.05),
			},
		},
		{
			name: "sweepaccount - optionals omitted",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sweepaccount", "default", "DsUZxxoHJSty8DCfwfartwTYbuhmVct7tJu")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSweepAccountCmd("default", "DsUZxxoHJSty8DCfwfartwTYbuhmVct7tJu", nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sweepaccount","params":["default","DsUZxxoHJSty8DCfwfartwTYbuhmVct7tJu"],"id":1}`,
			unmarshalled: &pfcjson.SweepAccountCmd{
				SourceAccount:      "default",
				DestinationAddress: "DsUZxxoHJSty8DCfwfartwTYbuhmVct7tJu",
			},
		},
		{
			name: "verifyseed",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("verifyseed", "abc")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewVerifySeedCmd("abc", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"verifyseed","params":["abc"],"id":1}`,
			unmarshalled: &pfcjson.VerifySeedCmd{
				Seed:    "abc",
				Account: nil,
			},
		},
		{
			name: "verifyseed optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("verifyseed", "abc", 5)
			},
			staticCmd: func() interface{} {
				account := pfcjson.Uint32(5)
				return pfcjson.NewVerifySeedCmd("abc", account)
			},
			marshalled: `{"jsonrpc":"1.0","method":"verifyseed","params":["abc",5],"id":1}`,
			unmarshalled: &pfcjson.VerifySeedCmd{
				Seed:    "abc",
				Account: pfcjson.Uint32(5),
			},
		},
		{
			name: "walletlock",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("walletlock")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewWalletLockCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"walletlock","params":[],"id":1}`,
			unmarshalled: &pfcjson.WalletLockCmd{},
		},
		{
			name: "walletpassphrase",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("walletpassphrase", "pass", 60)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewWalletPassphraseCmd("pass", 60)
			},
			marshalled: `{"jsonrpc":"1.0","method":"walletpassphrase","params":["pass",60],"id":1}`,
			unmarshalled: &pfcjson.WalletPassphraseCmd{
				Passphrase: "pass",
				Timeout:    60,
			},
		},
		{
			name: "walletpassphrasechange",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("walletpassphrasechange", "old", "new")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewWalletPassphraseChangeCmd("old", "new")
			},
			marshalled: `{"jsonrpc":"1.0","method":"walletpassphrasechange","params":["old","new"],"id":1}`,
			unmarshalled: &pfcjson.WalletPassphraseChangeCmd{
				OldPassphrase: "old",
				NewPassphrase: "new",
			},
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Marshal the command as created by the new static command
		// creation function.
		marshalled, err := pfcjson.MarshalCmd("1.0", testID, test.staticCmd())
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
		marshalled, err = pfcjson.MarshalCmd("1.0", testID, cmd)
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
