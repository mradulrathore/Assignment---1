package node

import (
	"fmt"
)

type Node struct {
	Id       int
	Name     string
	MetaData map[string]string
}

func (n *Node) String() string {
	return fmt.Sprintf("[%d, %s, %v", n.Id, n.Name, n.MetaData)
}
