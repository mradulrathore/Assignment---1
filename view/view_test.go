package view

// cmd go test -coverprofile=coverage.out

import (
	"testing"
)

func TestValidateConfirmation(t *testing.T) {

	userChoices := []string{"y", "n", "er", "yes", "no"}

	expected := []bool{true, true, false, false, false}

	for index, choice := range userChoices {
		if ok, err := ValidateConfirmation(choice); ok != expected[index] {
			if !expected[index] {
				t.Errorf("exception is occuring: %q", err)
			} else {
				t.Errorf("exception is not occuring: %q", err)
			}

		}
	}
}
