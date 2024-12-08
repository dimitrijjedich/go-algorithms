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

	// starting point for the insert is the root of the tree
	pointer := tree.Root
	for {
		// if the value of the new node is smaller than the one of the current -> left section of the tree
		if newNode.value < pointer.value {
			// if the is no left child (nil)
			if pointer.left == nil {
				// make the new node the left child of the current node
				pointer.left = newNode
				// end the insertion process
				break
			} else {
				// if the left child exists, switch the pointer to it
				pointer = pointer.left
			}
		} else {
			// the value of the new node is greater or equal to the current -> right section of the tree
			// if the is no right child (nil)
			if pointer.right == nil {
				// make the new node the right child of the current node
				pointer.right = newNode
				// end the insertion process
				break
			} else {
				// if the left child exists, switch the pointer to it
				pointer = pointer.right
			}
		}
	}
	// the new node is now child of pointer and needs to have it as parent
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
