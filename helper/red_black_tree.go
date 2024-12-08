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
	if parent.parent.left == parent {
		return parent.parent.right
	}
	return parent.parent.left
}

type RedBlackTree struct {
	Root *RedBlackNode
}

func (tree *RedBlackTree) Insert(value int) {
	newNode := &RedBlackNode{value: value, color: Red, parent: nil, left: nil, right: nil}
	if tree.Root == nil {
		tree.Root = newNode
		tree.Root.color = Black
		return
	}

	pointer := tree.Root
	for {
		if newNode.value < pointer.value && pointer.left != nil {
			pointer = pointer.left
		} else if newNode.value < pointer.value {
			pointer.left = newNode
			break
		} else if pointer.right != nil {
			pointer = pointer.right
		} else {
			pointer.right = newNode
			break
		}
	}
	newNode.parent = pointer
	if newNode.parent.color == Black {
		return
	}
	if newNode.parent.parent != nil && newNode.parent.color == Red && (newNode.uncle() == nil || newNode.uncle().color == Red) {
		newNode.parent.color = Black
		newNode.parent.parent.color = Red
		if newNode.uncle() != nil {
			newNode.uncle().color = Black
		}
	}
	grandparent := newNode.parent.parent
	parent := newNode.parent
	if (newNode.uncle() == nil || newNode.uncle().color == Black) && parent.color == Red {
		if grandparent == tree.Root {
			tree.Root = parent
			parent.right = grandparent
			parent.parent = nil
			grandparent.parent = parent
		}
		if (parent.left == newNode && grandparent.left == parent) || (parent.right == newNode || parent == grandparent.right) {
			parent.parent = grandparent.parent
			grandparent.parent = parent
			if parent.parent.right == grandparent {
				parent.parent.right = parent
				grandparent.right = parent.right
				parent.right = nil
			} else {
				parent.parent.left = parent
				grandparent.left = parent.left
				parent.left = nil
			}
		}
		parent.color = Black
		grandparent.color = Red
		return
	}
	if newNode.parent.color == Red && newNode.parent == tree.Root {
		newNode.parent.color = Black
		return
	}
}

func (node *RedBlackNode) Walk() []int {
	result := make([]int, 0)
	if node == nil {
		return result
	}
	result = append(node.left.Walk(), node.value)
	result = append(result, node.right.Walk()...)
	return result
}
