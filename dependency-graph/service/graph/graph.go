package graph

import (
	"fmt"
	"log"

	"github.com/mradulrathore/dependency-graph/service/node"
)

const (
	NodeNotExistErr = "node doesn't exist, id:%d"
	NodeExistErr    = "node exists (id: %d)"
	DuplicateIdErr  = "duplicate id:%d"
)

type GraphI interface {
	GetNodeParent(id int) (map[int]node.Node, error)
	GetNodeChild(id int) (map[int]node.Node, error)
	GetAncestors(int) ([]int, error)
	GetDescendants(int) ([]int, error)
	DeleteEdge(int, int) error
	DeleteNode(int) error
	AddEdge(int, int) error
	AddNode(int, string, map[int]string)
	CheckIdExist(int) (*node.Node, bool)
}

type graph struct {
	node  node.Node
	nodes map[int]node.Node
}

func NewGraph(node node.Node) *graph {
	return &graph{node: node}
}

func (g *graph) GetNodeParent(id int) (map[int]node.Node, error) {
	n, exist := g.nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}
	ancestors, err := n.GetParent()
	if err != nil {
		return nil, err
	}
	return ancestors, nil
}

func (g *graph) GetNodeChild(id int) (map[int]node.Node, error) {
	n, exist := g.nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}
	descendants, err := n.GetChild()
	if err != nil {
		return nil, err
	}
	return descendants, nil
}

func (g *graph) GetAncestors(id int) ([]int, error) {
	n, exist := g.nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	var ancestors []int
	ancestorsDFS(n, func(i int) { ancestors = append(ancestors, i) })

	return ancestors, nil
}

func ancestorsDFS(n node.Node, visitCallback func(int)) error {
	if n == nil {
		return nil
	}

	ancestors, err := n.GetParent()
	if err != nil {
		return err
	}

	for id, ancestor := range ancestors {
		visitCallback(id)
		ancestorsDFS(ancestor, visitCallback)
	}
	return nil
}

func (g *graph) GetDescendants(id int) ([]int, error) {
	n, exist := g.nodes[id]
	if !exist {
		err := fmt.Errorf(NodeNotExistErr, id)
		return nil, err
	}

	var descendants []int
	descendantsDFS(n, func(i int) { descendants = append(descendants, i) })

	return descendants, nil
}

func descendantsDFS(n node.Node, visitCallback func(int)) error {
	if n == nil {
		return nil
	}

	descendants, err := n.GetChild()
	if err != nil {
		return err
	}

	for id, descendant := range descendants {
		visitCallback(id)
		descendantsDFS(descendant, visitCallback)
	}

	return nil
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

	descendants, err := g.nodes[id1].GetChild()
	if err != nil {
		return err
	}

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

	ancestors, err := g.nodes[id2].GetParent()
	if err != nil {
		return err
	}

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

	parent, err := g.nodes[id].GetParent()
	if err != nil {
		return err
	}
	for _, node := range parent {
		descendants, err := node.GetChild()
		if err != nil {
			return err
		}
		delete(descendants, id)
	}

	child, err := g.nodes[id].GetChild()
	if err != nil {
		return err
	}
	for _, node := range child {
		ancestors, err := node.GetParent()
		if err != nil {
			return err
		}
		delete(ancestors, id)
	}

	delete(g.nodes, id)

	return nil
}

func (g *graph) AddEdge(id1, id2 int) error {
	if g.nodes == nil {
		g.nodes = make(map[int]node.Node)
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
	for _, a := range ancestors {
		if a == id2 {
			err := fmt.Errorf("cyclic dependency")
			log.Println(err)
			return err
		}
	}

	descendants, err := g.nodes[id1].GetChild()
	if err != nil {
		return err
	}
	if descendants == nil {
		descendants = make(map[int]node.Node)
	}
	descendants[id2] = g.nodes[id2]

	ancestorsNodes, err := g.nodes[id2].GetParent()
	if err != nil {
		return err
	}
	if ancestorsNodes == nil {
		ancestorsNodes = make(map[int]node.Node)
	}
	ancestorsNodes[id1] = g.nodes[id1]

	return nil
}

func (g *graph) AddNode(id int, name string, metaData map[string]string) error {
	if g.nodes == nil {
		g.nodes = make(map[int]node.Node)
	}

	_, err := g.CheckIdExist(id)
	if err != nil {
		return err
	}

	g.nodes[id] = g.nodes[id].Create(id, name, metaData)

	return nil
}

func (g *graph) CheckIdExist(id int) (node.Node, error) {
	node, exist := g.nodes[id]
	var err error
	if exist {
		err = fmt.Errorf(NodeExistErr, id)
	}
	return node, err
}
