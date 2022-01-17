package graph

import (
	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/node"
)

type Graph struct {
	Nodes map[int]*node.Node
}

func (g *Graph) String() string {
	s := ""
	for _, node := range g.Nodes {
		s += node.String() + " -> "
		parent := node.Parent
		for _, parentNode := range parent {
			s += parentNode.String() + " "
		}
		child := node.Child
		for _, childNode := range child {
			s += childNode.String() + " "
		}
		s += "\n"
	}
	return s
}
