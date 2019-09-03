package chaincfg

import (
	"github.com/picfight/pfcd/wire"
	"math"
	"time"
)

// SimNetParams defines the network parameters for the simulation test Picfightcoin
// network.  This network is similar to the normal test network except it is
// intended for private use within a group of individuals doing simulation
// testing.  The functionality is intended to differ in that the only nodes
// which are specifically specified are used to create the network rather than
// following normal discovery rules.  This is important as otherwise it would
// just turn into another public testnet.
var SimNetParams = Params{
	Name:             "simnet",
	NodeBuildVersion: build_id + ".simnet",
	Net:              wire.SimNet,
	DefaultPort:      "18555",
	DNSSeeds:         []DNSSeed{}, // NOTE: There must NOT be any seeds.

	// Blockchain parameters
	MaxTimeOffsetSeconds:    2 * 60 * 60,
	MinCoinbaseScriptLen:    2,
	MaxCoinbaseScriptLen:    100,
	MedianTimeBlocks:        11,
	TargetTotalSubsidy:      7777777,
	SubsidyProductionPeriod: time.Hour * time.Duration(24*365*SubsidyProductionYears),

	// Chain parameters
	GenesisBlock:     &simNetGenesisBlock,
	GenesisHash:      &simNetGenesisHash,
	PowLimit:         simNetPowLimit.ToBigInt(),
	PowLimitBits:     simNetPowLimit.ToCompact(),
	CoinbaseMaturity: 100,
	//SubsidyReductionInterval: 210000,
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
	RuleChangeActivationThreshold: 75, // 75% of MinerConfirmationWindow
	MinerConfirmationWindow:       100,
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
	Bech32HRPSegwit: "sb", // always sb for sim net

	// Address encoding magics
	PubKeyHashAddrID:        0x3f, // starts with S
	ScriptHashAddrID:        0x7b, // starts with s
	PrivateKeyID:            0x64, // starts with 4 (uncompressed) or F (compressed)
	WitnessPubKeyHashAddrID: 0x19, // starts with Gg
	WitnessScriptHashAddrID: 0x28, // starts with ?

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x20, 0xb9, 0x00}, // starts with sprv
	HDPublicKeyID:  [4]byte{0x04, 0x20, 0xbd, 0x3a}, // starts with spub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 115, // ASCII for s
}
