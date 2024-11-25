package helper

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (node *Node) insert(newNode *Node) {
	if node.left != nil && newNode.data < node.data {
		node.left.insert(newNode)
	} else if newNode.data < node.data {
		node.left = newNode
	} else if node.left != nil && newNode.data > node.data {
		node.right.insert(newNode)
	} else {
		node.right = newNode
	}
}
