package enum

//go:generate enumer -type=ItemType

const (
	Raw ItemType = iota
	Manufactured
	Imported
)

type ItemType int
