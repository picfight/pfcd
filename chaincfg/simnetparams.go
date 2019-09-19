// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/picfight/pfcd/wire"
)

// SimNetParams defines the network parameters for the simulation test network.
// This network is similar to the normal test network except it is intended for
// private use within a group of individuals doing simulation testing and full
// integration tests between different applications such as wallets, voting
// service providers, mining pools, block explorers, and other services that
// build on PicFight.
//
// The functionality is intended to differ in that the only nodes which are
// specifically specified are used to create the network rather than following
// normal discovery rules.  This is important as otherwise it would just turn
// into another public testnet.
var SimNetParams = Params{
	Name:        "simnet",
	Net:         wire.SimNet,
	DefaultPort: "18555",
	DNSSeeds:    nil, // NOTE: There must NOT be any seeds.

	// Chain parameters
	GenesisBlock:             &simNetGenesisBlock,
	GenesisHash:              &simNetGenesisHash,
	PowLimit:                 simNetPowLimit.ToBigInt(),
	PowLimitBits:             simNetPowLimit.ToCompact(),
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0, // Does not apply since ReduceMinDifficulty false
	GenerateSupported:        true,
	MaximumBlockSizes:        []int{1310720},
	MaxTxSize:                1000000,
	TargetTimePerBlock:       time.Second,
	WorkDiffAlpha:            1,
	WorkDiffWindowSize:       8,
	WorkDiffWindows:          4,
	TargetTimespan:           time.Second * 8, // TimePerBlock * WindowSize
	RetargetAdjustmentFactor: 4,

	// Subsidy parameters.
	BaseSubsidy:              50000000000,
	MulSubsidy:               100,
	DivSubsidy:               101,
	SubsidyReductionInterval: 128,
	WorkRewardProportion:     6,
	StakeRewardProportion:    3,
	BlockTaxProportion:       1,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: nil,

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationQuorum:     160, // 10 % of RuleChangeActivationInterval * TicketsPerBlock
	RuleChangeActivationMultiplier: 3,   // 75%
	RuleChangeActivationDivisor:    4,
	RuleChangeActivationInterval:   320, // 320 seconds

	// Enforce current block version once majority of the network has
	// upgraded.
	// 51% (51 / 100)
	// Reject previous block versions once a majority of the network has
	// upgraded.
	// 75% (75 / 100)
	BlockEnforceNumRequired: 51,
	BlockRejectNumRequired:  75,
	BlockUpgradeNumToCheck:  100,

	// AcceptNonStdTxs is a mempool param to either accept and relay
	// non standard txs to the network or reject them
	AcceptNonStdTxs: true,

	// Address encoding magics
	NetworkAddressPrefix: "S",
	PubKeyAddrID:         [2]byte{0x27, 0x6f}, // starts with Sk
	PubKeyHashAddrID:     [2]byte{0x0e, 0x91}, // starts with Ss
	PKHEdwardsAddrID:     [2]byte{0x0e, 0x71}, // starts with Se
	PKHSchnorrAddrID:     [2]byte{0x0e, 0x53}, // starts with SS
	ScriptHashAddrID:     [2]byte{0x0e, 0x6c}, // starts with Sc
	PrivateKeyID:         [2]byte{0x23, 0x07}, // starts with Ps

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x20, 0xb9, 0x03}, // starts with sprv
	HDPublicKeyID:  [4]byte{0x04, 0x20, 0xbd, 0x3d}, // starts with spub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	SLIP0044CoinType: 1,   // SLIP0044, Testnet (all coins)
	LegacyCoinType:   115, // ASCII for s, for backwards compatibility

	// PicFight PoS parameters
	MinimumStakeDiff:        20000,
	TicketPoolSize:          64,
	TicketsPerBlock:         5,
	TicketMaturity:          16,
	TicketExpiry:            384, // 6*TicketPoolSize
	CoinbaseMaturity:        16,
	SStxChangeMaturity:      1,
	TicketPoolSizeWeight:    4,
	StakeDiffAlpha:          1,
	StakeDiffWindowSize:     8,
	StakeDiffWindows:        8,
	StakeVersionInterval:    8 * 2 * 7,
	MaxFreshStakePerBlock:   20,            // 4*TicketsPerBlock
	StakeEnabledHeight:      16 + 16,       // CoinbaseMaturity + TicketMaturity
	StakeValidationHeight:   16 + (64 * 2), // CoinbaseMaturity + TicketPoolSize*2
	StakeBaseSigScript:      []byte{0xDE, 0xAD, 0xBE, 0xEF},
	StakeMajorityMultiplier: 3,
	StakeMajorityDivisor:    4,

	// PicFight organization related parameters
	//
	// Treasury address is a 3-of-3 P2SH going to a wallet with seed:
	// aardvark adroitness aardvark adroitness
	// aardvark adroitness aardvark adroitness
	// aardvark adroitness aardvark adroitness
	// aardvark adroitness aardvark adroitness
	// aardvark adroitness aardvark adroitness
	// aardvark adroitness aardvark adroitness
	// aardvark adroitness aardvark adroitness
	// aardvark adroitness aardvark adroitness
	// briefcase
	// (seed 0x0000000000000000000000000000000000000000000000000000000000000000)
	//
	// This same wallet owns the three ledger outputs for simnet.
	//
	// P2SH details for simnet treasury:
	//
	// redeemScript: 532103e8c60c7336744c8dcc7b85c27789950fc52aa4e48f895ebbfb
	// ac383ab893fc4c2103ff9afc246e0921e37d12e17d8296ca06a8f92a07fbe7857ed1d4
	// f0f5d94e988f21033ed09c7fa8b83ed53e6f2c57c5fa99ed2230c0d38edf53c0340d0f
	// c2e79c725a53ae
	//   (3-of-3 multisig)
	// Pubkeys used:
	//   SkQmxbeuEFDByPoTj41TtXat8tWySVuYUQpd4fuNNyUx51tF1csSs
	//   SkQn8ervNvAUEX5Ua3Lwjc6BAuTXRznDoDzsyxgjYqX58znY7w9e4
	//   SkQkfkHZeBbMW8129tZ3KspEh1XBFC1btbkgzs6cjSyPbrgxzsKqk
	//
	// Organization address is ScuQxvveKGfpG1ypt6u27F99Anf7EW3cqhq
	OrganizationPkScript:        hexDecode("a914cbb08d6ca783b533b2c7d24a51fbca92d937bf9987"),
	OrganizationPkScriptVersion: 0,
	BlockOneLedger:              BlockOneLedgerSimNet,
}
