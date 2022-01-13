package graph

import (
	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/node"
)

type Graph struct {
	Nodes []*node.Node
	Edges map[*node.Node][]*node.Node
}

func (g *Graph) String() string {
	s := ""
	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + " -> "
		neigh := g.Edges[g.Nodes[i]]
		for j := 0; j < len(neigh); j++ {
			s += neigh[j].String() + " "
		}
		s += "\n"
	}
	return s
}
