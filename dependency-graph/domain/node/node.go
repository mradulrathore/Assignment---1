package node

import (
	"fmt"
)

type Node struct {
	Id       int
	Name     string
	MetaData map[string]string
	Parent   *Node
	Child    *Node
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}

	parentId := -1
	if n.Parent != nil {
		parentId = n.Parent.Id
	}

	childId := -1
	if n.Child != nil {
		parentId = n.Child.Id
	}

	return fmt.Sprintf("[%d, %s, %v, %d, %d", n.Id, n.Name, n.MetaData, parentId, childId)
}
