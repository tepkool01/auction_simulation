package main

import (
	"flag"
	Auction "lookout_interview/auction"
)

const DefaultAuctionItems = "auctionItems.json" // DefaultAuctionItems is the location of the file to ingest, this can be a full or relative path
const DefaultIngestionRate = 1000               // DefaultIngestionRate is in milliseconds; 1000 ms is 1 second

func main() {
	// Grab the flags from the CLI to determine the file path and ingestion rate
	filePath := flag.String("path", DefaultAuctionItems, "file path to auction items JSON, i.e. auctionItems.json")
	ingestionRate := flag.Int("milliseconds", DefaultIngestionRate, "items to ingest per millisecond, i.e. 1000, 100, 5000")
	flag.Parse()

	// Leveraging a strategy that in the event of a contended item at the same price, the person with the initial bid will win
	kingOfTheHill := &Auction.KingOfTheHillStrategy{}

	// Create a new auction instance, run the auction, and observe the results in the console
	auction := Auction.NewAuction(*filePath, *ingestionRate)
	auction.Run(kingOfTheHill)
	auction.PrintResults()
}
