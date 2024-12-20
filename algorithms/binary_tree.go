package algorithms

import "github.com/dimitrijjedich/go-algorithms/helper"

func BinaryTreeSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	tree := helper.Tree{}

	for i := 0; i < len(arr); i++ {
		tree.Insert(arr[i])
	}
	if tree.Root != nil {
		result = tree.Root.Walk()
	}
	return result
}
