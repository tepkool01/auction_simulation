package auction

import "lookout_interview/auction/model"

type BidWarStrategy interface {
	Fight(residingWinner *model.ItemBid, opponent *model.ItemBid, currentBid int)
	GetWinner() *model.ItemBid
	GetNewBidAmount() int
}