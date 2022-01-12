package enum

//go:generate enumer -type=Course -json

const (
	A Course = iota
	B
	C
	D
	E
	F
)

type Course int
