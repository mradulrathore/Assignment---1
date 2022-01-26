package graph

import (
	"fmt"
	"log"
)

const (
	DependencyExistErr = "dependency exists"
	NodeNotExistErr    = "node doesn't exist, id:%d"
	NodeExistErr       = "node exists (id: %d)"
	DuplicateIdErr     = "duplicate id:%d"
)

type Graph interface {
	GetNodeParent(id int) (map[int]*node, error)
	GetNodeChild(id int) (map[int]*node, error)
	GetAncestors(int) (map[int]*node, error)
	GetDescendants(int) (map[int]*node, error)
	DeleteEdge(int, int) error
	DeleteNode(int) error
	AddEdge(int, int) error
	AddNode(int, string, map[string]string) error
	GetNodesID(map[int]*node) []int
}

type graph struct {
	nodes map[int]*node
}

type node struct {
	id       int
	name     string
	metaData map[string]string
	parent   map[int]*node
	child    map[int]*node
}

func NewGraph() *graph {
	return &graph{}
}

func (g *graph) GetNodeParent(id int) (map[int]*node, error) {
	n, exist := g.nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}
	ancestors := make(map[int]*node)
	for id, node := range n.parent {
		ancestors[id] = node
	}

	return ancestors, nil
}

func (g *graph) GetNodeChild(id int) (map[int]*node, error) {
	n, exist := g.nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}
	descendants := make(map[int]*node)
	for id, node := range n.child {
		descendants[id] = node
	}

	return descendants, nil
}

func (g *graph) GetAncestors(id int) (map[int]*node, error) {
	if g.nodes == nil {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}
	n, exist := g.nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	ancestors := make(map[int]*node)
	ancestorsDFS(n, func(i *node) { ancestors[i.id] = i })

	return ancestors, nil
}

func ancestorsDFS(n *node, visitCallback func(*node)) {
	if n == nil {
		return
	}

	ancestors := n.parent
	for _, ancestor := range ancestors {
		visitCallback(ancestor)
		ancestorsDFS(ancestor, visitCallback)
	}
}

func (g *graph) GetDescendants(id int) (map[int]*node, error) {
	if g.nodes == nil {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	n, exist := g.nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	descendants := make(map[int]*node)
	descendantsDFS(n, func(i *node) { descendants[i.id] = i })

	return descendants, nil
}

func descendantsDFS(n *node, visitCallback func(*node)) {
	if n == nil {
		return
	}

	descendants := n.child
	for _, descendant := range descendants {
		visitCallback(descendant)
		descendantsDFS(descendant, visitCallback)
	}
}

func (g *graph) DeleteEdge(id1 int, id2 int) error {
	if g.nodes == nil {
		err := fmt.Errorf(NodeNotExistErr, id1)
		return err
	}

	_, exist := g.nodes[id1]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id1)
		return err
	}

	_, exist = g.nodes[id2]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id2)
		return err
	}

	descendants := g.nodes[id1].child
	if descendants == nil {
		err := fmt.Errorf("dependency doesn't exist")
		log.Println(err)
		return err
	}

	_, exist = descendants[id2]
	if !exist {
		err := fmt.Errorf("dependency doesn't exist")
		log.Println(err)
		return err
	}

	ancestors := g.nodes[id2].parent

	delete(descendants, id2)
	delete(ancestors, id1)

	return nil
}

func (g *graph) DeleteNode(id int) error {
	if g.nodes == nil {
		err := fmt.Errorf(NodeNotExistErr, id)
		return err
	}

	_, exist := g.nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return err
	}

	parent := g.nodes[id].parent

	for _, node := range parent {
		descendants := node.child
		delete(descendants, id)
	}

	child := g.nodes[id].child

	for _, node := range child {
		ancestors := node.parent
		delete(ancestors, id)
	}

	delete(g.nodes, id)

	return nil
}

func (g *graph) AddEdge(id1, id2 int) error {
	if g.nodes == nil {
		err := fmt.Errorf(NodeNotExistErr, id1)
		return err
	}

	_, exist := g.nodes[id1]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id1)
		return err
	}

	_, exist = g.nodes[id2]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id1)
		return err
	}

	ancestors, err := g.GetAncestors(id1)
	if err != nil {
		return err
	}

	for id := range ancestors {
		if id == id2 {
			err := fmt.Errorf("cyclic dependency")
			log.Println(err)
			return err
		}
	}

	if g.nodes[id1].child == nil {
		g.nodes[id1].child = make(map[int]*node)
	}

	_, exist = g.nodes[id1].child[id2]
	if exist {
		err := fmt.Errorf(DependencyExistErr)
		return err
	}

	g.nodes[id1].child[id2] = g.nodes[id2]

	if g.nodes[id2].parent == nil {
		g.nodes[id2].parent = make(map[int]*node)
	}
	g.nodes[id2].parent[id1] = g.nodes[id1]

	return nil
}

func (g *graph) AddNode(id int, name string, metaData map[string]string) error {
	if g.nodes == nil {
		g.nodes = make(map[int]*node)
	}

	_, exist := g.nodes[id]
	if exist {
		err := fmt.Errorf(NodeExistErr, id)
		log.Println(err)
		return err
	}

	node := node{id: id, name: name, metaData: metaData}
	g.nodes[id] = &node

	return nil
}

func (g *graph) GetNodesID(nodes map[int]*node) []int {
	var ids []int
	for id := range nodes {
		ids = append(ids, id)
	}
	return ids
}
