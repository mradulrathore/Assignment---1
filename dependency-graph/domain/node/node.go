package node

import (
	"fmt"
)

type Node struct {
	Id       int
	Name     string
	MetaData map[string]string
	Parent   []*Node
	Child    []*Node
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}

	ancestorsId := []int{}
	if n.Parent != nil {
		for _, node := range n.Parent {
			ancestorsId = append(ancestorsId, node.Id)
		}
	}

	descendantsId := []int{}
	if n.Child != nil {
		for _, node := range n.Child {
			descendantsId = append(descendantsId, node.Id)
		}
	}

	return fmt.Sprintf("[%d, %s, %v, %v, %v", n.Id, n.Name, n.MetaData, ancestorsId, descendantsId)
}
