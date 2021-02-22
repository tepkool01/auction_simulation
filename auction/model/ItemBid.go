package model

// ItemBid struct which are similar to properties of a class. This also has the added benefit of running built-in validation
type ItemBid struct {
	ID           string `json:"id" validate:"ascii,required"`
	Type         string `json:"type" validate:"ascii,required,oneof=bid newItem"`
	Name         string `json:"name" validate:"ascii,required,min=1,max=50"`
	ItemName     string `json:"itemName" validate:"ascii,required,min=1,max=100"`
	Item         string `json:"item"`
	StartingBid  int    `json:"startingBid" validate:"ascii,required,numeric,min=1"`
	MaxBid       int    `json:"maxBid" validate:"ascii,required,numeric,gtefield=StartingBid"`
	BidIncrement int    `json:"bidIncrement" validate:"ascii,required,numeric,min=1"`
}

// GetMaxBid returns the max bid
func (a ItemBid) GetMaxBid() int {
	return a.MaxBid
}

// GetBidIncrement returns the increment the user will bid
func (a ItemBid) GetBidIncrement() int {
	return a.BidIncrement
}

// GetID retrieves the ID of the item in the array (GUID seemingly)
func (a ItemBid) GetID() string {
	return a.ID
}

// GetType returns bid (redundant?)
func (a ItemBid) GetType() string {
	return a.Type
}

// GetName returns the name of the item the user is bidding on
func (a ItemBid) GetName() string {
	return a.Name
}

// GetItem retrieves the ID of the item the user is bidding on
func (a ItemBid) GetItem() string {
	return a.Item
}

// GetStartingBid is the first offer a user will make on an item
func (a ItemBid) GetStartingBid() int {
	return a.StartingBid
}
