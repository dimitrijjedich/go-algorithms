package helper

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (node *Node) insert(newNode *Node) {
	if node.left == nil {
		node.left = newNode
	} else if newNode.data < node.data {
		node.left.insert(newNode)
	} else if node.right == nil {
		node.right = newNode
	} else {
		node.right.insert(newNode)
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
