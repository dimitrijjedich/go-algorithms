package helper

type Tree struct {
	root *Node
}

func (tree *Tree) add(data int) *Tree {
	newNode := &Node{data: data, left: nil, right: nil}
	if tree.root != nil {
		tree.root.insert(newNode)
	} else {
		tree.root = newNode
	}
	return tree
}
