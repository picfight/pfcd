// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2015-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package blockchain

import (
	"bytes"
	"fmt"

	"github.com/picfight/pfcd/blockchain/stake"
	"github.com/picfight/pfcd/chaincfg"
	"github.com/picfight/pfcd/pfcutil"
	"github.com/picfight/pfcd/txscript"
	"github.com/picfight/pfcd/wire"
)

// The number of values to precalculate on initialization of the subsidy
// cache.
const subsidyCacheInitWidth = 4

// SubsidyCache is a structure that caches calculated values of subsidy so that
// they're not constantly recalculated. The blockchain struct itself possesses a
// pointer to a preinitialized SubsidyCache.
type SubsidyCache struct {
	params *chaincfg.Params
}

// NewSubsidyCache initializes a new subsidy cache for a given height. It
// precalculates the values of the subsidy that are most likely to be seen by
// the client when it connects to the network.
func NewSubsidyCache(height int64, params *chaincfg.Params) *SubsidyCache {
	sc := SubsidyCache{
		params: params,
	}

	iteration := uint64(height / params.SubsidyReductionInterval)
	if iteration < subsidyCacheInitWidth {
		return &sc
	}

	for i := iteration - 4; i <= iteration; i++ {
		sc.CalcBlockSubsidy(int64(iteration) * params.SubsidyReductionInterval)
	}

	return &sc
}

// CalcBlockSubsidy returns the subsidy amount a block at the provided height
// should have. This is mainly used for determining how much the coinbase for
// newly generated blocks awards as well as validating the coinbase for blocks
// has the expected value.
func (s *SubsidyCache) CalcBlockSubsidy(height int64) int64 {
	// Block height 1 subsidy is 'special' and used to
	// distribute initial tokens, if any.
	if height == 1 {
		return s.params.BlockOneSubsidy()
	}

	return s.params.BaseSubsidy
}

// CalcBlockWorkSubsidy calculates the proof of work subsidy for a block as a
// proportion of the total subsidy.
func CalcBlockWorkSubsidy(subsidyCache *SubsidyCache, height int64, voters uint16, params *chaincfg.Params) int64 {
	subsidy := subsidyCache.CalcBlockSubsidy(height)

	proportionWork := int64(params.WorkRewardProportion)
	proportions := int64(params.TotalSubsidyProportions())
	subsidy *= proportionWork
	subsidy /= proportions

	// Ignore the voters field of the header before we're at a point
	// where there are any voters.
	if height < params.StakeValidationHeight {
		return subsidy
	}

	// If there are no voters, subsidy is 0. The block will fail later anyway.
	if voters == 0 {
		return 0
	}

	// Adjust for the number of voters. This shouldn't ever overflow if you start
	// with 50 * 10^8 Atoms and voters and potentialVoters are uint16.
	potentialVoters := params.TicketsPerBlock
	actual := (int64(voters) * subsidy) / int64(potentialVoters)

	return actual
}

// CalcStakeVoteSubsidy calculates the subsidy for a stake vote based on the height
// of its input SStx.
//
// Safe for concurrent access.
func CalcStakeVoteSubsidy(subsidyCache *SubsidyCache, height int64, params *chaincfg.Params) int64 {
	// Calculate the actual reward for this block, then further reduce reward
	// proportional to StakeRewardProportion.
	// Note that voters/potential voters is 1, so that vote reward is calculated
	// irrespective of block reward.
	subsidy := subsidyCache.CalcBlockSubsidy(height)

	proportionStake := int64(params.StakeRewardProportion)
	proportions := int64(params.TotalSubsidyProportions())
	subsidy *= proportionStake
	subsidy /= (proportions * int64(params.TicketsPerBlock))

	return subsidy
}

