package application

import (
	"fmt"
	"sort"

	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/graph"
	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/node"
)

var g graph.Graph
var ids []int

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
	insertId(id)
	g.AddNode(&n)
}

func insertId(id int) {
	index := SearchId(id)
	insertIdAt(index, id)
}

func SearchId(rollNo int) (index int) {
	index = sort.Search(len(ids), func(i int) bool {
		return ids[i] >= rollNo
	})
	return
}

func insertIdAt(index int, id int) {
	if index == len(ids) {
		ids = append(ids, id)
		return
	}

	ids = append(ids[:index+1], ids[index:]...)
	ids[index] = id
}

func GetAllId() (i []int) {
	i = ids
	return
}
