package picfightcoin

import (
	"time"

	"github.com/decred/dcrd/chaincfg"
	"github.com/jfixby/difficulty"
)

// These variables are the chain proof-of-work limit parameters for each default
// network.
var (
	//  picfightPowLimit value for the PicFight coin network.
	picfightPowLimit = difficulty.NewDifficultyFromHashString( //
		"00 00 ff ff ffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
)

// PicFightCoinNetParams defines the network parameters for the main Decred network.
var PicFightCoinNetParams = chaincfg.Params{
	Name:        "picfightcoin",
	Net:         PicfightCoinWire,
	DefaultPort: "9108",
	DNSSeeds: []chaincfg.DNSSeed{
		{"eu-01.seed.picfight.org", true},
		{"eu-02.seed.picfight.org", true},
	},

	// Chain parameters
	GenesisBlock:             &genesisBlock,
	GenesisHash:              &genesisHash,
	PowLimit:                 picfightPowLimit.ToBigInt(),
	PowLimitBits:             picfightPowLimit.ToCompact(),
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
	BaseSubsidy:              3119582664, // 21m
	MulSubsidy:               100,
	DivSubsidy:               101,
	SubsidyReductionInterval: 6144,
	WorkRewardProportion:     6,
	StakeRewardProportion:    3,
	BlockTaxProportion:       1,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []chaincfg.Checkpoint{},

	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationQuorum:     4032, // 10 % of RuleChangeActivationInterval * TicketsPerBlock
	RuleChangeActivationMultiplier: 3,    // 75%
	RuleChangeActivationDivisor:    4,
	RuleChangeActivationInterval:   2016 * 4, // 4 weeks
	Deployments:                    map[uint32][]chaincfg.ConsensusDeployment{},

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
	NetworkAddressPrefix: "J",
	PubKeyAddrID:         [2]byte{0x1b, 0x2d}, // starts with Jk
	PubKeyHashAddrID:     [2]byte{0x0a, 0x0f}, // starts with Js
	PKHEdwardsAddrID:     [2]byte{0x09, 0xef}, // starts with Je
	PKHSchnorrAddrID:     [2]byte{0x09, 0xd1}, // starts with JS
	ScriptHashAddrID:     [2]byte{0x09, 0xea}, // starts with Jc
	PrivateKeyID:         [2]byte{0x22, 0xce}, // starts with Pj

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

	// Decred organization related parameters
	// Organization address is Dcur2mcGjmENx4DhNqDctW5wJCVyT3Qeqkx
	OrganizationPkScript:        hexDecode("a914f5916158e3e2c4551c1796708db8367207ed13bb87"),
	OrganizationPkScriptVersion: 0,
	BlockOneLedger:              BlockOneLedgerPicfightCoin,
}
