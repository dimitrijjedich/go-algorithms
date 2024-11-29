package algorithms

func RadixSort(arr []int) []int {
	return arr
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	var partitions [10][]int
	for i := range partitions {
		partitions[i] = make([]int, 0)
	}
}
