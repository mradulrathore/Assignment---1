package error

import "errors"

var (
	NegativeAgeErr    = errors.New("age no can not be negative")
	EmptyFullName     = errors.New("name can not be empty")
	NegativeRollNoErr = errors.New("roll no can not be negative")
	EmptyRollNo       = errors.New("rollno can not be empty")
	EmptyAddr         = errors.New("address can not be empty")
)
