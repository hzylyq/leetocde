package stack

func BalancedBrackets(s string) bool {
	charMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []rune
	for _, val := range s {
		if val == '(' || val == '[' || val == '{' {
			stack = append(stack, val)
			continue
		}

		if val != ')' && val != ']' && val != '}' {
			continue
		}
		if len(stack) == 0 {
			stack = append(stack, val)
			continue
		}

		lastChar := stack[len(stack)-1]

		if charMap[val] == lastChar {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, val)
		}
	}

	return len(stack) == 0
}
