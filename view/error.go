package view

import (
	"errors"
)

var (
	InvalidUsrChoice = errors.New("enter either " + Accept + " or " + Deny)
)
