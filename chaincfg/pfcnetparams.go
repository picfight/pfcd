package chaincfg

import (
	"github.com/picfight/pfcd/wire"
	"github.com/picfight/picfightcoin"
	"time"
)

// PicFightCoinNetParams defines the network parameters for the main Decred network.
var PicFightCoinNetParams = Params{
	Name:        "picfightcoin",
	Net:         wire.PicfightCoinWire,
	DefaultPort: "9108",
	DNSSeeds: []DNSSeed{
		{picfightcoin.DNSSeeds()[0], true},
		{picfightcoin.DNSSeeds()[1], true},
	},

	// Chain parameters
	GenesisBlock:             &picfightGenesisBlock,
	GenesisHash:              &picfightGenesisHash,
	PowLimit:                 picfightcoin.NetworkPoWLimit().ToBigInt(),
	PowLimitBits:             picfightcoin.NetworkPoWLimit().ToCompact(),
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0, // Does not apply since ReduceMinDifficulty false
	GenerateSupported:        false,
	MaximumBlockSizes:        []int{393216},
	MaxTxSize:                393216,
	TargetTimePerBlock:       time.Minute * 5,
	WorkDiffAlpha:            1,
	WorkDiffWindowSize:       144,
	WorkDiffWindows:          20,
	TargetTimespan:           time.Minute * 5 * 144, // TimePerBlock * WindowSize
	RetargetAdjustmentFactor: 4,

	// Subsidy parameters.
	BaseSubsidy:              0,    // not used
	MulSubsidy:               100,  // not used
	DivSubsidy:               101,  // not used
	SubsidyReductionInterval: 6144, // not used
	WorkRewardProportion:     6,
	StakeRewardProportion:    3,
	BlockTaxProportion:       1,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{},

	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationQuorum:     4032, // 10 % of RuleChangeActivationInterval * TicketsPerBlock
	RuleChangeActivationMultiplier: 3,    // 75%
	RuleChangeActivationDivisor:    4,
	RuleChangeActivationInterval:   2016 * 4, // 4 weeks
	Deployments:                    map[uint32][]ConsensusDeployment{},

	// Enforce current block version once majority of the network has
	// upgraded.
	// 75% (750 / 1000)
	// Reject previous block versions once a majority of the network has
	// upgraded.
	// 95% (950 / 1000)
	BlockEnforceNumRequired: 750,
	BlockRejectNumRequired:  950,
	BlockUpgradeNumToCheck:  1000,

	// AcceptNonStdTxs is a mempool param to either accept and relay
	// non standard txs to the network or reject them
	AcceptNonStdTxs: false,

	// Address encoding magics
	NetworkAddressPrefix: picfightcoin.NetworkAddressPrefix,
	PubKeyAddrID:         picfightcoin.PubKeyAddrID,
	PubKeyHashAddrID:     picfightcoin.PubKeyHashAddrID,
	PKHEdwardsAddrID:     picfightcoin.PKHEdwardsAddrID,
	PKHSchnorrAddrID:     picfightcoin.PKHSchnorrAddrID,
	ScriptHashAddrID:     picfightcoin.ScriptHashAddrID,
	PrivateKeyID:         picfightcoin.PrivateKeyID,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x02, 0xfd, 0xa4, 0xe8}, // starts with dprv
	HDPublicKeyID:  [4]byte{0x02, 0xfd, 0xa9, 0x26}, // starts with dpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	SLIP0044CoinType: 42, // SLIP0044, Decred
	LegacyCoinType:   20, // for backwards compatibility

	// Decred PoS parameters
	MinimumStakeDiff:        2 * 1e8, // 2 Coin
	TicketPoolSize:          8192,
	TicketsPerBlock:         5,
	TicketMaturity:          256,
	TicketExpiry:            40960, // 5*TicketPoolSize
	CoinbaseMaturity:        256,
	SStxChangeMaturity:      1,
	TicketPoolSizeWeight:    4,
	StakeDiffAlpha:          1, // Minimal
	StakeDiffWindowSize:     144,
	StakeDiffWindows:        20,
	StakeVersionInterval:    144 * 2 * 7, // ~1 week
	MaxFreshStakePerBlock:   20,          // 4*TicketsPerBlock
	StakeEnabledHeight:      256 + 256,   // CoinbaseMaturity + TicketMaturity
	StakeValidationHeight:   4096,        // ~14 days
	StakeBaseSigScript:      []byte{0x00, 0x00},
	StakeMajorityMultiplier: 3,
	StakeMajorityDivisor:    4,

	OrganizationPkScript:        picfightcoin.OrganizationPkScript(),
	OrganizationPkScriptVersion: 0,
	BlockOneLedger:              BlockOneLedgerPicfightCoin(),
}
