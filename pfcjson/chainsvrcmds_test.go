// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2016 The Decred developers
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

// TestChainSvrCmds tests all of the chain server commands marshal and unmarshal
// into valid results include handling of optional fields being omitted in the
// marshalled command, while optional fields with defaults have the default
// assigned on unmarshalled commands.
func TestChainSvrCmds(t *testing.T) {
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
			name: "addnode",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("addnode", "127.0.0.1", pfcjson.ANRemove)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewAddNodeCmd("127.0.0.1", pfcjson.ANRemove)
			},
			marshalled:   `{"jsonrpc":"1.0","method":"addnode","params":["127.0.0.1","remove"],"id":1}`,
			unmarshalled: &pfcjson.AddNodeCmd{Addr: "127.0.0.1", SubCmd: pfcjson.ANRemove},
		},
		{
			name: "createrawtransaction",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("createrawtransaction", `[{"txid":"123","vout":1}]`,
					`{"456":0.0123}`)
			},
			staticCmd: func() interface{} {
				txInputs := []pfcjson.TransactionInput{
					{Txid: "123", Vout: 1},
				}
				amounts := map[string]float64{"456": .0123}
				return pfcjson.NewCreateRawTransactionCmd(txInputs, amounts, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"createrawtransaction","params":[[{"txid":"123","vout":1,"tree":0}],{"456":0.0123}],"id":1}`,
			unmarshalled: &pfcjson.CreateRawTransactionCmd{
				Inputs:  []pfcjson.TransactionInput{{Txid: "123", Vout: 1}},
				Amounts: map[string]float64{"456": .0123},
			},
		},
		{
			name: "createrawtransaction optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("createrawtransaction", `[{"txid":"123","vout":1,"tree":0}]`,
					`{"456":0.0123}`, int64(12312333333), int64(12312333333))
			},
			staticCmd: func() interface{} {
				txInputs := []pfcjson.TransactionInput{
					{Txid: "123", Vout: 1},
				}
				amounts := map[string]float64{"456": .0123}
				return pfcjson.NewCreateRawTransactionCmd(txInputs, amounts, pfcjson.Int64(12312333333), pfcjson.Int64(12312333333))
			},
			marshalled: `{"jsonrpc":"1.0","method":"createrawtransaction","params":[[{"txid":"123","vout":1,"tree":0}],{"456":0.0123},12312333333,12312333333],"id":1}`,
			unmarshalled: &pfcjson.CreateRawTransactionCmd{
				Inputs:   []pfcjson.TransactionInput{{Txid: "123", Vout: 1}},
				Amounts:  map[string]float64{"456": .0123},
				LockTime: pfcjson.Int64(12312333333),
				Expiry:   pfcjson.Int64(12312333333),
			},
		},
		{
			name: "debuglevel",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("debuglevel", "trace")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewDebugLevelCmd("trace")
			},
			marshalled: `{"jsonrpc":"1.0","method":"debuglevel","params":["trace"],"id":1}`,
			unmarshalled: &pfcjson.DebugLevelCmd{
				LevelSpec: "trace",
			},
		},
		{
			name: "decoderawtransaction",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("decoderawtransaction", "123")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewDecodeRawTransactionCmd("123")
			},
			marshalled:   `{"jsonrpc":"1.0","method":"decoderawtransaction","params":["123"],"id":1}`,
			unmarshalled: &pfcjson.DecodeRawTransactionCmd{HexTx: "123"},
		},
		{
			name: "decodescript",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("decodescript", "00")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewDecodeScriptCmd("00")
			},
			marshalled:   `{"jsonrpc":"1.0","method":"decodescript","params":["00"],"id":1}`,
			unmarshalled: &pfcjson.DecodeScriptCmd{HexScript: "00"},
		},
		{
			name: "estimatesmartfee",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("estimatesmartfee", 6, pfcjson.EstimateSmartFeeConservative)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewEstimateSmartFeeCmd(6, pfcjson.EstimateSmartFeeConservative)
			},
			marshalled:   `{"jsonrpc":"1.0","method":"estimatesmartfee","params":[6,"conservative"],"id":1}`,
			unmarshalled: &pfcjson.EstimateSmartFeeCmd{Confirmations: 6, Mode: pfcjson.EstimateSmartFeeConservative},
		},
		{
			name: "generate",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("generate", 1)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGenerateCmd(1)
			},
			marshalled: `{"jsonrpc":"1.0","method":"generate","params":[1],"id":1}`,
			unmarshalled: &pfcjson.GenerateCmd{
				NumBlocks: 1,
			},
		},
		{
			name: "getaddednodeinfo",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getaddednodeinfo", true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetAddedNodeInfoCmd(true, nil)
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getaddednodeinfo","params":[true],"id":1}`,
			unmarshalled: &pfcjson.GetAddedNodeInfoCmd{DNS: true, Node: nil},
		},
		{
			name: "getaddednodeinfo optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getaddednodeinfo", true, "127.0.0.1")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetAddedNodeInfoCmd(true, pfcjson.String("127.0.0.1"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getaddednodeinfo","params":[true,"127.0.0.1"],"id":1}`,
			unmarshalled: &pfcjson.GetAddedNodeInfoCmd{
				DNS:  true,
				Node: pfcjson.String("127.0.0.1"),
			},
		},
		{
			name: "getbestblock",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getbestblock")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBestBlockCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getbestblock","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetBestBlockCmd{},
		},
		{
			name: "getbestblockhash",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getbestblockhash")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBestBlockHashCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getbestblockhash","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetBestBlockHashCmd{},
		},
		{
			name: "getblock",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblock", "123")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBlockCmd("123", nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getblock","params":["123"],"id":1}`,
			unmarshalled: &pfcjson.GetBlockCmd{
				Hash:      "123",
				Verbose:   pfcjson.Bool(true),
				VerboseTx: pfcjson.Bool(false),
			},
		},
		{
			name: "getblock required optional1",
			newCmd: func() (interface{}, error) {
				// Intentionally use a source param that is
				// more pointers than the destination to
				// exercise that path.
				verbosePtr := pfcjson.Bool(true)
				return pfcjson.NewCmd("getblock", "123", &verbosePtr)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBlockCmd("123", pfcjson.Bool(true), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getblock","params":["123",true],"id":1}`,
			unmarshalled: &pfcjson.GetBlockCmd{
				Hash:      "123",
				Verbose:   pfcjson.Bool(true),
				VerboseTx: pfcjson.Bool(false),
			},
		},
		{
			name: "getblock required optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblock", "123", true, true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBlockCmd("123", pfcjson.Bool(true), pfcjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getblock","params":["123",true,true],"id":1}`,
			unmarshalled: &pfcjson.GetBlockCmd{
				Hash:      "123",
				Verbose:   pfcjson.Bool(true),
				VerboseTx: pfcjson.Bool(true),
			},
		},
		{
			name: "getblockchaininfo",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblockchaininfo")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBlockChainInfoCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getblockchaininfo","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetBlockChainInfoCmd{},
		},
		{
			name: "getblockcount",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblockcount")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBlockCountCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getblockcount","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetBlockCountCmd{},
		},
		{
			name: "getblockhash",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblockhash", 123)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBlockHashCmd(123)
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getblockhash","params":[123],"id":1}`,
			unmarshalled: &pfcjson.GetBlockHashCmd{Index: 123},
		},
		{
			name: "getblockheader",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblockheader", "123")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBlockHeaderCmd("123", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getblockheader","params":["123"],"id":1}`,
			unmarshalled: &pfcjson.GetBlockHeaderCmd{
				Hash:    "123",
				Verbose: pfcjson.Bool(true),
			},
		},
		{
			name: "getblocksubsidy",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblocksubsidy", 123, 256)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBlockSubsidyCmd(123, 256)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getblocksubsidy","params":[123,256],"id":1}`,
			unmarshalled: &pfcjson.GetBlockSubsidyCmd{
				Height: 123,
				Voters: 256,
			},
		},
		{
			name: "getblocktemplate",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblocktemplate")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetBlockTemplateCmd(nil)
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getblocktemplate","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetBlockTemplateCmd{Request: nil},
		},
		{
			name: "getblocktemplate optional - template request",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblocktemplate", `{"mode":"template","capabilities":["longpoll","coinbasetxn"]}`)
			},
			staticCmd: func() interface{} {
				template := pfcjson.TemplateRequest{
					Mode:         "template",
					Capabilities: []string{"longpoll", "coinbasetxn"},
				}
				return pfcjson.NewGetBlockTemplateCmd(&template)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getblocktemplate","params":[{"mode":"template","capabilities":["longpoll","coinbasetxn"]}],"id":1}`,
			unmarshalled: &pfcjson.GetBlockTemplateCmd{
				Request: &pfcjson.TemplateRequest{
					Mode:         "template",
					Capabilities: []string{"longpoll", "coinbasetxn"},
				},
			},
		},
		{
			name: "getblocktemplate optional - template request with tweaks",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblocktemplate", `{"mode":"template","capabilities":["longpoll","coinbasetxn"],"sigoplimit":500,"sizelimit":100000000,"maxversion":2}`)
			},
			staticCmd: func() interface{} {
				template := pfcjson.TemplateRequest{
					Mode:         "template",
					Capabilities: []string{"longpoll", "coinbasetxn"},
					SigOpLimit:   500,
					SizeLimit:    100000000,
					MaxVersion:   2,
				}
				return pfcjson.NewGetBlockTemplateCmd(&template)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getblocktemplate","params":[{"mode":"template","capabilities":["longpoll","coinbasetxn"],"sigoplimit":500,"sizelimit":100000000,"maxversion":2}],"id":1}`,
			unmarshalled: &pfcjson.GetBlockTemplateCmd{
				Request: &pfcjson.TemplateRequest{
					Mode:         "template",
					Capabilities: []string{"longpoll", "coinbasetxn"},
					SigOpLimit:   int64(500),
					SizeLimit:    int64(100000000),
					MaxVersion:   2,
				},
			},
		},
		{
			name: "getblocktemplate optional - template request with tweaks 2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getblocktemplate", `{"mode":"template","capabilities":["longpoll","coinbasetxn"],"sigoplimit":true,"sizelimit":100000000,"maxversion":2}`)
			},
			staticCmd: func() interface{} {
				template := pfcjson.TemplateRequest{
					Mode:         "template",
					Capabilities: []string{"longpoll", "coinbasetxn"},
					SigOpLimit:   true,
					SizeLimit:    100000000,
					MaxVersion:   2,
				}
				return pfcjson.NewGetBlockTemplateCmd(&template)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getblocktemplate","params":[{"mode":"template","capabilities":["longpoll","coinbasetxn"],"sigoplimit":true,"sizelimit":100000000,"maxversion":2}],"id":1}`,
			unmarshalled: &pfcjson.GetBlockTemplateCmd{
				Request: &pfcjson.TemplateRequest{
					Mode:         "template",
					Capabilities: []string{"longpoll", "coinbasetxn"},
					SigOpLimit:   true,
					SizeLimit:    int64(100000000),
					MaxVersion:   2,
				},
			},
		},
		{
			name: "getcfilter",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getcfilter", "123", "extended")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetCFilterCmd("123", "extended")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getcfilter","params":["123","extended"],"id":1}`,
			unmarshalled: &pfcjson.GetCFilterCmd{
				Hash:       "123",
				FilterType: "extended",
			},
		},
		{
			name: "getcfilterheader",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getcfilterheader", "123", "extended")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetCFilterHeaderCmd("123", "extended")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getcfilterheader","params":["123","extended"],"id":1}`,
			unmarshalled: &pfcjson.GetCFilterHeaderCmd{
				Hash:       "123",
				FilterType: "extended",
			},
		},
		{
			name: "getchaintips",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getchaintips")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetChainTipsCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getchaintips","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetChainTipsCmd{},
		},
		{
			name: "getconnectioncount",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getconnectioncount")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetConnectionCountCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getconnectioncount","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetConnectionCountCmd{},
		},
		{
			name: "getcurrentnet",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getcurrentnet")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetCurrentNetCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getcurrentnet","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetCurrentNetCmd{},
		},
		{
			name: "getdifficulty",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getdifficulty")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetDifficultyCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getdifficulty","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetDifficultyCmd{},
		},
		{
			name: "getgenerate",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getgenerate")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetGenerateCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getgenerate","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetGenerateCmd{},
		},
		{
			name: "gethashespersec",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("gethashespersec")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetHashesPerSecCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"gethashespersec","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetHashesPerSecCmd{},
		},
		{
			name: "getinfo",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getinfo")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetInfoCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getinfo","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetInfoCmd{},
		},
		{
			name: "getmempoolinfo",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getmempoolinfo")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetMempoolInfoCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getmempoolinfo","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetMempoolInfoCmd{},
		},
		{
			name: "getmininginfo",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getmininginfo")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetMiningInfoCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getmininginfo","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetMiningInfoCmd{},
		},
		{
			name: "getnetworkinfo",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getnetworkinfo")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetNetworkInfoCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getnetworkinfo","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetNetworkInfoCmd{},
		},
		{
			name: "getnettotals",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getnettotals")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetNetTotalsCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getnettotals","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetNetTotalsCmd{},
		},
		{
			name: "getnetworkhashps",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getnetworkhashps")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetNetworkHashPSCmd(nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getnetworkhashps","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetNetworkHashPSCmd{
				Blocks: pfcjson.Int(120),
				Height: pfcjson.Int(-1),
			},
		},
		{
			name: "getnetworkhashps optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getnetworkhashps", 200)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetNetworkHashPSCmd(pfcjson.Int(200), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getnetworkhashps","params":[200],"id":1}`,
			unmarshalled: &pfcjson.GetNetworkHashPSCmd{
				Blocks: pfcjson.Int(200),
				Height: pfcjson.Int(-1),
			},
		},
		{
			name: "getnetworkhashps optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getnetworkhashps", 200, 123)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetNetworkHashPSCmd(pfcjson.Int(200), pfcjson.Int(123))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getnetworkhashps","params":[200,123],"id":1}`,
			unmarshalled: &pfcjson.GetNetworkHashPSCmd{
				Blocks: pfcjson.Int(200),
				Height: pfcjson.Int(123),
			},
		},
		{
			name: "getpeerinfo",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getpeerinfo")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetPeerInfoCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getpeerinfo","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetPeerInfoCmd{},
		},
		{
			name: "getrawmempool",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getrawmempool")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetRawMempoolCmd(nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawmempool","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetRawMempoolCmd{
				Verbose: pfcjson.Bool(false),
			},
		},
		{
			name: "getrawmempool optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getrawmempool", false)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetRawMempoolCmd(pfcjson.Bool(false), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawmempool","params":[false],"id":1}`,
			unmarshalled: &pfcjson.GetRawMempoolCmd{
				Verbose: pfcjson.Bool(false),
			},
		},
		{
			name: "getrawmempool optional 2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getrawmempool", false, "all")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetRawMempoolCmd(pfcjson.Bool(false), pfcjson.String("all"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawmempool","params":[false,"all"],"id":1}`,
			unmarshalled: &pfcjson.GetRawMempoolCmd{
				Verbose: pfcjson.Bool(false),
				TxType:  pfcjson.String("all"),
			},
		},
		{
			name: "getrawtransaction",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getrawtransaction", "123")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetRawTransactionCmd("123", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawtransaction","params":["123"],"id":1}`,
			unmarshalled: &pfcjson.GetRawTransactionCmd{
				Txid:    "123",
				Verbose: pfcjson.Int(0),
			},
		},
		{
			name: "getrawtransaction optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getrawtransaction", "123", 1)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetRawTransactionCmd("123", pfcjson.Int(1))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawtransaction","params":["123",1],"id":1}`,
			unmarshalled: &pfcjson.GetRawTransactionCmd{
				Txid:    "123",
				Verbose: pfcjson.Int(1),
			},
		},
		{
			name: "getstakeversions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getstakeversions", "deadbeef", 1)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetStakeVersionsCmd("deadbeef", 1)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getstakeversions","params":["deadbeef",1],"id":1}`,
			unmarshalled: &pfcjson.GetStakeVersionsCmd{
				Hash:  "deadbeef",
				Count: 1,
			},
		},
		{
			name: "gettxout",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("gettxout", "123", 1)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetTxOutCmd("123", 1, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"gettxout","params":["123",1],"id":1}`,
			unmarshalled: &pfcjson.GetTxOutCmd{
				Txid:           "123",
				Vout:           1,
				IncludeMempool: pfcjson.Bool(true),
			},
		},
		{
			name: "gettxout optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("gettxout", "123", 1, true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetTxOutCmd("123", 1, pfcjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"gettxout","params":["123",1,true],"id":1}`,
			unmarshalled: &pfcjson.GetTxOutCmd{
				Txid:           "123",
				Vout:           1,
				IncludeMempool: pfcjson.Bool(true),
			},
		},
		{
			name: "gettxoutsetinfo",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("gettxoutsetinfo")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetTxOutSetInfoCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"gettxoutsetinfo","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetTxOutSetInfoCmd{},
		},
		{
			name: "getvoteinfo",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getvoteinfo", 1)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetVoteInfoCmd(1)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getvoteinfo","params":[1],"id":1}`,
			unmarshalled: &pfcjson.GetVoteInfoCmd{
				Version: 1,
			},
		},
		{
			name: "getwork",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getwork")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetWorkCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getwork","params":[],"id":1}`,
			unmarshalled: &pfcjson.GetWorkCmd{
				Data: nil,
			},
		},
		{
			name: "getwork optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("getwork", "00112233")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewGetWorkCmd(pfcjson.String("00112233"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getwork","params":["00112233"],"id":1}`,
			unmarshalled: &pfcjson.GetWorkCmd{
				Data: pfcjson.String("00112233"),
			},
		},
		{
			name: "help",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("help")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewHelpCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"help","params":[],"id":1}`,
			unmarshalled: &pfcjson.HelpCmd{
				Command: nil,
			},
		},
		{
			name: "help optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("help", "getblock")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewHelpCmd(pfcjson.String("getblock"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"help","params":["getblock"],"id":1}`,
			unmarshalled: &pfcjson.HelpCmd{
				Command: pfcjson.String("getblock"),
			},
		},
		{
			name: "node option remove",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("node", pfcjson.NRemove, "1.1.1.1")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewNodeCmd("remove", "1.1.1.1", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"node","params":["remove","1.1.1.1"],"id":1}`,
			unmarshalled: &pfcjson.NodeCmd{
				SubCmd: pfcjson.NRemove,
				Target: "1.1.1.1",
			},
		},
		{
			name: "node option disconnect",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("node", pfcjson.NDisconnect, "1.1.1.1")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewNodeCmd("disconnect", "1.1.1.1", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"node","params":["disconnect","1.1.1.1"],"id":1}`,
			unmarshalled: &pfcjson.NodeCmd{
				SubCmd: pfcjson.NDisconnect,
				Target: "1.1.1.1",
			},
		},
		{
			name: "node option connect",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("node", pfcjson.NConnect, "1.1.1.1", "perm")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewNodeCmd("connect", "1.1.1.1", pfcjson.String("perm"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"node","params":["connect","1.1.1.1","perm"],"id":1}`,
			unmarshalled: &pfcjson.NodeCmd{
				SubCmd:        pfcjson.NConnect,
				Target:        "1.1.1.1",
				ConnectSubCmd: pfcjson.String("perm"),
			},
		},
		{
			name: "ping",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("ping")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewPingCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"ping","params":[],"id":1}`,
			unmarshalled: &pfcjson.PingCmd{},
		},
		{
			name: "searchrawtransactions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("searchrawtransactions", "1Address")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSearchRawTransactionsCmd("1Address", nil, nil, nil, nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"searchrawtransactions","params":["1Address"],"id":1}`,
			unmarshalled: &pfcjson.SearchRawTransactionsCmd{
				Address:     "1Address",
				Verbose:     pfcjson.Int(1),
				Skip:        pfcjson.Int(0),
				Count:       pfcjson.Int(100),
				VinExtra:    pfcjson.Int(0),
				Reverse:     pfcjson.Bool(false),
				FilterAddrs: nil,
			},
		},
		{
			name: "searchrawtransactions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("searchrawtransactions", "1Address", 0)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSearchRawTransactionsCmd("1Address",
					pfcjson.Int(0), nil, nil, nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"searchrawtransactions","params":["1Address",0],"id":1}`,
			unmarshalled: &pfcjson.SearchRawTransactionsCmd{
				Address:     "1Address",
				Verbose:     pfcjson.Int(0),
				Skip:        pfcjson.Int(0),
				Count:       pfcjson.Int(100),
				VinExtra:    pfcjson.Int(0),
				Reverse:     pfcjson.Bool(false),
				FilterAddrs: nil,
			},
		},
		{
			name: "searchrawtransactions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("searchrawtransactions", "1Address", 0, 5)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSearchRawTransactionsCmd("1Address",
					pfcjson.Int(0), pfcjson.Int(5), nil, nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"searchrawtransactions","params":["1Address",0,5],"id":1}`,
			unmarshalled: &pfcjson.SearchRawTransactionsCmd{
				Address:     "1Address",
				Verbose:     pfcjson.Int(0),
				Skip:        pfcjson.Int(5),
				Count:       pfcjson.Int(100),
				VinExtra:    pfcjson.Int(0),
				Reverse:     pfcjson.Bool(false),
				FilterAddrs: nil,
			},
		},
		{
			name: "searchrawtransactions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("searchrawtransactions", "1Address", 0, 5, 10)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSearchRawTransactionsCmd("1Address",
					pfcjson.Int(0), pfcjson.Int(5), pfcjson.Int(10), nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"searchrawtransactions","params":["1Address",0,5,10],"id":1}`,
			unmarshalled: &pfcjson.SearchRawTransactionsCmd{
				Address:     "1Address",
				Verbose:     pfcjson.Int(0),
				Skip:        pfcjson.Int(5),
				Count:       pfcjson.Int(10),
				VinExtra:    pfcjson.Int(0),
				Reverse:     pfcjson.Bool(false),
				FilterAddrs: nil,
			},
		},
		{
			name: "searchrawtransactions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("searchrawtransactions", "1Address", 0, 5, 10, 1)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSearchRawTransactionsCmd("1Address",
					pfcjson.Int(0), pfcjson.Int(5), pfcjson.Int(10), pfcjson.Int(1), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"searchrawtransactions","params":["1Address",0,5,10,1],"id":1}`,
			unmarshalled: &pfcjson.SearchRawTransactionsCmd{
				Address:     "1Address",
				Verbose:     pfcjson.Int(0),
				Skip:        pfcjson.Int(5),
				Count:       pfcjson.Int(10),
				VinExtra:    pfcjson.Int(1),
				Reverse:     pfcjson.Bool(false),
				FilterAddrs: nil,
			},
		},
		{
			name: "searchrawtransactions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("searchrawtransactions", "1Address", 0, 5, 10, 1, true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSearchRawTransactionsCmd("1Address",
					pfcjson.Int(0), pfcjson.Int(5), pfcjson.Int(10),
					pfcjson.Int(1), pfcjson.Bool(true), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"searchrawtransactions","params":["1Address",0,5,10,1,true],"id":1}`,
			unmarshalled: &pfcjson.SearchRawTransactionsCmd{
				Address:     "1Address",
				Verbose:     pfcjson.Int(0),
				Skip:        pfcjson.Int(5),
				Count:       pfcjson.Int(10),
				VinExtra:    pfcjson.Int(1),
				Reverse:     pfcjson.Bool(true),
				FilterAddrs: nil,
			},
		},
		{
			name: "searchrawtransactions",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("searchrawtransactions", "1Address", 0, 5, 10, 1, true, []string{"1Address"})
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSearchRawTransactionsCmd("1Address",
					pfcjson.Int(0), pfcjson.Int(5), pfcjson.Int(10),
					pfcjson.Int(1), pfcjson.Bool(true), &[]string{"1Address"})
			},
			marshalled: `{"jsonrpc":"1.0","method":"searchrawtransactions","params":["1Address",0,5,10,1,true,["1Address"]],"id":1}`,
			unmarshalled: &pfcjson.SearchRawTransactionsCmd{
				Address:     "1Address",
				Verbose:     pfcjson.Int(0),
				Skip:        pfcjson.Int(5),
				Count:       pfcjson.Int(10),
				VinExtra:    pfcjson.Int(1),
				Reverse:     pfcjson.Bool(true),
				FilterAddrs: &[]string{"1Address"},
			},
		},
		{
			name: "sendrawtransaction",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendrawtransaction", "1122")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSendRawTransactionCmd("1122", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendrawtransaction","params":["1122"],"id":1}`,
			unmarshalled: &pfcjson.SendRawTransactionCmd{
				HexTx:         "1122",
				AllowHighFees: pfcjson.Bool(false),
			},
		},
		{
			name: "sendrawtransaction optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("sendrawtransaction", "1122", false)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSendRawTransactionCmd("1122", pfcjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendrawtransaction","params":["1122",false],"id":1}`,
			unmarshalled: &pfcjson.SendRawTransactionCmd{
				HexTx:         "1122",
				AllowHighFees: pfcjson.Bool(false),
			},
		},
		{
			name: "setgenerate",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("setgenerate", true)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSetGenerateCmd(true, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"setgenerate","params":[true],"id":1}`,
			unmarshalled: &pfcjson.SetGenerateCmd{
				Generate:     true,
				GenProcLimit: pfcjson.Int(-1),
			},
		},
		{
			name: "setgenerate optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("setgenerate", true, 6)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSetGenerateCmd(true, pfcjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"setgenerate","params":[true,6],"id":1}`,
			unmarshalled: &pfcjson.SetGenerateCmd{
				Generate:     true,
				GenProcLimit: pfcjson.Int(6),
			},
		},
		{
			name: "stop",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("stop")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewStopCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"stop","params":[],"id":1}`,
			unmarshalled: &pfcjson.StopCmd{},
		},
		{
			name: "submitblock",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("submitblock", "112233")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewSubmitBlockCmd("112233", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"submitblock","params":["112233"],"id":1}`,
			unmarshalled: &pfcjson.SubmitBlockCmd{
				HexBlock: "112233",
				Options:  nil,
			},
		},
		{
			name: "submitblock optional",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("submitblock", "112233", `{"workid":"12345"}`)
			},
			staticCmd: func() interface{} {
				options := pfcjson.SubmitBlockOptions{
					WorkID: "12345",
				}
				return pfcjson.NewSubmitBlockCmd("112233", &options)
			},
			marshalled: `{"jsonrpc":"1.0","method":"submitblock","params":["112233",{"workid":"12345"}],"id":1}`,
			unmarshalled: &pfcjson.SubmitBlockCmd{
				HexBlock: "112233",
				Options: &pfcjson.SubmitBlockOptions{
					WorkID: "12345",
				},
			},
		},
		{
			name: "validateaddress",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("validateaddress", "1Address")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewValidateAddressCmd("1Address")
			},
			marshalled: `{"jsonrpc":"1.0","method":"validateaddress","params":["1Address"],"id":1}`,
			unmarshalled: &pfcjson.ValidateAddressCmd{
				Address: "1Address",
			},
		},
		{
			name: "verifychain",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("verifychain")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewVerifyChainCmd(nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"verifychain","params":[],"id":1}`,
			unmarshalled: &pfcjson.VerifyChainCmd{
				CheckLevel: pfcjson.Int64(3),
				CheckDepth: pfcjson.Int64(288),
			},
		},
		{
			name: "verifychain optional1",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("verifychain", 2)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewVerifyChainCmd(pfcjson.Int64(2), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"verifychain","params":[2],"id":1}`,
			unmarshalled: &pfcjson.VerifyChainCmd{
				CheckLevel: pfcjson.Int64(2),
				CheckDepth: pfcjson.Int64(288),
			},
		},
		{
			name: "verifychain optional2",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("verifychain", 2, 500)
			},
			staticCmd: func() interface{} {
				return pfcjson.NewVerifyChainCmd(pfcjson.Int64(2), pfcjson.Int64(500))
			},
			marshalled: `{"jsonrpc":"1.0","method":"verifychain","params":[2,500],"id":1}`,
			unmarshalled: &pfcjson.VerifyChainCmd{
				CheckLevel: pfcjson.Int64(2),
				CheckDepth: pfcjson.Int64(500),
			},
		},
		{
			name: "verifymessage",
			newCmd: func() (interface{}, error) {
				return pfcjson.NewCmd("verifymessage", "1Address", "301234", "test")
			},
			staticCmd: func() interface{} {
				return pfcjson.NewVerifyMessageCmd("1Address", "301234", "test")
			},
			marshalled: `{"jsonrpc":"1.0","method":"verifymessage","params":["1Address","301234","test"],"id":1}`,
			unmarshalled: &pfcjson.VerifyMessageCmd{
				Address:   "1Address",
				Signature: "301234",
				Message:   "test",
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
			t.Errorf("\n%s\n%s", marshalled, test.marshalled)
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

// TestChainSvrCmdErrors ensures any errors that occur in the command during
// custom mashal and unmarshal are as expected.
func TestChainSvrCmdErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		result     interface{}
		marshalled string
		err        error
	}{
		{
			name:       "template request with invalid type",
			result:     &pfcjson.TemplateRequest{},
			marshalled: `{"mode":1}`,
			err:        &json.UnmarshalTypeError{},
		},
		{
			name:       "invalid template request sigoplimit field",
			result:     &pfcjson.TemplateRequest{},
			marshalled: `{"sigoplimit":"invalid"}`,
			err:        pfcjson.Error{Code: pfcjson.ErrInvalidType},
		},
		{
			name:       "invalid template request sizelimit field",
			result:     &pfcjson.TemplateRequest{},
			marshalled: `{"sizelimit":"invalid"}`,
			err:        pfcjson.Error{Code: pfcjson.ErrInvalidType},
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		err := json.Unmarshal([]byte(test.marshalled), &test.result)
		if reflect.TypeOf(err) != reflect.TypeOf(test.err) {
			t.Errorf("Test #%d (%s) wrong error type - got `%T` (%v), got `%T`",
				i, test.name, err, err, test.err)
			continue
		}

		if terr, ok := test.err.(pfcjson.Error); ok {
			gotErrorCode := err.(pfcjson.Error).Code
			if gotErrorCode != terr.Code {
				t.Errorf("Test #%d (%s) mismatched error code "+
					"- got %v (%v), want %v", i, test.name,
					gotErrorCode, terr, terr.Code)
				continue
			}
		}
	}
}
