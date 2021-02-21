package model

type ItemListing struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Description string `json:"description"`
	TimeOfAuction int `json:"timeOfAuction"`
}

func (a ItemListing) GetID() string {
	return a.Id
}

func (a ItemListing) GetType() string {
	return a.Type
}

func (a ItemListing) GetName() string {
	return a.Name
}

func (a ItemListing) GetTimeOfAuction() int {
	return  a.TimeOfAuction
}