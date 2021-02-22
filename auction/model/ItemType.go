package model

type ItemType int

const (
	BID ItemType = iota
	LISTING
)

func (t ItemType) String() string {
	return [...]string{"bid", "newItem"}[t]
}