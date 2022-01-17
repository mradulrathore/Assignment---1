package repository

import (
	"errors"
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
	// type userTest struct {
	// 	name         string
	// 	age          int
	// 	address      string
	// 	rollNo       int
	// 	coursesEnrol []string
	// }
	// file, err := os.OpenFile("user-data", os.O_RDWR|os.O_CREATE, 0755)
	// if err != nil {
	// 	t.Errorf("Scenario: %s \n got: %v, expected: %v", "save file", err, nil)
	// }

	// user := []userTest{{
	// 	name:         "Mradul",
	// 	age:          21,
	// 	address:      "Indore",
	// 	rollNo:       43,
	// 	coursesEnrol: []string{"A", "C", "D", "E", "F"},
	// }}

	// tests := []struct {
	// 	scenario string
	// 	req      []userTest
	// 	err      error
	// }{{
	// 	scenario: "save file",
	// 	req:      user,
	// 	err:      nil,
	// }}

	// for _, tc := range tests {
	// 	err := Save(file, tc.req)
	// 	if err != nil && tc.err == nil {
	// 		t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
	// 	} else if err == nil && tc.err != nil {
	// 		t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
	// 	}
	// }
}

func TestRetrieve(t *testing.T) {
	file1, err := os.OpenFile("user-data", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "save file", err, nil)
	}

	tests := []struct {
		scenario string
		file     *os.File
		err      error
	}{{
		scenario: "retrieve file (file exist)",
		file:     file1,
		err:      nil,
	}, {
		scenario: "retrieve file (file doesn't exist)",
		file:     nil,
		err:      errors.New("file doesn't exist"),
	}}

	for _, tc := range tests {
		_, err := RetrieveData(tc.file)
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}
