package node

type Node struct {
	Id       int
	Name     string
	MetaData map[string]string
	Parent   map[int]*Node
	Child    map[int]*Node
}
