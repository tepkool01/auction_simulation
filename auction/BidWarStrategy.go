package auction

import "lookout_interview/auction/model"

// BidWarStrategy is the interface for various types of strategies in the event an item results in a bidding war
type BidWarStrategy interface {
	Fight(residingWinner *model.ItemBid, opponent *model.ItemBid, item *model.ItemListing)
	GetWinner() *model.ItemBid
	GetHighestBid() int
}
