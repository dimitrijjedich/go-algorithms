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
	if newNode.parent.parent != nil && newNode.parent.color == Red && (newNode.uncle() == nil || newNode.uncle().color == Red) {
		newNode.parent.color = Black
		newNode.uncle().color = Black
		newNode.parent.parent.color = Red
	}
	grandparent := newNode.parent.parent
	parent := newNode.parent
	if (newNode.uncle() == nil || newNode.uncle().color == Black) && parent.color == Red {
		if grandparent == tree.root {
			tree.root = parent
			parent.right = grandparent
			parent.parent = nil
			grandparent.parent = parent
		}
		if (parent.left == newNode && grandparent.left == parent) || (parent.right == newNode || parent == grandparent.right) {
			if grandparent.parent.right == grandparent {
				parent.parent = grandparent.parent
				parent.parent.right = parent
				grandparent.parent = parent
				grandparent.right = parent.right
				parent.right = nil
			} else {
				parent.parent = grandparent.parent
				parent.parent.left = parent
				grandparent.parent = parent
				grandparent.left = parent.left
				parent.left = nil
			}
		}
		parent.color = Black
		grandparent.color = Red
	}
}
