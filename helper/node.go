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

func (node *Node) Walk() []int {
	result := make([]int, 0)
	for node := node.left; node.left != nil; node = node.left {
	}
	result = append(result, node.data)
	if node.right != nil {
		result = append(result, node.right.Walk()...)
	}
	return result
}
