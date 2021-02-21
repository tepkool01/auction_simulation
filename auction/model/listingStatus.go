package model

import "time"

type ListingStatus struct {
	ItemListing *ItemListing
	HighestBidder *ItemBid
	BidPrice int
	StartTime int
}

//func NewListingStatus(listing *ItemListing, bidder *ItemBid) *ListingStatus {
//	return &ListingStatus{
//		ItemListing:   listing,
//		HighestBidder: bidder,
//		StartTime:     int(time.Now().Unix()),
//	}
//}

func (s *ListingStatus) SetHighestBidder(bidder *ItemBid) {
	s.HighestBidder = bidder
}

func (s *ListingStatus) SetBidPrice(price int) {
	s.BidPrice = price
}

func (s *ListingStatus) IsClosed() bool {
	return time.Now().Unix() > int64(s.StartTime + s.getTimeOfAuction())
}

func (s *ListingStatus) getTimeOfAuction() int {
	if s.ItemListing == nil {
		return 0
	}
	return s.ItemListing.GetTimeOfAuction()
}