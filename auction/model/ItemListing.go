package model

// ItemListing holds the information for the items that will be bid on, validation is contained in the properties
type ItemListing struct {
	ID            string `json:"id" validate:"ascii,required"`
	Type          string `json:"type" validate:"ascii,required,oneof=bid newItem"`
	Name          string `json:"name" validate:"ascii,required,min=1,max=50"`
	Description   string `json:"description"`
	TimeOfAuction int    `json:"timeOfAuction" validate:"ascii,required,numeric,min=1"`
}

// GetID returns the ID of the item
func (a ItemListing) GetID() string {
	return a.ID
}

// GetType will return newItem for this struct
func (a ItemListing) GetType() string {
	return a.Type
}

// GetName returns the name of the item
func (a ItemListing) GetName() string {
	return a.Name
}

// GetTimeOfAuction returns how long the auction will last
func (a ItemListing) GetTimeOfAuction() int {
	return a.TimeOfAuction
}
