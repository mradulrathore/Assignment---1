package ports

import (
	"errors"
)

var (
	InvalidUsrChoice   = errors.New("enter either " + Accept + " or " + Deny)
	DuplicateCourseErr = errors.New("duplicate course")
)
