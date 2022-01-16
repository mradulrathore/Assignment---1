package repository

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {

	tests := []struct {
		scenario string
		err      error
	}{{
		scenario: "open file",
		err:      nil,
	}}

	for _, tc := range tests {
		_, err := Open()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

//TODO
func TestSave(t *testing.T) {

	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{{
		scenario: "save file",
		err:      nil,
	}}

	_ = tests
}

//TODO
func TestRetrieve(t *testing.T) {

	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{{
		scenario: "retrieve file",
		err:      nil,
	}}

	_ = tests
}
