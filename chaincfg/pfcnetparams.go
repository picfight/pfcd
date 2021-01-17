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
		{"eu-01.seed.picfight.org", true},
		{"eu-02.seed.picfight.org", true},
		//
		{"us-01.seed.picfight.org", true},
		{"us-02.seed.picfight.org", true},
		//
		{"ap-01.seed.picfight.org", true},
		{"ap-02.seed.picfight.org", true},
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

	// Subsidy parameters.DecodeAddress
	SubsidyCalculator:     picfightcoin.PicFightCoinSubsidy,
	WorkRewardProportion:  picfightcoin.PicFightCoinSubsidy().WorkRewardProportion(),  //
	StakeRewardProportion: picfightcoin.PicFightCoinSubsidy().StakeRewardProportion(), //
	BlockTaxProportion:    picfightcoin.PicFightCoinSubsidy().BlockTaxProportion(),    //

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
		{CheckpointFlag: CHECKPOINT_FLAG_HASH_MUST_NOT_BE_EQUAL, Height: 98710, Hash: newHashFromStr("0000000eabc543ce2cc81e22d621e8820d3a3a1b6b71202c8d720fea7bba0566")},
	},
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

	// PoS parameters
	MinimumStakeDiff:        picfightcoin.MinStakeDifficulty(),
	TicketPoolSize:          8192,
	TicketsPerBlock:         picfightcoin.PicFightCoinSubsidy().TicketsPerBlock(),
	TicketMaturity:          256,
	TicketExpiry:            40960, // 5*TicketPoolSize
	CoinbaseMaturity:        256,
	SStxChangeMaturity:      1,
	TicketPoolSizeWeight:    4,
	StakeDiffAlpha:          1, // Minimal
	StakeDiffWindowSize:     144,
	StakeDiffWindows:        20,
	StakeVersionInterval:    144 * 2 * 7,                                                // ~1 week
	MaxFreshStakePerBlock:   20,                                                         // 4*TicketsPerBlock
	StakeEnabledHeight:      256 + 256,                                                  // CoinbaseMaturity + TicketMaturity
	StakeValidationHeight:   picfightcoin.PicFightCoinSubsidy().StakeValidationHeight(), // 4096,        // ~14 days
	StakeBaseSigScript:      []byte{0x00, 0x00},
	StakeMajorityMultiplier: 3,
	StakeMajorityDivisor:    4,

	OrganizationPkScript:        picfightcoin.OrganizationPkScript(),
	OrganizationPkScriptVersion: 0,
	BlockOneLedger:              BlockOneLedgerPicfightCoin(),
}
