package chaincfg

import (
	"github.com/picfight/pfcd/wire"
	"time"
)

// MainNetParams defines the network parameters for the main Picfightcoin network.
var MainNetParams = Params{
	Name:             "mainnet",
	NodeBuildVersion: build_id + ".mainnet",
	Net:              wire.MainNet,
	DefaultPort:      "8333",
	DNSSeeds: []DNSSeed{
		{"eu-01.mainnet.picfight.org", true},
	},

	// Blockchain parameters
	MaxTimeOffsetSeconds:    2 * 60 * 60,
	MinCoinbaseScriptLen:    2,
	MaxCoinbaseScriptLen:    100,
	MedianTimeBlocks:        11,
	TargetTotalSubsidy:      7777777,
	SubsidyProductionPeriod: time.Hour * time.Duration(24*365*SubsidyProductionYears),

	// Chain parameters
	GenesisBlock:     &genesisBlock,
	GenesisHash:      &genesisHash,
	PowLimit:         mainPowLimit,
	PowLimitBits:     0x1e00ffff,
	CoinbaseMaturity: 100,
	//SubsidyReductionInterval: 210000,
	TargetTimespan:           time.Hour * 24 * 14, // 14 days
	TargetTimePerBlock:       time.Minute * 1,     // 1 minute
	RetargetAdjustmentFactor: 4,                   // 25% less, 400% more
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0,
	GenerateSupported:        false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: nil,

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 1916, // 95% of MinerConfirmationWindow
	MinerConfirmationWindow:       2016, //
	Deployments: [DefinedDeployments]ConsensusDeployment{
		DeploymentTestDummy: {
			BitNumber:  28,
			StartTime:  1546300801, // January 1, 2019 UTC
			ExpireTime: 1577836799, // December 31, 2019 UTC
		},
	},

	// Mempool parameters
	RelayNonStdTxs: false,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "bc", // always bc for main net

	// Address encoding magics
	PubKeyHashAddrID:        0x00, // starts with 1
	ScriptHashAddrID:        0x05, // starts with 3
	PrivateKeyID:            0x80, // starts with 5 (uncompressed) or K (compressed)
	WitnessPubKeyHashAddrID: 0x06, // starts with p2
	WitnessScriptHashAddrID: 0x0A, // starts with 7Xh

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 0,
}
