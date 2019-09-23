// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/decred/dcrd/wire"
)

// PicFightCoinParams defines the network parameters for the PicFight coin network.
var PicFightCoinParams = Params{
	Name:        "picfight",
	Net:         wire.MainNet,
	DefaultPort: "9108",
	DNSSeeds: []DNSSeed{
		{"eu-01.picfight.org", true},
		{"eu-02.picfight.org", true},
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
	Checkpoints: []Checkpoint{
	},

	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationQuorum:     4032, // 10 % of RuleChangeActivationInterval * TicketsPerBlock
	RuleChangeActivationMultiplier: 3,    // 75%
	RuleChangeActivationDivisor:    4,
	RuleChangeActivationInterval:   2016 * 4, // 4 weeks

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
	NetworkAddressPrefix: "D",
	PubKeyAddrID:         [2]byte{0x13, 0x86}, // starts with Dk
	PubKeyHashAddrID:     [2]byte{0x07, 0x3f}, // starts with Ds
	PKHEdwardsAddrID:     [2]byte{0x07, 0x1f}, // starts with De
	PKHSchnorrAddrID:     [2]byte{0x07, 0x01}, // starts with DS
	ScriptHashAddrID:     [2]byte{0x07, 0x1a}, // starts with Dc
	PrivateKeyID:         [2]byte{0x22, 0xde}, // starts with Pm

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
	BlockOneLedger:              BlockOneLedgerPicFight,
}
