package service

type Node interface {
	GetParent() (map[int]*Node, error)
	GetChild() (map[int]*Node, error)
}

type node struct {
	id       int
	name     string
	metaData map[string]string
	parent   map[int]*node
	child    map[int]*node
}

func NewNode() *node {
	return &node{}
}

func (n *node) Create(id int, name string, metaData map[string]string) *node {
	return &node{
		id:       id,
		name:     name,
		metaData: metaData,
	}
}
