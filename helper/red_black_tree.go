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
	if node.parent == nil || node.parent.parent == nil {
		return nil
	}
	if node.parent.parent.left == node.parent {
		return node.parent.parent.right
	}
	return node.parent.parent.left
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

	tree.transform(newNode)
}

func (tree *RedBlackTree) transform(node *RedBlackNode) {
	if node.parent.color == Black {
		return
	}
	if node.parent.parent != nil && node.parent.color == Red && (node.uncle() == nil || node.uncle().color == Red) {
		node.parent.color = Black
		node.parent.parent.color = Red
		if node.uncle() != nil {
			node.uncle().color = Black
		}
	}
	grandparent := node.parent.parent
	parent := node.parent
	if (node.uncle() == nil || node.uncle().color == Black) && parent.color == Red {
		if grandparent == tree.Root {
			tree.Root = parent
			parent.right = grandparent
			parent.parent = nil
			grandparent.parent = parent
		}
		if (parent.left == node && grandparent.left == parent) || (parent.right == node || parent == grandparent.right) {
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
	if node.parent.color == Red && node.parent == tree.Root {
		node.parent.color = Black
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
