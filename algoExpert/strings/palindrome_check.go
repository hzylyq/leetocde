package strings

// IsPalindrome O(n)/O(n)
func IsPalindrome(str string) bool {
	reversed := ""

	for i := len(str) - 1; i > 0; i-- {
		reversed += string(str[i])
	}

	return str == reversed
}

// IsPalindrome2 O(n)/O(1)
func IsPalindrome2(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		j := len(str) - i - 1
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
