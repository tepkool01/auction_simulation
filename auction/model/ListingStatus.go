package model

import "time"

// ListingStatus provides the current bidding status for a single item
type ListingStatus struct {
	ItemListing   *ItemListing
	HighestBidder *ItemBid
	BidPrice      int
	StartTime     int
}

// SetHighestBidder sets the user object for highest bidder
func (s *ListingStatus) SetHighestBidder(bidder *ItemBid) {
	s.HighestBidder = bidder
}

// SetBidPrice sets the current highest bid price for this item
func (s *ListingStatus) SetBidPrice(price int) {
	s.BidPrice = price
}

// IsClosed determines if the bidding for a particular item has been closed by expiration
func (s *ListingStatus) IsClosed() bool {
	return time.Now().Unix() > int64(s.StartTime+s.getTimeOfAuction())
}

func (s *ListingStatus) getTimeOfAuction() int {
	if s.ItemListing == nil {
		return 0
	}
	return s.ItemListing.GetTimeOfAuction()
}
