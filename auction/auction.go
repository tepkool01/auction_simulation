package auction

import (
	"log"
	"time"

	"lookout_interview/auction/model"

	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
)

// Auction is the implementation of running an auction and observing the results
type Auction struct {
	Items           []model.Item                    // array of multi-type items in the JSON document, i.e. an auction item and bidder
	IngestionRate   int                             // IngestionRate is a
	ListingStatuses map[string]*model.ListingStatus // ListingStatuses is a map to hold the current status of the auction
}

// NewAuction instantiates a new auction with data from a JSON file and prepares entries as specific types
func NewAuction(filePath string, ingestionRate int) (*Auction, error) {
	auction := new(Auction)
	auction.SetIngestionRate(ingestionRate)

	// Validator which checks for valid structs
	v := validator.New()

	// Retrieves the JSON array from the file as an anonymous type (interface) to be parsed later
	rawItemArray, err := getArrayFromFile(filePath)
	if err != nil {
		log.Println("could not extract data from file. malformed or non-existent file")
		return nil, err
	}

	// Iterating over the JSON data to create a more usable interaction with the items via defined structs
	var auctionItemType model.Item
	for _, item := range rawItemArray {
		// Switching type based on type property to place it into the correct struct
		auctionItemType = model.ItemBid{}
		if item["type"] == model.LISTING.String() {
			auctionItemType = model.ItemListing{}
		}
		_ = mapstructure.Decode(item, &auctionItemType)

		// Validate struct is valid before adding it to the slice
		err := v.Struct(auctionItemType)
		if err != nil {
			log.Println("item does not conform to prescribed specification")
			return nil, err
		}

		// Add the items to the list!
		auction.Items = append(auction.Items, auctionItemType)
	}

	return auction, nil
}

// Run starts the auction with a given bid war strategy
func (a *Auction) Run(bidWarStrategy BidWarStrategy) {
	// Initializing the map to prevent nil map errors
	a.ListingStatuses = make(map[string]*model.ListingStatus)

	// Iterate over all of the items in the auction, at the set ingestion rate, and aggregate the results
	for _, item := range a.GetItems() {
		// Starting a timer to time the execution of the following block, which will be used to calculate the sleeping rate
		executionStart := time.Now()

		// Checking the most likely model first (bid)
		if itemBid, ok := item.(model.ItemBid); ok {
			// Check if item has been added to the map and the bidding is still available for that item
			if itemListing, ok := a.ListingStatuses[itemBid.GetItem()]; ok && !a.ListingStatuses[itemBid.GetItem()].IsClosed() {
				if itemListing.HighestBidder == nil {
					// No one has bid for this item, if they want it, they can have it
					itemListing.HighestBidder = &itemBid
					itemListing.SetBidPrice(itemBid.GetStartingBid())

					printBidAction(&itemBid, itemListing.ItemListing, itemBid.GetStartingBid())
				} else {
					// Executing the bid war based on the selected strategy
					bidWarStrategy.Fight(itemListing.HighestBidder, &itemBid, itemListing.ItemListing)

					// Updating the bid status of this item with the new winner and bid price
					itemListing.SetHighestBidder(bidWarStrategy.GetWinner())
					itemListing.SetBidPrice(bidWarStrategy.GetHighestBid())

					// Let's indicate to the console that there is a new price/winner for this item
					printBidWarResult(bidWarStrategy.GetWinner(), itemListing.ItemListing, bidWarStrategy.GetHighestBid())
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
				log.Printf("New Item available for bidding: %s (%s)", itemListing.GetName(), itemListing.GetID())
			}
		}

		// Print current status each event (per the requirements)
		a.PrintResults()

		// Factor in the above execution time to make the iterations consistently spaced
		executionDuration := time.Since(executionStart)
		sleepTime := a.GetIngestionRate() - int(executionDuration/1000000) // convert nano seconds to milliseconds for execution duration
		if sleepTime > 0 {                                                 // make sure we don't sleep a negative number
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		}
	}
	log.Println("Auction concluded")
}

// GetListingStatuses returns a map of the current results of the auction
func (a *Auction) GetListingStatuses() map[string]*model.ListingStatus {
	return a.ListingStatuses
}

// SetIngestionRate sets the rate of processing items in the auction array. I.E. 1 item per 500 milliseconds
func (a *Auction) SetIngestionRate(milliseconds int) {
	a.IngestionRate = milliseconds
}

// GetIngestionRate returns the ingestion rate in milliseconds
func (a *Auction) GetIngestionRate() int {
	return a.IngestionRate
}

// GetItems returns an array of auction items (combination of ItemListing and ItemBid)
func (a *Auction) GetItems() []model.Item {
	return a.Items
}

// PrintResults sends the results of the auction to the console
func (a *Auction) PrintResults() {
	log.Println("====================AUCTION LISTINGS====================")
	for _, item := range a.GetListingStatuses() {
		// Checking for nil pointers
		bidderName := "N/A"
		bidderID := "N/A"
		if item.HighestBidder != nil {
			bidderName = item.HighestBidder.GetName()
			bidderID = item.HighestBidder.GetID()
		}
		log.Printf("Item: %s (%s) - Amount: $%d - User: %s (%s)\n",
			item.ItemListing.GetName(),
			item.ItemListing.GetID(),
			item.BidPrice,
			bidderName,
			bidderID,
		)
	}
	log.Println("========================================================")
	log.Println("")
}

func printBidWarResult(winner *model.ItemBid, item *model.ItemListing, bidPrice int) {
	log.Printf("Item Contended: %s bid $%d for the %s and is the current highest bidder!\n", winner.GetName(), bidPrice, item.GetName())
}

func printBidAction(user *model.ItemBid, item *model.ItemListing, bidPrice int) {
	log.Printf("Bid Action: %s bid $%d for the %s and is the current highest bidder!\n", user.GetName(), bidPrice, item.GetName())
}
