package algorithms

func ShakerSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < n-1; i++ {
			if result[i] > result[i+1] {
				result[i], result[i+1] = result[i+1], result[i]
				sorted = false
			}
		}
		for i := n - 1; i > 0; i-- {
			if result[i] < result[i-1] {
				result[i], result[i-1] = result[i-1], result[i]
				sorted = false
			}
		}
	}
	return result
}
