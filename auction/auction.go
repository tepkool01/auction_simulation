package auction

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"lookout_interview/auction/model"
	"time"
)

type Auction struct {
	Items []model.Item
	IngestionRate int
	ListingStatuses map[string]*model.ListingStatus // a map to hold the current status of the auction
}

// NewAuction instantiates a new auction with data from a JSON file
func NewAuction(filePath string, ingestionRate int) *Auction {
	auction := new(Auction)
	auction.SetIngestionRate(ingestionRate)

	rawItemArray, err := getArrayFromFile(filePath)
	if err != nil {
		return nil
	}

	// Iterating over the JSON data to create a more usable interaction with the items via defined structs
	var auctionItemType model.Item
	for _, item := range rawItemArray {
		auctionItemType = model.ItemBid{}
		if item["type"] == model.LISTING.String() {
			fmt.Println("HEYYYYYYYYY")
			auctionItemType = model.ItemListing{}
		}
		err = mapstructure.Decode(item, &auctionItemType)
		if err != nil {
			log.Println("skipping item due to decoding failure")
			continue
		}
		auction.Items = append(auction.Items, auctionItemType)
	}
	fmt.Println(auction.Items)

	return auction
}

func (a *Auction) Run() {
	// Initializing the map to prevent nil map errors
	a.ListingStatuses = make(map[string]*model.ListingStatus)

	for _, item := range a.GetItems() {
		executionStart := time.Now()

		// Checking the most likely model first (bid)
		fmt.Println(item.GetType())
		if itemBid, ok := item.(model.ItemBid); ok {
			// todo: break out winning strategy into function or somewhere else


			// Check if item has been added to the map and the bidding is still available for that item
			if itemListing, ok := a.ListingStatuses[itemBid.GetItem()]; ok && !a.ListingStatuses[itemBid.GetItem()].IsClosed() {
				if itemListing.HighestBidder == nil {
					// No one has bid for this item, if they want it, they can have it
					itemListing.HighestBidder = &itemBid
					itemListing.SetBidPrice(itemBid.GetStartingBid())
				} else {
					fmt.Println("WARRRRRRRRRRRRRRRRRRRRR")
					// Execute bid war
					// todo: possibly use the strategy pattern here for altering the true max bid
					volatileBidPrice := itemListing.BidPrice // the final bid price that will be observed as the winning bid
					opponentBid := itemBid.GetStartingBid() // the starting point for the opponent to increment from
					previousWinnerBid := itemListing.BidPrice // using saved bid price to start from, because they had to get there somehow

					for volatileBidPrice < itemBid.GetMaxBid() && volatileBidPrice < itemListing.HighestBidder.GetMaxBid() {
						// Opponent (current user) trying to outbid the current price
						for opponentBid <= volatileBidPrice {
							opponentBid += itemBid.GetBidIncrement()
						}
						volatileBidPrice = opponentBid


						for previousWinnerBid <= volatileBidPrice {
							previousWinnerBid += itemListing.HighestBidder.GetBidIncrement()
						}
						volatileBidPrice = previousWinnerBid
					}
					// Figure out the winner based on who is maxed out
					if itemBid.GetMaxBid() == itemListing.HighestBidder.GetMaxBid() {
						// The bidder just pushed the user to their highest bid, so it remains as the existing winner, no change
					} else if itemListing.HighestBidder.GetMaxBid() > itemBid.GetMaxBid() {
						// The bidder can't beat the existing user, and just made him pay more, update the bid price for this item
						itemListing.SetBidPrice(volatileBidPrice)
					} else {
						// A new winner is found!
						itemListing.SetBidPrice(volatileBidPrice)
						itemListing.SetHighestBidder(&itemBid)
					}


				}
			}





		} else {
			// Type assertion check for the ItemListing model
			if itemListing, ok := item.(model.ItemListing); ok {
				a.ListingStatuses[itemListing.GetID()] = &model.ListingStatus{
					ItemListing:   &itemListing,
					HighestBidder: nil,
					StartTime:     int(time.Now().Unix()),
				}
			}
		}


		// Factor in the above execution time to make the iterations consistently spaced
		executionDuration := time.Since(executionStart)
		sleepTime := a.GetIngestionRate() - int(executionDuration / 1000000) // convert nano seconds to milliseconds for execution duration
		if sleepTime > 0 { // make sure we don't sleep a negative number
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		}
		fmt.Println(time.Now())
	}
	fmt.Println("Results")
	a.PrintResults()
}

// todo: move this to an 'OutPuts' function that pretty prints results
func (a *Auction) PrintResults() {

	for _, item := range a.GetListingStatuses() {
		fmt.Printf("%s - %d - %s\n", item.ItemListing.GetName(), item.BidPrice, item.HighestBidder.GetName())
	}
}

func (a *Auction) GetListingStatuses() map[string]*model.ListingStatus {
	return a.ListingStatuses
}

func (a *Auction) SetIngestionRate(milliseconds int) {
	a.IngestionRate = milliseconds
}

func (a *Auction) GetIngestionRate() int {
	return a.IngestionRate
}

func (a *Auction) GetItems() []model.Item {
	return a.Items
}