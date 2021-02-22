package model

// ItemType is the enumeration for item types
type ItemType int

// BID and LISTING are the enum types that we accept
const (
	BID ItemType = iota
	LISTING
)

// String returns the string representation of the enum
func (t ItemType) String() string {
	return [...]string{"bid", "newItem"}[t]
}
