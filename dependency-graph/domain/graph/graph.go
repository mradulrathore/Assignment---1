package graph

import (
	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/node"
)

type Graph struct {
	nodes []*node.Node
	edges map[*node.Node][]*node.Node
}

func (g *Graph) AddNode(n *node.Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *node.Node) {
	if g.edges == nil {
		g.edges = make(map[*node.Node][]*node.Node)
	}
	g.edges[n1] = append(g.edges[n1], n2)

}

func (g *Graph) String() string {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		neigh := g.edges[g.nodes[i]]
		for j := 0; j < len(neigh); j++ {
			s += neigh[j].String() + " "
		}
		s += "\n"
	}
	return s
}
