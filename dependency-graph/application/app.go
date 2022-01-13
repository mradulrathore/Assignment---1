package application

import (
	"fmt"

	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/graph"
	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/node"
)

var g graph.Graph

func Init() {
	n1 := node.Node{
		Id:   1,
		Name: "A",
	}
	n2 := node.Node{
		Id:   2,
		Name: "B",
	}
	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddEdge(&n1, &n2)

	fmt.Println(g.String())

}

func AddNode(id int, name string, metaData map[string]string) {
	n := node.Node{
		Id:       id,
		Name:     name,
		MetaData: metaData,
	}
	g.AddNode(&n)
}
