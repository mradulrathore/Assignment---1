package graph

import (
	"fmt"
	"log"

	"github.com/mradulrathore/dependency-graph/service/node"
)

type Graph struct {
	Nodes map[int]*node.Node
}

var g Graph

const (
	NodeNotExistErr = "node doesn't exist, id:%d"
	NodeExistMsg    = "node exists (id: %d)"
)

func GetParent(id int) (map[int]*node.Node, error) {
	n, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	return n.Parent, nil
}

func GetChild(id int) (map[int]*node.Node, error) {
	n, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	return n.Child, nil
}

func GetAncestors(id int) ([]int, error) {
	n, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	var ancestors []int
	ancestorsDFS(n, func(i int) {
		ancestors = append(ancestors, i)
	})

	return ancestors, nil
}

func ancestorsDFS(n *node.Node, visitCallback func(int)) {
	if n == nil {
		return
	}

	for id, ancestor := range n.Parent {
		visitCallback(id)
		ancestorsDFS(ancestor, visitCallback)
	}
}

func GetDescendants(id int) ([]int, error) {
	n, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	var descendants []int
	descendantsDFS(n, func(i int) {
		descendants = append(descendants, i)
	})

	return descendants, nil
}

func descendantsDFS(n *node.Node, visitCallback func(int)) {
	if n == nil {
		return
	}

	for id, descendants := range n.Child {
		visitCallback(id)
		descendantsDFS(descendants, visitCallback)
	}
}

func DeleteEdge(id1 int, id2 int) error {
	if g.Nodes == nil {
		err := fmt.Errorf(NodeNotExistErr, id1)
		return err
	}

	_, exist := g.Nodes[id1]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id1)
		return err
	}

	_, exist = g.Nodes[id2]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id2)
		return err
	}

	if g.Nodes[id1].Child == nil {
		err := fmt.Errorf("dependency doesn't exist")
		log.Println(err)
		return err
	}

	_, exist = g.Nodes[id1].Child[id2]
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
	if g.Nodes == nil {
		err := fmt.Errorf(NodeNotExistErr, id)
		return err
	}

	_, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return err
	}

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
		delete(node.Parent, id)
	}

	delete(g.Nodes, id)

	return nil
}

func AddEdge(id1, id2 int) error {
	if g.Nodes == nil {
		g.Nodes = make(map[int]*node.Node)
	}

	_, exist := g.Nodes[id1]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id1)
		return err
	}

	_, exist = g.Nodes[id2]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id2)
		return err
	}

	if g.Nodes[id1].Child == nil {
		g.Nodes[id1].Child = make(map[int]*node.Node)
	}
	g.Nodes[id1].Child[id2] = g.Nodes[id2]

	if g.Nodes[id2].Parent == nil {
		g.Nodes[id2].Parent = make(map[int]*node.Node)
	}
	g.Nodes[id2].Parent[id1] = g.Nodes[id1]

	return nil
}

func AddNode(id int, name string, metaData map[string]string) error {
	if g.Nodes == nil {
		g.Nodes = make(map[int]*node.Node)
	}

	_, exist := g.Nodes[id]
	if exist {
		err := fmt.Errorf(NodeExistMsg, id)
		return err
	}

	n := node.Node{
		Id:       id,
		Name:     name,
		MetaData: metaData,
	}
	g.Nodes[id] = &n

	return nil
}

func CheckIdExist(id int) (*node.Node, bool) {
	n, exist := g.Nodes[id]
	return n, exist
}
