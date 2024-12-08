package algorithms

func Search(needle []rune, haystack []rune) int {
	n := len(needle)
	maxLen := len(haystack)
	skipTable := skipTable(needle)

	position := n - 1
	for position < maxLen {
		i := n - 1
		for ; i >= 0 && needle[i] == haystack[position]; i, position = i-1, position-1 {
		}
		if i == -1 {
			return position + 1
		}
		if shift, exist := skipTable[haystack[position]]; exist {
			position = position + shift
		} else {
			position = position + n
		}
	}
	return -1
}

func skipTable(needle []rune) map[rune]int {
	result := make(map[rune]int)
	n := len(needle)
	for i := 0; i < n; i++ {
		result[needle[i]] = n - 1 - i
	}
	return result
}
