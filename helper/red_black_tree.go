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

	tree.transform(newNode)
}

func (tree *RedBlackTree) transform(node *RedBlackNode) {
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
