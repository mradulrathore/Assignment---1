package item

import (
	"errors"
)

var (
	NegativeQuantErr = errors.New("quantity can not be negative")
	NegativePriceErr = errors.New("price can not be negative")
)
