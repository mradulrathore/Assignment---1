package error

import (
	"errors"

	"github.com/mradulrathore/onboarding-assignments/constant"
)

var (
	NoItmErr         = errors.New("pleae specify item type")
	NegativeQuantErr = errors.New("quantity can not be negative")
	NegativePriceErr = errors.New("price can not be negative")
	InvalideItmType  = errors.New("item type can only be raw, manufactured or imported")
	InvalidUsrChoice = errors.New("enter either " + constant.Accept + " or " + constant.Deny)
)
