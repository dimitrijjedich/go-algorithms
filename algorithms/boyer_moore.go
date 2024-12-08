package algorithms

func Search(needle []rune, hayStack []rune) int {
	n := len(needle)
	maxLen := len(hayStack)
	skipTable := shiftArray(needle)

	position := n - 1
	for position < maxLen {
		println("Current position", position)
		i := n - 1
		for ; i >= 0 && needle[i] == hayStack[position]; i, position = i-1, position-1 {
		}
		if i == -1 {
			return position + 1
		}
		_, exist := skipTable[hayStack[position]]
		if exist {
			position = position + skipTable[hayStack[position]]
		} else {
			position = position + n
		}
	}
	return -1
}

func shiftArray(needle []rune) map[rune]int {
	result := make(map[rune]int)
	n := len(needle)
	for i := n - 1; i >= 0; i-- {
		_, exists := result[needle[i]]
		if !exists {
			result[needle[i]] = n - 1 - i
		}
	}
	return result
}
