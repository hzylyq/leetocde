package strings

import "strings"

func CaesarCipherEncryptor(str string, key int) string {
	shift, offset := rune(key%26), rune(26)

	runes := []rune(str)

	for i, char := range str {
		if char >= 'a' && char+shift <= 'z' {
			char += shift
		} else {
			char += shift - offset
		}
		runes[i] = char
	}
	return string(runes)
}

func CaesarCipherEncryptor2(str string, key int) string {
	runes := []rune(str)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i, char := range str {
		index := strings.Index(alphabet, string(char))
		if index == -1 {
			return ""
		}

		newIndex := (index + key) % 26
		runes[i] = rune(alphabet[newIndex])
	}
	return string(runes)
}
