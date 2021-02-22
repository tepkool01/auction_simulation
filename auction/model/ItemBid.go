package model

type ItemBid struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	ItemName     string `json:"itemName"`
	Item         string `json:"item"`
	StartingBid  int    `json:"startingBid"`
	MaxBid       int    `json:"maxBid"`
	BidIncrement int    `json:"bidIncrement"`
}

func (a ItemBid) GetMaxBid() int {
	return a.MaxBid
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
