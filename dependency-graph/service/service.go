package service

import (
	"fmt"
	"log"

	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/graph"
	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/node"
)

var g graph.Graph

var NodeNotExistMsg = "node doesn't exist"

func GetParent(id int) (map[int]*node.Node, error) {
	n, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistMsg)
		return nil, err
	}
	return n.Parent, nil
}

func GetChild(id int) (map[int]*node.Node, error) {
	n, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistMsg)
		return nil, err
	}
	return n.Child, nil
}

func GetAncestors(id int) ([]int, error) {
	n, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistMsg)
		return nil, err
	}

	var ancestors []int
	cb := func(i int) {
		ancestors = append(ancestors, i)
	}
	ancestorsDFS(n, cb)

	return ancestors, nil
}

func ancestorsDFS(n *node.Node, visitCb func(int)) {
	visited := map[int]bool{}

	if n == nil {
		return
	}

	for id, ancestor := range n.Parent {
		if visited[id] {
			continue
		}
		visited[id] = true
		visitCb(id)
		ancestorsDFS(ancestor, visitCb)
	}
}

func GetDescendants(id int) ([]int, error) {
	n, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistMsg)
		return nil, err
	}

	var descendants []int
	cb := func(i int) {
		descendants = append(descendants, i)
	}
	descendantsDFS(n, cb)

	return descendants, nil
}

func descendantsDFS(n *node.Node, visitCb func(int)) {
	visited := map[int]bool{}

	if n == nil {
		return
	}

	for id, descendants := range n.Child {
		if visited[id] {
			continue
		}
		visited[id] = true
		visitCb(id)
		descendantsDFS(descendants, visitCb)
	}
}

func DeleteEdge(id1 int, id2 int) error {
	_, exist := g.Nodes[id1].Child[id2]
	if !exist {
		err := fmt.Errorf("dependency doesn't exist")
		log.Println(err)
		return err
	}

	delete(g.Nodes[id1].Child, id2)
	delete(g.Nodes[id2].Parent, id1)

	return nil
}

func DeleteNode(id int) error {
	_, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistMsg)
		return err
	}
	delete(g.Nodes, id)

	parent, err := GetParent(id)
	if err != nil {
		return err
	}
	for _, node := range parent {
		delete(node.Child, id)
	}

	child, err := GetChild(id)
	if err != nil {
		return err
	}
	for _, node := range child {
		delete(node.Child, id)
	}

	return nil
}

func AddEdge(id1, id2 int) {
	if g.Nodes[id1].Child == nil {
		g.Nodes[id1].Child = make(map[int]*node.Node)
	}
	g.Nodes[id1].Child[id2] = g.Nodes[id2]

	if g.Nodes[id2].Parent == nil {
		g.Nodes[id2].Parent = make(map[int]*node.Node)
	}
	g.Nodes[id2].Parent[id1] = g.Nodes[id1]
}

func AddNode(id int, name string, metaData map[string]string) {
	if g.Nodes == nil {
		g.Nodes = make(map[int]*node.Node)
	}

	n := node.Node{
		Id:       id,
		Name:     name,
		MetaData: metaData,
	}
	g.Nodes[id] = &n
}

func CheckIdExist(id int) (*node.Node, bool) {
	n, exist := g.Nodes[id]
	return n, exist
}
