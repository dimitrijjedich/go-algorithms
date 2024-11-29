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
	divider := 10
	for {
		if len(partitions[0]) == n {
			return partitions[0]
		}
		for i := 0; i < n; i++ {
			partition := result[i] % divider
			partitions[partition] = append(partitions[partition], result[i])
		}
		result = make([]int, 0)
		for i := 0; i < 10; i++ {
			result = append(result, partitions[i]...)
		}
	}
}
