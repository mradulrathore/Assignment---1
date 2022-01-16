package service

import (
	"errors"
	"testing"
)

// cmd go test -coverprofile=coverage.out

func TestLoadData(t *testing.T) {
	tests := []struct {
		scenario string
		err      error
	}{{
		scenario: "load data",
		err:      nil,
	}}

	for _, tc := range tests {
		err := LoadData()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}

}

func TestCheckDataExistence(t *testing.T) {
	tests := []struct {
		scenario string
		rollno   int
		exist    bool
	}{{
		scenario: "check data existence",
		rollno:   14,
		exist:    false,
	}}

	for _, tc := range tests {
		exist := CheckDataExistence(tc.rollno)
		if tc.exist != exist {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, exist, tc.exist)
		}
	}
}
func TestGetAll(t *testing.T) {
	tests := []struct {
		scenario string
		field    string
		order    int
		err      error
	}{{
		scenario: "fetch data sorted by name in ascending order",
		field:    "name",
		order:    1,
		err:      nil,
	}, {
		scenario: "fetch data sorted by age in ascending order",
		field:    "age",
		order:    1,
		err:      nil,
	}, {
		scenario: "fetch data sorted by name in descending order",
		field:    "name",
		order:    2,
		err:      nil,
	}, {
		scenario: "fetch data sorted by address in ascending order",
		field:    "address",
		order:    1,
		err:      nil,
	}, {
		scenario: "fetch data sorted by rollno in ascending order",
		field:    "rollno",
		order:    1,
		err:      nil,
	}}

	for _, tc := range tests {
		_, err := GetAll(tc.field, tc.order)
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestDeleteByRollNo(t *testing.T) {
	tests := []struct {
		scenario string
		rollno   int
		err      error
	}{{
		scenario: "delete by rollno",
		rollno:   15,
		err:      errors.New("rollno doesn't exist"),
	}}

	for _, tc := range tests {
		err := DeleteByRollNo(tc.rollno)
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}
