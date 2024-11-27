package algorithms

func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	center := n / 2

	left := MergeSort(arr[:center])
	right := MergeSort(arr[0:center])
	return arr
}
