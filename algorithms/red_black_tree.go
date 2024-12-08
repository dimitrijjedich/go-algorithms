package algorithms

import "github.com/dimitrijjedich/go-algorithms/helper"

func RedBlackTreeSort(arr []int) []int {
	result := make([]int, len(arr))
	tree := helper.RedBlackTree{}
	for _, v := range arr {
		tree.Insert(v)
	}
	if tree.Root != nil {
		result = tree.Root.Walk()
	}
	return result
}