// CalcBlockTaxSubsidy calculates the subsidy for the organization address in the
// coinbase.
func CalcBlockTaxSubsidy(subsidyCache *SubsidyCache, height int64, voters uint16, params *chaincfg.Params) (dev, art int64) {
	if params.BlockDevTaxProportion == 0 || params.BlockArtTaxProportion == 0 {
		return 0, 0
	}

	subsidyTotal := subsidyCache.CalcBlockSubsidy(height)

	proportionTaxDev := int64(params.BlockDevTaxProportion)
	proportionTaxArt := int64(params.BlockArtTaxProportion)
	proportions := int64(params.TotalSubsidyProportions())

	subsidyDev := subsidyTotal * (proportionTaxDev)
	subsidyDev /= proportions

	subsidyArt := subsidyTotal * (proportionTaxArt)
	subsidyArt /= proportions

	// Assume all voters 'present' before stake voting is turned on.
	if height < params.StakeValidationHeight {
		voters = 5
	}

	// If there are no voters, subsidy is 0. The block will fail later anyway.
	if voters == 0 && height >= params.StakeValidationHeight {
		return 0, 0
	}

	// Adjust for the number of voters. This shouldn't ever overflow if you start
	// with 50 * 10^8 Atoms and voters and potentialVoters are uint16.
	potentialVoters := params.TicketsPerBlock
	adjustedDev := (int64(voters) * subsidyDev) / int64(potentialVoters)
	adjustedArt := (int64(voters) * subsidyArt) / int64(potentialVoters)

	return adjustedDev, adjustedArt
}

// BlockOneCoinbasePaysTokens checks to see if the first block coinbase pays
// out to the network initial token ledger.
func BlockOneCoinbasePaysTokens(tx *pfcutil.Tx, params *chaincfg.Params) error {
	// If no ledger is specified, just return true.
	if len(params.BlockOneLedger) == 0 {
		return nil
	}

	if tx.MsgTx().LockTime != 0 {
		errStr := fmt.Sprintf("block 1 coinbase has invalid locktime")
		return ruleError(ErrBlockOneTx, errStr)
	}

	if tx.MsgTx().Expiry != wire.NoExpiryValue {
		errStr := fmt.Sprintf("block 1 coinbase has invalid expiry")
		return ruleError(ErrBlockOneTx, errStr)
	}

	if tx.MsgTx().TxIn[0].Sequence != wire.MaxTxInSequenceNum {
		errStr := fmt.Sprintf("block 1 coinbase not finalized")
		return ruleError(ErrBlockOneInputs, errStr)
	}

	if len(tx.MsgTx().TxOut) == 0 {
		errStr := fmt.Sprintf("coinbase outputs empty in block 1")
		return ruleError(ErrBlockOneOutputs, errStr)
	}

	ledger := params.BlockOneLedger
	if len(ledger) != len(tx.MsgTx().TxOut) {
		errStr := fmt.Sprintf("wrong number of outputs in block 1 coinbase; "+
			"got %v, expected %v", len(tx.MsgTx().TxOut), len(ledger))
		return ruleError(ErrBlockOneOutputs, errStr)
	}

	// Check the addresses and output amounts against those in the ledger.
	for i, txout := range tx.MsgTx().TxOut {
		if txout.Version != txscript.DefaultScriptVersion {
			errStr := fmt.Sprintf("bad block one output version; want %v, got %v",
				txscript.DefaultScriptVersion, txout.Version)
			return ruleError(ErrBlockOneOutputs, errStr)
		}

		// There should only be one address.
		_, addrs, _, err :=
			txscript.ExtractPkScriptAddrs(txout.Version, txout.PkScript, params)
		if err != nil {
			return ruleError(ErrBlockOneOutputs, err.Error())
		}
		if len(addrs) != 1 {
			errStr := fmt.Sprintf("too many addresses in output")
			return ruleError(ErrBlockOneOutputs, errStr)
		}

		addrLedger, err := pfcutil.DecodeAddress(ledger[i].Address)
		if err != nil {
			return err
		}

		if !bytes.Equal(addrs[0].ScriptAddress(), addrLedger.ScriptAddress()) {
			errStr := fmt.Sprintf("address in output %v has non matching "+
				"address; got %v (hash160 %x), want %v (hash160 %x)",
				i,
				addrs[0].EncodeAddress(),
				addrs[0].ScriptAddress(),
				addrLedger.EncodeAddress(),
				addrLedger.ScriptAddress())
			return ruleError(ErrBlockOneOutputs, errStr)
		}

		if txout.Value != ledger[i].Amount {
			errStr := fmt.Sprintf("address in output %v has non matching "+
				"amount; got %v, want %v", i, txout.Value, ledger[i].Amount)
			return ruleError(ErrBlockOneOutputs, errStr)
		}
	}

	return nil
}

