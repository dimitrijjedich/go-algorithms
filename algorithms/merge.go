package algorithms

func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	center := n / 2

	left := MergeSort(arr[:center])
	right := MergeSort(arr[0:center])

	nLeft := len(left)
	nRight := len(right)
	i, j := 0, 0
	result := make([]int, 0)
	for i < nLeft && j < nRight {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return result
}
