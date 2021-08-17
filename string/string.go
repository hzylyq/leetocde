package string

// 20. 有效的括号
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pair := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []rune{}
	for _, ch := range s {
		if _, ok := pair[ch]; ok {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || ch != pair[stack[len(stack)-1]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

//
func longestCommonPrefix(strs []string) string {
	return ""
}

// checkRecord 551. 学生出勤记录 I
func checkRecord(s string) bool {
	a, latest := 0, 0
	for _, val := range s {
		if val == 'A' {
			a++
			if a >= 2 {
				return false
			}
		}
		if val == 'L' {
			latest++
			if latest >= 3 {
				return false
			}
		} else {
			latest = 0
		}
	}

	return true
}

// 1576. 替换所有的问号
func ModifyString(s string) string {
	temp := make([]byte, len(s))
	copyStr := []byte(s)

	for idx, char := range s {
		if char != '?' {
			temp[idx] = s[idx]
			continue
		}

		for j := byte('a'); j <= 'z'; j++ {
			if (idx == 0 || j != copyStr[idx-1]) && (idx == len(s)-1 || j != copyStr[idx+1]) {
				temp[idx] = j
				copyStr[idx] = j
				break
			}
		}
	}

	return string(temp)
}

// 1678. 设计 Goal 解析器
func Interpret(command string) string {
	var result string
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			result += "G"
		} else if command[i] == '(' {
			if command[i+1] == 'a' {
				result += "al"
				i += 3
			} else {
				result += "o"
				i += 1
			}
		}
	}

	return result
}

//
func FirstUniqChar(s string) int {
	var res [26]int

	for _, ch := range s {
		res[ch-'a']++
	}

	for i, ch := range s {
		if res[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
