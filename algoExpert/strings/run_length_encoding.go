package strings

import "strconv"

func RunLengthEncoding(str string) string {
	var encodeStringCharacters []byte
	currentRunLength := 1

	for i := 1; i < len(str); i++ {
		currChar := str[i]
		prevChar := str[i-1]

		if currChar != prevChar || currentRunLength == 9 {
			encodeStringCharacters = append(encodeStringCharacters, strconv.Itoa(currentRunLength)[0])
			encodeStringCharacters = append(encodeStringCharacters, prevChar)
			currentRunLength = 0
		}
		currentRunLength++
	}

	encodeStringCharacters = append(encodeStringCharacters, strconv.Itoa(currentRunLength)[0])
	encodeStringCharacters = append(encodeStringCharacters, str[len(str)-1])
	return string(encodeStringCharacters)
}
