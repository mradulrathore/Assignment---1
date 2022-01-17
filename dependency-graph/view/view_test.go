package view

import (
	"errors"
	"log"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	fileAllOptions := setInputSave("all_options_test.txt")
	fileDuplicateId := setInputSave("duplicate_id_test.txt")
	fileAddDependency := setInputSave("duplicate_id_test.txt")

	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{
		{
			scenario: "dependency graph all options test",
			req:      fileAllOptions,
			err:      nil,
		}, {
			scenario: "adding node having duplicate id test",
			req:      fileDuplicateId,
			err:      errors.New("duplicate id"),
		}, {
			scenario: "adding dependency when node doesn't exist test",
			req:      fileAddDependency,
			err:      errors.New("node doesnt exist"),
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		os.Stdin = tc.req
		err := Init()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if err := tc.req.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func setInputSave(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
	}
	return file
}
