package algorithms

func QuickSort(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}

	left := make([]int, 0)
	right := make([]int, 0)
	pivot := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	result := append(left, pivot)
	result = append(result, right...)
	return result
}
