package string

// 5. 最长回文子串
// 给你一个字符串 s，找到 s 中最长的回文子串
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)

		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {

	}
	return left + 1, right - 1
}

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

	var stack []rune
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

// balancedStringSplit 1221. 分割平衡字符串
// 在一个 平衡字符串 中，'L' 和 'R' 字符的数量是相同的。
// 给你一个平衡字符串s，请你将它分割成尽可能多的平衡字符串。
// 注意：分割得到的每个字符串都必须是平衡字符串。
// 返回可以通过分割得到的平衡字符串的 最大数量 。
func balancedStringSplit(s string) int {
	var ans int
	var d int
	for _, char := range s {
		if char == 'R' {
			d++
		} else {
			d--
		}

		if d == 0 {
			ans++
		}
	}
	return ans
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
