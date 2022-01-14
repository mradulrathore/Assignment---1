package view

import (
	"os"
	"testing"
)

//TODO
func TestInit(t *testing.T) {

	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{{
		scenario: "add user",
		err:      nil,
	}, {
		scenario: "delete by rollno",
		err:      nil,
	}, {
		scenario: "save user",
		err:      nil,
	}, {
		scenario: "confirm validation",
		err:      nil,
	}}

	_ = tests
}
