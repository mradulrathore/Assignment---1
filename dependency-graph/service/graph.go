package service

import (
	"fmt"
	"log"
)

const (
	NodeNotExistErr = "node doesn't exist, id:%d"
	NodeExistErr    = "node exists (id: %d)"
	DuplicateIdErr  = "duplicate id:%d"
)

type GraphI interface {
	GetAncestors(int) ([]int, error)
	GetDescendants(int) ([]int, error)
	DeleteEdge(int, int) error
	DeleteNode(int) error
	AddEdge(int, int) error
	AddNode(int, string, map[int]string)
	CheckIdExist(int) (*node, bool)
}

type Graph struct {
	Nodes map[int]*node
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) GetParent(id int) (map[int]*node, error) {
	n, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	return n.parent, nil
}

func (g *Graph) GetChild(id int) (map[int]*node, error) {
	n, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	return n.child, nil
}

func (g *Graph) GetAncestors(id int) ([]int, error) {
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

func ancestorsDFS(n *node, visitCallback func(int)) {
	if n == nil {
		return
	}

	for id, ancestor := range n.parent {
		visitCallback(id)
		ancestorsDFS(ancestor, visitCallback)
	}
}

func (g *Graph) GetDescendants(id int) ([]int, error) {
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

func descendantsDFS(n *node, visitCallback func(int)) {
	if n == nil {
		return
	}

	for id, descendants := range n.child {
		visitCallback(id)
		descendantsDFS(descendants, visitCallback)
	}
}

func (g *Graph) DeleteEdge(id1 int, id2 int) error {
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

	if g.Nodes[id1].child == nil {
		err := fmt.Errorf("dependency doesn't exist")
		log.Println(err)
		return err
	}

	_, exist = g.Nodes[id1].child[id2]
	if !exist {
		err := fmt.Errorf("dependency doesn't exist")
		log.Println(err)
		return err
	}

	delete(g.Nodes[id1].child, id2)
	delete(g.Nodes[id2].parent, id1)

	return nil
}

func (g *Graph) DeleteNode(id int) error {
	if g.Nodes == nil {
		err := fmt.Errorf(NodeNotExistErr, id)
		return err
	}

	_, exist := g.Nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return err
	}

	parent, err := g.GetParent(id)
	if err != nil {
		return err
	}
	for _, node := range parent {
		delete(node.child, id)
	}

	child, err := g.GetChild(id)
	if err != nil {
		return err
	}
	for _, node := range child {
		delete(node.parent, id)
	}

	delete(g.Nodes, id)

	return nil
}

func (g *Graph) AddEdge(id1, id2 int) error {
	if g.Nodes == nil {
		g.Nodes = make(map[int]*node)
	}

	_, exist := g.Nodes[id1]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id1)
		return err
	}

	_, exist = g.Nodes[id2]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id1)
		return err
	}

	ancestors, err := g.GetAncestors(id1)
	if err != nil {
		return err
	}
	for _, a := range ancestors {
		if a == id2 {
			err := fmt.Errorf("cyclic dependency")
			log.Println(err)
			return err
		}
	}

	if g.Nodes[id1].child == nil {
		g.Nodes[id1].child = make(map[int]*node)
	}
	g.Nodes[id1].child[id2] = g.Nodes[id2]

	if g.Nodes[id2].parent == nil {
		g.Nodes[id2].parent = make(map[int]*node)
	}
	g.Nodes[id2].parent[id1] = g.Nodes[id1]

	return nil
}

func (g *Graph) AddNode(id int, name string, metaData map[string]string) error {
	if g.Nodes == nil {
		g.Nodes = make(map[int]*node)
	}

	_, err := g.CheckIdExist(id)
	if err != nil {
		return err
	}

	node := NewNode()
	g.Nodes[id] = node.Create(id, name, metaData)

	return nil
}

func (g *Graph) CheckIdExist(id int) (*node, error) {
	node, exist := g.Nodes[id]
	var err error
	if exist {
		err = fmt.Errorf(NodeExistErr, id)
	}
	return node, err
}
