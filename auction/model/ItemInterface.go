package model

// Item is the interface for the structs that will contain the items in the JSON array
type Item interface {
	GetID() string
	GetType() string
	GetName() string
}
