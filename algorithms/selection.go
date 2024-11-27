package algorithms

func SelectionSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)
	for i := 0; i < n; i++ {
		selected := i
		for j := i + 1; j < n; j++ {
			if result[selected] > result[j] {
				selected = j
			}
		}
		result[i], result[selected] = result[selected], result[i]
	}
	return result
}
