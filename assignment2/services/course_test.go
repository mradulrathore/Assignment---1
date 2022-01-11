package services

// cmd go test -coverprofile=coverage.out

import (
	"testing"
)

func TestValidateCourse(t *testing.T) {
	userChoices := []string{"A,B,C", "A,B,C,D,E", "A", "ABC", "P,Q,R", "A,A,B,B"}

	expected := []bool{false, true, false, false, false, false}

	for index, choice := range userChoices {
		if ok, err := ValidateCourse(choice); ok != expected[index] {
			if !expected[index] {
				t.Errorf("exception is occuring: %q", err)
			} else {
				t.Errorf("exception is not occuring: %q", err)
			}

		}
	}
}
