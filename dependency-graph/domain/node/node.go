package node

import (
	"fmt"
)

type Node struct {
	Id       int
	Name     string
	MetaData map[string]string
	Parent   map[int]*Node
	Child    map[int]*Node
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}

	ancestorsId := []int{}
	if n.Parent != nil {
		for id, _ := range n.Parent {
			ancestorsId = append(ancestorsId, id)
		}
	}

	descendantsId := []int{}
	if n.Child != nil {
		for id, _ := range n.Child {
			descendantsId = append(descendantsId, id)
		}
	}

	return fmt.Sprintf("[%d, %s, %v, %v, %v", n.Id, n.Name, n.MetaData, ancestorsId, descendantsId)
}
