package algorithms

func GnomeSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	for i := 0; i < n-1; {
		if result[i] > result[i+1] {
			result[i], result[i+1] = result[i+1], result[i]
			i = i - 1
			if i < 0 {
				i = 0
			}
		} else {
			i = i + 1
		}
	}
	return result
}
