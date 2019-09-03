package chaincfg

import (
	"github.com/picfight/pfcd/wire"
	"time"
)

// TestNet3Params defines the network parameters for the test Picfightcoin network
// (version 3).  Not to be confused with the regression test network, this
// network is sometimes simply called "testnet".
var TestNet3Params = Params{
	Name:             "testnet3",
	NodeBuildVersion: build_id + ".testnet3",
	Net:              wire.TestNet3,
	DefaultPort:      "18333",
	DNSSeeds: []DNSSeed{
		{"eu-01.testnet.picfight.org", true},
		{"eu-02.testnet.picfight.org", true},
	},

	// Blockchain parameters
	MaxTimeOffsetSeconds:    2 * 60 * 60,
	MinCoinbaseScriptLen:    2,
	MaxCoinbaseScriptLen:    100,
	MedianTimeBlocks:        11,
	TargetTotalSubsidy:      7777777,
	SubsidyProductionPeriod: time.Hour * time.Duration(24*365*SubsidyProductionYears),

	// Chain parameters
	GenesisBlock:     &testNet3GenesisBlock,
	GenesisHash:      &testNet3GenesisHash,
	PowLimit:         testNet3PowLimit.ToBigInt(),
	PowLimitBits:     testNet3PowLimit.ToCompact(),
	CoinbaseMaturity: 100,
	//SubsidyReductionInterval: 210000,
	TargetTimespan:           time.Hour * 24 * 14, // 14 days
	TargetTimePerBlock:       time.Minute * 1,     // 1 minute
	RetargetAdjustmentFactor: 4,                   // 25% less, 400% more
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Minute * 20, // TargetTimePerBlock * 2
	GenerateSupported:        false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: nil,

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 1512, // 75% of MinerConfirmationWindow
	MinerConfirmationWindow:       2016,
	Deployments: [DefinedDeployments]ConsensusDeployment{
		DeploymentTestDummy: {
			BitNumber:  28,
			StartTime:  1546300801, // January 1, 2019 UTC
			ExpireTime: 1577836799, // December 31, 2019 UTC
		},
	},

	// Mempool parameters
	RelayNonStdTxs: true,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "tb", // always tb for test net

	// Address encoding magics
	PubKeyHashAddrID:        0x6f, // starts with m or n
	ScriptHashAddrID:        0xc4, // starts with 2
	WitnessPubKeyHashAddrID: 0x03, // starts with QW
	WitnessScriptHashAddrID: 0x28, // starts with T7n
	PrivateKeyID:            0xef, // starts with 9 (uncompressed) or c (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}
