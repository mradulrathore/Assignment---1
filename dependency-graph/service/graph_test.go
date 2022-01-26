package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListNodeParent(t *testing.T) {
	graph := NewGraph()

	err := graph.AddNode(1, "A", nil)
	require.Nil(t, err)

	err = graph.AddNode(2, "B", nil)
	require.Nil(t, err)

	err = graph.AddEdge(1, 2)
	require.Nil(t, err)

	err = graph.AddNode(3, "C", nil)
	require.Nil(t, err)

	err = graph.AddNode(4, "D", nil)
	require.Nil(t, err)

	tests := []struct {
		scenario string
		nodeId   int
		checkId  int
		exist    bool
		err      error
	}{
		{
			scenario: "List node parents when edge between nodes doesn't exist",
			nodeId:   4,
			checkId:  3,
			exist:    false,
			err:      nil,
		}, {
			scenario: "List node parents when edge between nodes exist",
			nodeId:   2,
			checkId:  1,
			exist:    true,
			err:      nil,
		}, {
			scenario: "List node parents when node doesn't exist",
			nodeId:   21,
			checkId:  -1,
			exist:    false,
			err:      fmt.Errorf("node doesn't exist, id:%d", 21),
		},
	}

	for _, tc := range tests {
		nodes, err := graph.ListNodeParent(tc.nodeId)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		_, exist := nodes[tc.checkId]
		if exist != tc.exist {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, exist, tc.exist)
		}
	}
}

func TestListNodeChild(t *testing.T) {
	graph := NewGraph()

	err := graph.AddNode(1, "A", nil)
	require.Nil(t, err)

	err = graph.AddNode(2, "B", nil)
	require.Nil(t, err)

	err = graph.AddEdge(1, 2)
	require.Nil(t, err)

	err = graph.AddNode(3, "C", nil)
	require.Nil(t, err)

	err = graph.AddNode(4, "D", nil)
	require.Nil(t, err)

	tests := []struct {
		scenario string
		nodeId   int
		checkId  int
		exist    bool
		err      error
	}{
		{
			scenario: "List node parents when edge between nodes doesn't exist",
			nodeId:   4,
			checkId:  3,
			exist:    false,
			err:      nil,
		}, {
			scenario: "List node parents when edge between nodes exist",
			nodeId:   1,
			checkId:  2,
			exist:    true,
			err:      nil,
		}, {
			scenario: "List node parents when node doesn't exist",
			nodeId:   21,
			checkId:  -1,
			exist:    false,
			err:      fmt.Errorf("node doesn't exist, id:%d", 21),
		},
	}

	for _, tc := range tests {
		nodes, err := graph.ListNodeChild(tc.nodeId)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		_, exist := nodes[tc.checkId]
		if exist != tc.exist {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, exist, tc.exist)
		}
	}
}

func TestListAncestors(t *testing.T) {
	graph := NewGraph()

	err := graph.AddNode(1, "A", nil)
	require.Nil(t, err)

	err = graph.AddNode(2, "B", nil)
	require.Nil(t, err)

	err = graph.AddNode(3, "C", nil)
	require.Nil(t, err)

	err = graph.AddNode(4, "D", nil)
	require.Nil(t, err)

	err = graph.AddEdge(1, 2)
	require.Nil(t, err)

	err = graph.AddEdge(2, 3)
	require.Nil(t, err)

	err = graph.AddEdge(3, 4)
	require.Nil(t, err)

	tests := []struct {
		scenario  string
		nodeId    int
		ancestors []int
		err       error
	}{
		{
			scenario:  "List node ancestors when node exist",
			nodeId:    4,
			ancestors: []int{1, 2, 3},
			err:       nil,
		}, {
			scenario:  "List node ancestors when node doesn't exist",
			nodeId:    12,
			ancestors: []int{},
			err:       fmt.Errorf("node doesn't exist, id:%d", 12),
		}, {
			scenario:  "List node ancestors when node exist",
			nodeId:    1,
			ancestors: []int{},
			err:       nil,
		},
	}

	for _, tc := range tests {
		nodes, err := graph.ListAncestors(tc.nodeId)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if len(tc.ancestors) != len(nodes) {
			t.Errorf("Scenario: %s \n got no of ancestors: %v, expected no of ancestors: %v", tc.scenario, len(nodes), len(tc.ancestors))
		}

		for _, id := range tc.ancestors {
			_, exist := nodes[id]
			if exist != true {
				t.Errorf("Scenario: %s \n %d got: %v, expected: %v", tc.scenario, id, exist, true)
			}
		}
	}
}

