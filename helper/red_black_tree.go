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
	// create a new RedBlackNode with the provided value
	// initiate the node as Red with no parents and no children
	newNode := &RedBlackNode{value: value, color: Red, parent: nil, left: nil, right: nil}

	// if the root of the tree is empty (nil)
	if tree.Root == nil {
		// make the new node the new root of the tree, ...
		tree.Root = newNode
		// ...recolor the node to Black (root is always black)...
		tree.Root.color = Black
		// ...and finish the process
		return
	}

	pointer := tree.Root
	for {
		if newNode.value < pointer.value {
			if pointer.left == nil {
				pointer.left = newNode
				break
			} else {
				pointer = pointer.left
			}
		} else {
			if pointer.right == nil {
				pointer.right = newNode
				break
			} else {
				pointer = pointer.right
			}
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
