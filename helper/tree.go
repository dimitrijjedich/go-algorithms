package helper

type Tree struct {
	Root *Node
}

func (tree *Tree) Insert(data int) *Tree {
	newNode := &Node{data: data, left: nil, right: nil}
	if tree.Root != nil {
		tree.Root.insert(newNode)
	} else {
		tree.Root = newNode
	}
	return tree
}
