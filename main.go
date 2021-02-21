package main

import (
	"flag"
	Auction "lookout_interview/auction"
)

const DefaultAuctionItems = "auctionItems.json"
const DefaultIngestionRate = 1000 // DefaultIngestionRate is in milliseconds; 1000 ms is 1 second

func main() {
	// Grab the flags from the CLI
	filePath := flag.String("path", DefaultAuctionItems, "file path to auction items JSON, i.e. auctionItems.json")
	ingestionRate := flag.Int("milliseconds", DefaultIngestionRate, "items to ingest per millisecond, i.e. 1000, 100, 5000")
	flag.Parse()

	auction := Auction.NewAuction(*filePath, *ingestionRate)
	auction.Run()
	//fmt.Println(auction.GetItems())
	//for _, val := range auctionItems {
	//	fmt.Println(val.GetID())
	//}

}

//func strDereference(str *string) string {
//	if str == nil {
//		return ""
//	}
//	return *str
//}
//
//func floatDereference(num *float64) float64 {
//	if num == nil {
//		return 1.0
//	}
//	return *num
//}