package service

// import (
// 	"errors"
// 	"testing"
// )

// func TestGetParent(t *testing.T) {
// 	tests := []struct {
// 		scenario string
// 		req      int
// 		err      error
// 	}{
// 		{
// 			scenario: "get parent",
// 			req:      1,
// 			err:      errors.New("node doesn't exist"),
// 		},
// 	}

// 	for _, tc := range tests {
// 		_, err := GetParent(tc.req)
// 		if err != nil && tc.err == nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		} else if err == nil && tc.err != nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		}
// 	}
// }

// func TestGetChild(t *testing.T) {
// 	tests := []struct {
// 		scenario string
// 		req      int
// 		err      error
// 	}{
// 		{
// 			scenario: "get ancestors",
// 			req:      1,
// 			err:      errors.New("node doesn't exist"),
// 		},
// 	}

// 	for _, tc := range tests {
// 		_, err := GetChild(tc.req)
// 		if err != nil && tc.err == nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		} else if err == nil && tc.err != nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		}
// 	}
// }

// func TestGetAncestors(t *testing.T) {

// 	tests := []struct {
// 		scenario string
// 		req      int
// 		err      error
// 	}{
// 		{
// 			scenario: "get ancestors",
// 			req:      1,
// 			err:      errors.New("node doesn't exist"),
// 		},
// 	}

// 	for _, tc := range tests {
// 		_, err := GetAncestors(tc.req)
// 		if err != nil && tc.err == nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		} else if err == nil && tc.err != nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		}
// 	}
// }

// func TestGetDescendants(t *testing.T) {
// 	tests := []struct {
// 		scenario string
// 		req      int
// 		err      error
// 	}{
// 		{
// 			scenario: "get descendants",
// 			req:      1,
// 			err:      errors.New("node doesn't exist"),
// 		},
// 	}

// 	for _, tc := range tests {
// 		_, err := GetDescendants(tc.req)
// 		if err != nil && tc.err == nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		} else if err == nil && tc.err != nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		}
// 	}
// }

// func TestDeleteEdge(t *testing.T) {
// 	tests := []struct {
// 		scenario string
// 		req      []int
// 		err      error
// 	}{
// 		{
// 			scenario: "delete edge",
// 			req:      []int{1, 2},
// 			err:      errors.New("node doesn't exist"),
// 		},
// 	}

// 	for _, tc := range tests {
// 		err := DeleteEdge(tc.req[0], tc.req[1])
// 		if err != nil && tc.err == nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		} else if err == nil && tc.err != nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		}
// 	}
// }

// func TestDeleteNode(t *testing.T) {
// 	tests := []struct {
// 		scenario string
// 		req      int
// 		err      error
// 	}{
// 		{
// 			scenario: "delete node",
// 			req:      1,
// 			err:      errors.New("node doesn't exist"),
// 		},
// 	}

// 	for _, tc := range tests {
// 		err := DeleteNode(tc.req)
// 		if err != nil && tc.err == nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		} else if err == nil && tc.err != nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		}
// 	}
// }

// func TestAddEdge(t *testing.T) {
// 	tests := []struct {
// 		scenario string
// 		req      []int
// 		err      error
// 	}{
// 		{
// 			scenario: "add edge",
// 			req:      []int{1, 2},
// 			err:      errors.New("node doesn't exist"),
// 		},
// 	}

// 	for _, tc := range tests {
// 		err := AddEdge(tc.req[0], tc.req[1])
// 		if err != nil && tc.err == nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		} else if err == nil && tc.err != nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		}
// 	}
// }

// func TestInit(t *testing.T) {
// 	type testStruct struct {
// 		id       int
// 		name     string
// 		metadata map[string]string
// 	}

// 	tests := []struct {
// 		scenario string
// 		req      testStruct
// 		err      error
// 	}{
// 		{
// 			scenario: "add node",
// 			req: testStruct{
// 				id:       11,
// 				name:     "A",
// 				metadata: nil,
// 			},
// 			err: nil,
// 		},
// 	}

// 	for _, tc := range tests {
// 		err := AddNode(tc.req.id, tc.req.name, nil)
// 		if err != nil && tc.err == nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		} else if err == nil && tc.err != nil {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
// 		}
// 	}
// }

// func TestCheckIdExist(t *testing.T) {
// 	tests := []struct {
// 		scenario string
// 		req      int
// 		res      bool
// 	}{
// 		{
// 			scenario: "check id existence",
// 			req:      1,
// 			res:      false,
// 		},
// 	}

// 	for _, tc := range tests {
// 		_, exist := CheckIdExist(tc.req)
// 		if exist != tc.res {
// 			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, exist, tc.res)
// 		}
// 	}
// }
