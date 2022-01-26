package vutil

import "errors"

func CheckPositive(value interface{}) error {
	val := value.(int)
	if val <= 0 {
		return errors.New("must be positive")
	}
	return nil
}