func TestListDescendants(t *testing.T) {
	graph := NewGraph()

	err := graph.AddNode(1, "A", nil)
	require.Nil(t, err)

	err = graph.AddNode(2, "B", nil)
	require.Nil(t, err)

	err = graph.AddNode(3, "C", nil)
	require.Nil(t, err)

	err = graph.AddNode(4, "D", nil)
	require.Nil(t, err)

	err = graph.AddEdge(1, 2)
	require.Nil(t, err)

	err = graph.AddEdge(2, 3)
	require.Nil(t, err)

	err = graph.AddEdge(3, 4)
	require.Nil(t, err)

	tests := []struct {
		scenario    string
		nodeId      int
		descendants []int
		err         error
	}{
		{
			scenario:    "List node descendants when node exist",
			nodeId:      1,
			descendants: []int{2, 3, 4},
			err:         nil,
		}, {
			scenario:    "List node descendants when node doesn't exist",
			nodeId:      12,
			descendants: []int{},
			err:         fmt.Errorf("node doesn't exist, id:%d", 12),
		}, {
			scenario:    "List node descendants when node exist",
			nodeId:      4,
			descendants: []int{},
			err:         nil,
		},
	}

	for _, tc := range tests {
		nodes, err := graph.ListDescendants(tc.nodeId)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if len(tc.descendants) != len(nodes) {
			t.Errorf("Scenario: %s \n got no of descendants: %v, expected no of descendants: %v", tc.scenario, len(nodes), len(tc.descendants))
		}

		for _, id := range tc.descendants {
			_, exist := nodes[id]
			if exist != true {
				t.Errorf("Scenario: %s \n %d got: %v, expected: %v", tc.scenario, id, exist, true)
			}
		}
	}
}

func TestDeleteEdge(t *testing.T) {
	graph := NewGraph()

	err := graph.AddNode(1, "A", nil)
	require.Nil(t, err)

	err = graph.AddNode(2, "B", nil)
	require.Nil(t, err)

	err = graph.AddNode(3, "C", nil)
	require.Nil(t, err)

	err = graph.AddNode(4, "D", nil)
	require.Nil(t, err)

	err = graph.AddNode(5, "D", nil)
	require.Nil(t, err)

	err = graph.AddEdge(1, 2)
	require.Nil(t, err)

	err = graph.AddEdge(2, 3)
	require.Nil(t, err)

	err = graph.AddEdge(3, 4)
	require.Nil(t, err)

	err = graph.AddEdge(1, 5)
	require.Nil(t, err)

	tests := []struct {
		scenario string
		nodeId1  int
		nodeId2  int
		err      error
	}{
		{
			scenario: "delete edge when edge exist",
			nodeId1:  1,
			nodeId2:  2,
			err:      nil,
		}, {
			scenario: "delete edge when node doesn't exist",
			nodeId1:  12,
			nodeId2:  1,
			err:      fmt.Errorf("node doesn't exist, id:%d", 12),
		}, {
			scenario: "delete edge when node doesn't exist",
			nodeId1:  14,
			nodeId2:  1,
			err:      fmt.Errorf("node doesn't exist, id:%d", 14),
		}, {
			scenario: "delete edge when node doesn't exist",
			nodeId1:  1,
			nodeId2:  14,
			err:      fmt.Errorf("node doesn't exist, id:%d", 14),
		}, {
			scenario: "delete edge which doesn't exist",
			nodeId1:  4,
			nodeId2:  1,
			err:      fmt.Errorf("edge doesn't exist"),
		}, {
			scenario: "delete edge when node doesn't exist",
			nodeId1:  1,
			nodeId2:  3,
			err:      fmt.Errorf("node doesn't exist"),
		},
	}

	for _, tc := range tests {
		err := graph.DeleteEdge(tc.nodeId1, tc.nodeId2)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

	}
}

