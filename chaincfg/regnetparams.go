package chaincfg

import (
	"github.com/picfight/pfcd/wire"
	"math"
	"time"
)

// RegressionNetParams defines the network parameters for the regression test
// Picfightcoin network.  Not to be confused with the test Picfightcoin network (version
// 3), this network is sometimes simply called "testnet".
var RegressionNetParams = Params{
	Name:             "regtest",
	NodeBuildVersion: build_id + ".regtest",
	Net:              wire.TestNet,
	DefaultPort:      "18444",
	DNSSeeds:         []DNSSeed{},

	// Blockchain parameters
	MaxTimeOffsetSeconds:    2 * 60 * 60,
	MinCoinbaseScriptLen:    2,
	MaxCoinbaseScriptLen:    100,
	MedianTimeBlocks:        11,
	TargetTotalSubsidy:      7777777,
	SubsidyProductionPeriod: time.Hour * time.Duration(24*365*SubsidyProductionYears),

	// Chain parameters
	GenesisBlock:     &regTestGenesisBlock,
	GenesisHash:      &regTestGenesisHash,
	PowLimit:         regressionPowLimit.ToBigInt(),
	PowLimitBits:     regressionPowLimit.ToCompact(),
	CoinbaseMaturity: 100,
	//SubsidyReductionInterval: 150,
	TargetTimespan:           time.Hour * 24 * 14, // 14 days
	TargetTimePerBlock:       time.Minute * 1,     // 1 minute
	RetargetAdjustmentFactor: 4,                   // 25% less, 400% more
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Minute * 20, // TargetTimePerBlock * 2
	GenerateSupported:        true,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: nil,

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 108, // 75%  of MinerConfirmationWindow
	MinerConfirmationWindow:       144,
	Deployments: [DefinedDeployments]ConsensusDeployment{
		DeploymentTestDummy: {
			BitNumber:  28,
			StartTime:  0,             // Always available for vote
			ExpireTime: math.MaxInt64, // Never expires
		},
	},

	// Mempool parameters
	RelayNonStdTxs: true,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "bcrt", // always bcrt for reg test net

	// Address encoding magics
	PubKeyHashAddrID: 0x6f, // starts with m or n
	ScriptHashAddrID: 0xc4, // starts with 2
	PrivateKeyID:     0xef, // starts with 9 (uncompressed) or c (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}
