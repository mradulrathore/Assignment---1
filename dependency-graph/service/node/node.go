package node

const (
	NodeNotExistErr = "node doesn't exist, id:%d"
)

type Node interface {
	Create(id int, name string, metaData map[string]string) *node
	GetParent() (map[int]Node, error)
	GetChild() (map[int]Node, error)
}

type node struct {
	id       int
	name     string
	metaData map[string]string
	parent   map[int]Node
	child    map[int]Node
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

func (n *node) GetParent() (map[int]Node, error) {
	return n.parent, nil
}

func (n *node) GetChild() (map[int]Node, error) {
	return n.child, nil
}
