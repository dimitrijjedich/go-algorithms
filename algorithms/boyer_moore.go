package algorithms

import "errors"

func Search(needle []rune, haystack []rune) (int, error) {
	n := len(needle)
	if n == 0 {
		return -1, errors.New("needle can not be empty")
	}
	maxLen := len(haystack)
	if maxLen == 0 {
		return -1, errors.New("haystack can not be empty")
	}
	if maxLen < n {
		return -1, errors.New("haystack can not be longer than needle")
	}
	skipTable := skipTable(needle)

	position := n - 1
	for position < maxLen {
		i := n - 1
		for ; i >= 0 && needle[i] == haystack[position]; i, position = i-1, position-1 {
		}
		if i == -1 {
			return position + 1, nil
		}
		if shift, exist := skipTable[haystack[position]]; exist {
			position = position + shift
		} else {
			position = position + n
		}
	}
	return -1, errors.New("needle not found")
}

func skipTable(needle []rune) map[rune]int {
	result := make(map[rune]int)
	n := len(needle)
	for i := 0; i < n; i++ {
		result[needle[i]] = n - 1 - i
	}
	return result
}
