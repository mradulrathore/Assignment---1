package view

import (
	"log"
	"os"
	"testing"

	"github.com/mradulrathore/dependency-graph/view"
)

func TestInit(t *testing.T) {
	fileAllOptions := setInputSave("test_file.txt")
	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{
		{
			scenario: "dependency graph all options test",
			req:      fileAllOptions,
			err:      nil,
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		os.Stdin = tc.req
		err := view.Init()
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
