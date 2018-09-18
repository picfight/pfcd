// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2015-2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package blockchain

import (
	"testing"

	"github.com/picfight/pfcd/chaincfg"
	"fmt"
)

func TestBlockSubsidy(t *testing.T) {
	mainnet := &chaincfg.MainNetParams
	subsidyCache := NewSubsidyCache(0, mainnet)

	totalSubsidy := mainnet.BlockOneSubsidy()
	for i := int64(0); ; i++ {
		// Genesis block or first block.
		if i == 0 || i == 1 {
			continue
		}
		fmt.Printf("block %d", i)
		fmt.Println()

		if i%mainnet.SubsidyReductionInterval == 0 {
			numBlocks := mainnet.SubsidyReductionInterval
			// First reduction internal, which is reduction interval - 2
			// to skip the genesis block and block one.
			if i == mainnet.SubsidyReductionInterval {
				numBlocks -= 2
			}
			height := i - numBlocks

			work := CalcBlockWorkSubsidy(subsidyCache, height,
				mainnet.TicketsPerBlock, mainnet)
			stake := CalcStakeVoteSubsidy(subsidyCache, height,
				mainnet) * int64(mainnet.TicketsPerBlock)
			dev, art := CalcBlockTaxSubsidy(subsidyCache, height,
				mainnet.TicketsPerBlock, mainnet)
			if (i > 60*24*356*5) { // 5Yrs
				break
			}

			inc := (work + stake + dev + art)
			add := inc * numBlocks
			totalSubsidy += add

			// First reduction internal, subtract the stake subsidy for
			// blocks before the staking system is enabled.
			if i == mainnet.SubsidyReductionInterval {
				totalSubsidy -= stake * (mainnet.StakeValidationHeight - 2)
			}
		}
	}

	if totalSubsidy != 2099999999800912 {
		t.Errorf("Bad total subsidy; want 2099999999800912, got %v", totalSubsidy)
	}
}
