// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/picfight/pfcd/wire"
)

// MainNetParams defines the network parameters for the main PicFight network.
var MainNetParams = Params{
	Name:        "mainnet",
	Net:         wire.MainNet,
	DefaultPort: "9708",
	DNSSeeds: []DNSSeed{
		{"mainnet-seed.picfight.org", true},
		{"mainnet-seed.eu-1.picfight.org", true},
		{"mainnet-seed.us-1.picfight.org", true},
	},

	// Chain parameters
	GenesisBlock:             &genesisBlock,
	GenesisHash:              &genesisHash,
	PowLimit:                 mainPowLimit,
	PowLimitBits:             0x1e00ffff,
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0, // Does not apply since ReduceMinDifficulty false
	GenerateSupported:        false,
	MaximumBlockSizes:        []int{393216},
	MaxTxSize:                393216,
	TargetTimePerBlock:       time.Second * 60, //
	WorkDiffAlpha:            1,
	WorkDiffWindowSize:       144,
	WorkDiffWindows:          20,
	TargetTimespan:           time.Second * 60 * 144, // BlockTime * WindowSize
	RetargetAdjustmentFactor: 4,

	// Subsidy parameters.
	BaseSubsidy:           int64(1 * 1e8), // 1 coin
	WorkRewardProportion:  30,             // 30%
	StakeRewardProportion: 30,             // 30%
	BlockArtTaxProportion: 30,             // 30%
	BlockDevTaxProportion: 10,             // 10%

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{},

	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationQuorum:     4032, // 10 % of RuleChangeActivationInterval * TicketsPerBlock
	RuleChangeActivationMultiplier: 3,    // 75%
	RuleChangeActivationDivisor:    4,
	RuleChangeActivationInterval:   2016 * 4, // 4 weeks
	Deployments: map[uint32][]ConsensusDeployment{
		4: {{
			Vote: Vote{
				Id:          VoteIDSDiffAlgorithm,
				Description: "Change stake difficulty algorithm as defined in DCP0001",
				Mask:        0x0006, // Bits 1 and 2
				Choices: []Choice{{
					Id:          "abstain",
					Description: "abstain voting for change",
					Bits:        0x0000,
					IsAbstain:   true,
					IsNo:        false,
				}, {
					Id:          "no",
					Description: "keep the existing algorithm",
					Bits:        0x0002, // Bit 1
					IsAbstain:   false,
					IsNo:        true,
				}, {
					Id:          "yes",
					Description: "change to the new algorithm",
					Bits:        0x0004, // Bit 2
					IsAbstain:   false,
					IsNo:        false,
				}},
			},
			StartTime:  1493164800, // Apr 26th, 2017
			ExpireTime: 1524700800, // Apr 26th, 2018
		}, {
			Vote: Vote{
				Id:          VoteIDLNSupport,
				Description: "Request developers begin work on Lightning Network (LN) integration",
				Mask:        0x0018, // Bits 3 and 4
				Choices: []Choice{{
					Id:          "abstain",
					Description: "abstain from voting",
					Bits:        0x0000,
					IsAbstain:   true,
					IsNo:        false,
				}, {
					Id:          "no",
					Description: "no, do not work on integrating LN support",
					Bits:        0x0008, // Bit 3
					IsAbstain:   false,
					IsNo:        true,
				}, {
					Id:          "yes",
					Description: "yes, begin work on integrating LN support",
					Bits:        0x0010, // Bit 4
					IsAbstain:   false,
					IsNo:        false,
				}},
			},
			StartTime:  1493164800, // Apr 26th, 2017
			ExpireTime: 1508976000, // Oct 26th, 2017
		}},
		5: {{
			Vote: Vote{
				Id:          VoteIDLNFeatures,
				Description: "Enable features defined in DCP0002 and DCP0003 necessary to support Lightning Network (LN)",
				Mask:        0x0006, // Bits 1 and 2
				Choices: []Choice{{
					Id:          "abstain",
					Description: "abstain voting for change",
					Bits:        0x0000,
					IsAbstain:   true,
					IsNo:        false,
				}, {
					Id:          "no",
					Description: "keep the existing consensus rules",
					Bits:        0x0002, // Bit 1
					IsAbstain:   false,
					IsNo:        true,
				}, {
					Id:          "yes",
					Description: "change to the new consensus rules",
					Bits:        0x0004, // Bit 2
					IsAbstain:   false,
					IsNo:        false,
				}},
			},
			StartTime:  1505260800, // Sep 13th, 2017
			ExpireTime: 1536796800, // Sep 13th, 2018
		}},
	},

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
	SLIP0044CoinType: 42, // SLIP0044, PicFight
	LegacyCoinType:   20, // for backwards compatibility

	// PicFight PoS parameters
	MinimumStakeDiff:        1 * 1e8, // 2 Coin
	TicketPoolSize:          8192,
	TicketsPerBlock:         5,
	TicketMaturity:          256 * extentionParameter,
	TicketExpiry:            40960, // 5*TicketPoolSize
	CoinbaseMaturity:        256 * extentionParameter,
	SStxChangeMaturity:      1,
	TicketPoolSizeWeight:    4,
	StakeDiffAlpha:          1, // Minimal
	StakeDiffWindowSize:     144,
	StakeDiffWindows:        20,
	StakeVersionInterval:    144 * 2 * 7 * extentionParameter, // ~1 week
	MaxFreshStakePerBlock:   20,                               // 4*TicketsPerBlock
	StakeEnabledHeight:      (256 + 256) * extentionParameter, // CoinbaseMaturity + TicketMaturity
	StakeValidationHeight:   4096 * extentionParameter,        // ~14 days
	StakeBaseSigScript:      []byte{0x00, 0x00},
	StakeMajorityMultiplier: 3,
	StakeMajorityDivisor:    4,

	// PicFight organization related parameters
	// Organization address is ?
	OrganizationDevelopersPkScript:        hexDecode(""),
	OrganizationDevelopersPkScriptVersion: 0,

	OrganizationArtistsPkScript:        hexDecode(""),
	OrganizationArtistsPkScriptVersion: 0,

	BlockOneLedger: BlockOneLedgerMainNet,
}
