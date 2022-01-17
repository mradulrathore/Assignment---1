package graph

import (
	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/node"
)

type Graph struct {
	Nodes map[int]*node.Node
}