func TestDeleteNode(t *testing.T) {
	graph := NewGraph()

	err := graph.AddNode(1, "A", nil)
	require.Nil(t, err)

	err = graph.AddNode(2, "B", nil)
	require.Nil(t, err)

	err = graph.AddNode(3, "C", nil)
	require.Nil(t, err)

	err = graph.AddNode(4, "D", nil)
	require.Nil(t, err)

	err = graph.AddEdge(1, 2)
	require.Nil(t, err)

	err = graph.AddEdge(2, 3)
	require.Nil(t, err)

	err = graph.AddEdge(3, 4)
	require.Nil(t, err)

	tests := []struct {
		scenario     string
		nodeId       int
		parentNodeId int
		err          error
	}{
		{
			scenario:     "delete node when node exist",
			nodeId:       2,
			parentNodeId: 1,
			err:          nil,
		}, {
			scenario:     "delete node when node doesn't exist",
			nodeId:       12,
			parentNodeId: -1,
			err:          fmt.Errorf("node doesn't exist, id:%d", 12),
		}, {
			scenario:     "List node ancestors when node exist",
			nodeId:       4,
			parentNodeId: 3,
			err:          nil,
		},
	}

	for _, tc := range tests {
		err := graph.DeleteNode(tc.nodeId)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
		if tc.err == nil {
			nodes, err := graph.ListNodeChild(tc.parentNodeId)
			if err != nil {
				t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, nil)
			}
			_, exist := nodes[tc.nodeId]
			if exist != false {
				t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, exist, false)
			}
		}
	}
}

func TestAddEdge(t *testing.T) {
	graph := NewGraph()

	err := graph.AddNode(11, "A", nil)
	require.Nil(t, err)

	err = graph.AddNode(12, "B", nil)
	require.Nil(t, err)

	err = graph.AddNode(13, "C", nil)
	require.Nil(t, err)

	err = graph.AddNode(14, "D", nil)
	require.Nil(t, err)

	tests := []struct {
		scenario     string
		parentNodeId int
		childNodeId  int
		err          error
	}{{
		scenario:     "add edge when nodes exist",
		parentNodeId: 11,
		childNodeId:  12,
		err:          nil,
	}, {
		scenario:     "add edge when nodes do not exist",
		parentNodeId: 110,
		childNodeId:  2,
		err:          fmt.Errorf("node doesn't exist, id:%d", 12),
	}, {
		scenario:     "add edge when nodes do not exist",
		parentNodeId: 11,
		childNodeId:  120,
		err:          fmt.Errorf("node doesn't exist, id:%d", 12),
	}, {
		scenario:     "cyclic dependency",
		parentNodeId: 12,
		childNodeId:  11,
		err:          fmt.Errorf("cyclic dependency"),
	}, {
		scenario:     "dependency exists",
		parentNodeId: 11,
		childNodeId:  12,
		err:          fmt.Errorf("dependency exists"),
	}}

	for _, tc := range tests {
		err := graph.AddEdge(tc.parentNodeId, tc.childNodeId)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestAddNode(t *testing.T) {
	graph := NewGraph()

	metaData := make(map[string]string)
	metaData["color"] = "Red"
	metaData["vistied"] = "Not yet"

	tests := []struct {
		scenario string
		id       int
		name     string
		metaData map[string]string
		err      error
	}{{
		scenario: "new node",
		id:       1,
		name:     "A",
		metaData: metaData,
		err:      nil,
	}, {
		scenario: "new node when node exists",
		id:       1,
		name:     "B",
		metaData: metaData,
		err:      fmt.Errorf("node exists (id:%d)", 12),
	}}

	for _, tc := range tests {
		err := graph.AddNode(tc.id, tc.name, tc.metaData)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}
