package chaincfg

import (
	"github.com/jfixby/coin"
	"math"
)

type DecredSubsidyCalculator struct {
	Net *Params
}

func (c DecredSubsidyCalculator) ExpectedTotalNetworkSubsidy() coin.Amount {
	return coin.Amount{2103834590794301}
	// value received by block-by-block testing
}

func (DecredSubsidyCalculator) NumberOfGeneratingBlocks() int64 {
	return math.MaxInt64
}

func (c DecredSubsidyCalculator) PreminedCoins() coin.Amount {
	return coin.Amount{c.Net.BlockOneSubsidy()}
}

func (c DecredSubsidyCalculator) CalcBlockWorkSubsidy(height int64, voters uint16, StakeValidationHeight int64) int64 {
	subsidy := c.CalcBlockSubsidy(height)

	proportionWork := int64(c.Net.WorkRewardProportion)
	proportions := int64(c.Net.TotalSubsidyProportions())
	subsidy *= proportionWork
	subsidy /= proportions

	// Ignore the voters field of the header before we're at a point
	// where there are any voters.
	if height < c.Net.StakeValidationHeight {
		return subsidy
	}

	// If there are no voters, subsidy is 0. The block will fail later anyway.
	if voters == 0 {
		return 0
	}

	// Adjust for the number of voters. This shouldn't ever overflow if you start
	// with 50 * 10^8 Atoms and voters and potentialVoters are uint16.
	potentialVoters := c.Net.TicketsPerBlock
	actual := (int64(voters) * subsidy) / int64(potentialVoters)

	return actual
}

func (c DecredSubsidyCalculator) CalcStakeVoteSubsidy(height int64) int64 {
	// Calculate the actual reward for this block, then further reduce reward
	// proportional to StakeRewardProportion.
	// Note that voters/potential voters is 1, so that vote reward is calculated
	// irrespective of block reward.
	subsidy := c.CalcBlockSubsidy(height)

	proportionStake := int64(c.Net.StakeRewardProportion)
	proportions := int64(c.Net.TotalSubsidyProportions())
	subsidy *= proportionStake
	subsidy /= (proportions * int64(c.Net.TicketsPerBlock))

	return subsidy
}

func (DecredSubsidyCalculator) FirstGeneratingBlockIndex() int64 {
	// 0 - genesis block
	// 1 - premine block
	// and the
	return 2 // - is the first generating block
}

func (c DecredSubsidyCalculator) CalcBlockTaxSubsidy(height int64, voters uint16, StakeValidationHeight int64) int64 {
	if c.Net.BlockTaxProportion == 0 {
		return 0
	}

	subsidy := c.CalcBlockSubsidy(height)

	proportionTax := int64(c.Net.BlockTaxProportion)
	proportions := int64(c.Net.TotalSubsidyProportions())
	subsidy *= proportionTax
	subsidy /= proportions

	// Assume all voters 'present' before stake voting is turned on.
	if height < c.Net.StakeValidationHeight {
		voters = 5
	}

	// If there are no voters, subsidy is 0. The block will fail later anyway.
	if voters == 0 && height >= c.Net.StakeValidationHeight {
		return 0
	}

	// Adjust for the number of voters. This shouldn't ever overflow if you start
	// with 50 * 10^8 Atoms and voters and potentialVoters are uint16.
	potentialVoters := c.Net.TicketsPerBlock
	adjusted := (int64(voters) * subsidy) / int64(potentialVoters)

	return adjusted
}

func (c DecredSubsidyCalculator) CalcBlockSubsidy(height int64) int64 {
	// Block height 1 subsidy is 'special' and used to
	// distribute initial tokens, if any.
	if height == 1 {
		return c.Net.BlockOneSubsidy()
	}

	iteration := uint64(height / c.Net.SubsidyReductionInterval)

	if iteration == 0 {
		return c.Net.BaseSubsidy
	}

	if subsidyCache == nil {
		subsidyCache = make(map[uint64]int64)
	}

	// First, check the cache.
	cachedValue, existsInCache := subsidyCache[iteration]
	if existsInCache {
		return cachedValue
	}

	// Is the previous one in the cache? If so, calculate
	// the subsidy from the previous known value and store
	// it in the database and the cache.
	cachedValue, existsInCache = subsidyCache[iteration-1]
	if existsInCache {
		cachedValue *= c.Net.MulSubsidy
		cachedValue /= c.Net.DivSubsidy

		subsidyCache[iteration] = cachedValue

		return cachedValue
	}

	// Calculate the subsidy from scratch and store in the
	// cache. TODO If there's an older item in the cache,
	// calculate it from that to save time.
	subsidy := c.Net.BaseSubsidy
	for i := uint64(0); i < iteration; i++ {
		subsidy *= c.Net.MulSubsidy
		subsidy /= c.Net.DivSubsidy
	}

	subsidyCache[iteration] = subsidy

	return subsidy
}

var subsidyCache map[uint64]int64
