package auction

import (
	"lookout_interview/auction/model"
)

// KingOfTheHillStrategy is a strategy based on the assignment constraints indicating that the residing winner of an item will win in ties.
type KingOfTheHillStrategy struct {
	Winner     *model.ItemBid
	HighestBid int
}

// Fight under the condition of KingOfTheHill is a bid-war that will push user's bids to their opponent's max to maintain/win the item
func (k *KingOfTheHillStrategy) Fight(residingWinner *model.ItemBid, opponent *model.ItemBid, item *model.ItemListing) {
	// Computing the true max bid under the idea that a max bid might not be obtainable under certain increments
	opponentTrueMax := k.computeTrueMax(opponent)
	residingTrueMax := k.computeTrueMax(residingWinner)

	// The bidder can't beat the existing user, and made the residing user increase their cost basis for this item
	if residingTrueMax > opponentTrueMax {
		increasedBid := residingWinner.GetStartingBid()
		// Iterating to find out the minimum amount to beat the opponent, starting at the startingBid to get an increment that would reflect the initial conditions
		for increasedBid < opponentTrueMax {
			increasedBid += residingWinner.GetBidIncrement()
		}
		k.Winner = residingWinner
		k.HighestBid = increasedBid
		return
	}
	// A new winner is found, because the opponent can (and does) outbid the current holder
	if residingTrueMax < opponentTrueMax {
		opponentBid := opponent.GetStartingBid()
		for opponentBid < residingTrueMax {
			opponentBid += opponent.GetBidIncrement()
		}
		k.Winner = opponent
		k.HighestBid = opponentBid
		return
	}

	// The bidder just pushed the user to their highest bid, so it remains as the existing winner, no change
	k.Winner = residingWinner
	k.HighestBid = opponentTrueMax
}

// GetWinner returns the winner of the bid contention
func (k *KingOfTheHillStrategy) GetWinner() *model.ItemBid {
	return k.Winner
}

// GetHighestBid returns the highest bid, after the bidding war dust settles
func (k *KingOfTheHillStrategy) GetHighestBid() int {
	return k.HighestBid
}

// computeTrueMax resolves the 'true' max bid, because some max bids are not obtainable with certain increments, i.e. max bid of 10 with increment of 3
func (k *KingOfTheHillStrategy) computeTrueMax(user *model.ItemBid) int {
	remainder := (user.MaxBid - user.StartingBid) % user.BidIncrement
	if remainder == 0 {
		return user.MaxBid
	}
	return user.MaxBid - remainder
}
