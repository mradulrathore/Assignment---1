package view

// cmd go test -coverprofile=coverage.out

import (
	"testing"

	custErr "github.com/mradulrathore/onboarding-assignments/error"
)

func TestValidateConfirmation(t *testing.T) {

	var tests = []struct {
		userChoice string
		err        error
	}{
		{userChoice: "y", err: nil},
		{userChoice: "n", err: nil},
		{userChoice: "er", err: custErr.InvalidUsrChoice},
		{userChoice: "yes", err: custErr.InvalidUsrChoice},
		{userChoice: "no", err: custErr.InvalidUsrChoice},
	}

	for _, tc := range tests {
		if err := ValidateConfirmation(tc.userChoice); err != tc.err {
			t.Errorf("got: %v, expected: %v", err, tc.err)
		}
	}
}
