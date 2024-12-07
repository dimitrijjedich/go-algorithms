package helper

type Color bool

const (
	Red   Color = true
	Black Color = false
)

type RedBlackNode struct {
	value  int
	color  Color
	parent *RedBlackNode
	left   *RedBlackNode
	right  *RedBlackNode
}

func (node *RedBlackNode) uncle() *RedBlackNode {
	parent := node.parent
	if parent.left == node {
		return parent.right
	}
	return parent.left
}

type RedBlackTree struct {
	root *RedBlackNode
}

func (tree *RedBlackTree) Insert(value int) {
	newNode := &RedBlackNode{value: value, color: Red, parent: nil, left: nil, right: nil}
	if tree.root == nil {
		tree.root = newNode
		tree.root.color = Black
		return
	}
	pointer := tree.root
	for {
		if newNode.value < pointer.value && pointer.left != nil {
			pointer = pointer.left
		} else if newNode.value < pointer.value {
			pointer.left = newNode
			break
		} else if pointer.right != nil {
			pointer = pointer.right
		} else {
			pointer.right.left = newNode
			break
		}
	}
	newNode.parent = pointer
	if newNode.parent.color == Black {
		return
	}
}
