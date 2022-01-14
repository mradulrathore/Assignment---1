package enum

//go:generate enumer -type=ItemType

type ItemType int

const (
	Raw ItemType = iota
	Manufactured
	Imported
)