// CoinbasePaysTax checks to see if a given block's coinbase correctly pays
// tax to the developer organization.
func CoinbasePaysTax(subsidyCache *SubsidyCache, tx *pfcutil.Tx, height int64, voters uint16, params *chaincfg.Params) error {
	// Taxes only apply from block 2 onwards.
	if height <= 1 {
		return nil
	}

	// Tax is disabled.
	if params.BlockDevTaxProportion == 0 || params.BlockArtTaxProportion == 0 {
		return nil
	}

	if len(tx.MsgTx().TxOut) == 0 {
		errStr := fmt.Sprintf("invalid coinbase (no outputs)")
		return ruleError(ErrNoTxOutputs, errStr)
	}
	{ // developers
		outputID := 0
		taxOutput := tx.MsgTx().TxOut[outputID]
		if taxOutput.Version != params.OrganizationDevelopersPkScriptVersion {
			return ruleError(ErrNoTax,
				"coinbase tax output uses incorrect script version")
		}
		if !bytes.Equal(taxOutput.PkScript, params.OrganizationDevelopersPkScript) {
			return ruleError(ErrNoTax,
				"coinbase tax output script does not match the "+
					"required script")
		}

		// Get the amount of subsidy that should have been paid out to
		// the organization, then check it.
		dev, _ := CalcBlockTaxSubsidy(subsidyCache, height, voters, params)
		outputName := fmt.Sprintf("Developers[%v]", outputID)
		if dev != taxOutput.Value {
			errStr := fmt.Sprintf("amount in output %s has non matching org "+
				"calculated amount; got %v, want %v",
				outputName,
				taxOutput.Value,
				dev)
			return ruleError(ErrNoTax, errStr)
		}
	}
	{ // artists
		outputID := 1
		taxOutput := tx.MsgTx().TxOut[outputID]
		if taxOutput.Version != params.OrganizationArtistsPkScriptVersion {
			return ruleError(ErrNoTax,
				"coinbase tax output uses incorrect script version")
		}
		if !bytes.Equal(taxOutput.PkScript, params.OrganizationArtistsPkScript) {
			return ruleError(ErrNoTax,
				"coinbase tax output script does not match the "+
					"required script")
		}

		// Get the amount of subsidy that should have been paid out to
		// the organization, then check it.
		_, art := CalcBlockTaxSubsidy(subsidyCache, height, voters, params)
		outputName := fmt.Sprintf("Artists[%v]", outputID)
		if art != taxOutput.Value {
			errStr := fmt.Sprintf("amount in output %s has non matching org "+
				"calculated amount; got %v, want %v",
				outputName,
				taxOutput.Value,
				art)
			return ruleError(ErrNoTax, errStr)
		}
	}

	return nil
}

// CalculateAddedSubsidy calculates the amount of subsidy added by a block
// and its parent. The blocks passed to this function MUST be valid blocks
// that have already been confirmed to abide by the consensus rules of the
// network, or the function might panic.
func CalculateAddedSubsidy(block, parent *pfcutil.Block) int64 {
	var subsidy int64
	if headerApprovesParent(&block.MsgBlock().Header) {
		subsidy += parent.MsgBlock().Transactions[0].TxIn[0].ValueIn
	}

	for _, stx := range block.MsgBlock().STransactions {
		if stake.IsSSGen(stx) {
			subsidy += stx.TxIn[0].ValueIn
		}
	}

	return subsidy
}
