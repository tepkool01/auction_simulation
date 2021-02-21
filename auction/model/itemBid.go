package model

type ItemBid struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	ItemName string `json:"itemName"`
	Item string `json:"item"`
	StartingBid int `json:"startingBid"`
	MaxBid int `json:"maxBid"`
	BidIncrement int `json:"bidIncrement"`
}

// GetMaxBid resolves the 'true' max bid, because some max bids are not obtainable with certain increments, i.e. max bid of 10 with increment of 3
func (a ItemBid) GetMaxBid() int {
	remainder := (a.MaxBid - a.StartingBid) % a.BidIncrement
	if remainder == 0 {
		return a.MaxBid
	}
	return a.MaxBid - remainder
}

func (a ItemBid) GetBidIncrement() int {
	return a.BidIncrement
}

func (a ItemBid) GetID() string {
	return a.Id
}

func (a ItemBid) GetType() string {
	return a.Type
}

func (a ItemBid) GetName() string {
	return a.Name
}

func (a ItemBid) GetItem() string {
	return a.Item
}

func (a ItemBid) GetStartingBid() int {
	return a.StartingBid
}

