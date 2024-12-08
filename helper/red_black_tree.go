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
	// the node has no parent -> there will be no uncle
	// the node has no grandparent -> there will be no uncle
	if node.parent == nil || node.parent.parent == nil {
		return nil
	}
	// the grandparent left child is the parent of the node
	if node.parent.parent.left == node.parent {
		// the uncle is the right child of the grandparent
		return node.parent.parent.right
	}
	// the grandparent left child is not the parent of the node therefore it must be the uncle
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

	// fix now potentially broken tree
	tree.transform(newNode)
}

func (tree *RedBlackTree) transform(node *RedBlackNode) {
	// if the node is also the root of the tree
	if node == tree.Root {
		// change the color to black and exit
		node.color = Black
		return
	}
	// since the node is not the root of the tree, it has at least a parent
	// if the node has a black parent...
	if node.parent.color == Black {
		// do nothing and exit
		return
	} else {
		// the parent is red and so is the uncle ()
		if node.uncle() != nil && node.uncle().color == Red {
			// change color of parent and uncle to black
			node.parent.color = Black
			node.uncle().color = Black
			// and grandparent to red
			node.parent.parent.color = Red
			// and repeat the process for the grandparent
			tree.transform(node.parent.parent)
		} else {
			// the uncle is black (OR NIl)
			if node == node.parent.left && node.parent == node.parent.parent.left {
				// node and parent are left children
			} else if node == node.parent.right && node.parent == node.parent.parent.left {
				// node is right child while parent is left child
			} else if node == node.parent.right && node.parent == node.parent.parent.left {
				// node and parent are right children
			} else {
				// node is left child while parent is right child
			}
		}
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
