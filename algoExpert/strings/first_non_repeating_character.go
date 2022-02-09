package strings

func FirstNonRepeatingCharacter(str string) int {
	charMap := make(map[rune]int)

	for _, char := range str {
		charMap[char]++
	}

	for idx, char := range str {
		if charMap[char] == 1 {
			return idx
		}
	}
	return -1
}
