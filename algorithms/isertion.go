package algorithms

func InsertionSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	for i := 0; i < n; i++ {
		for j, element := i-1, result[i]; j >= 0; j-- {
			if element < result[j] {
				result[j], result[j+1] = result[j+1], result[j]
			} else {
				result[j+1] = element
				break
			}
		}
	}
	return result
}
