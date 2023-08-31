package string

import (
	"strconv"
)

// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	rk, ans := -1, 0

	existCh := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(existCh, s[i-1])
		}
		for rk+1 < len(s) && !existCh[s[rk+1]] {
			rk++
			existCh[s[rk]] = true
		}

		ans = max(ans, rk-i+1)
	}

	return ans
}

// LongestPalindrome 5. 最长回文子串
// 给你一个字符串 s，找到 s 中最长的回文子串
func LongestPalindrome(s string) string {
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

func LongestPalindromeDp(s string) string {
	// if len(s) <= 2 {
	// 	return s
	// }

	// Dp := [][]int{}

	// // todo

	// for {
	// 	break
	// }

	return ""
}

// 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
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

// 28. 找出字符串中第一个匹配项的下标
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}

	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}

	return -1
}

// 58. 最后一个单词的长度
func lengthOfLastWord(s string) int {
	var res int

	length := len(s)
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' && res == 0 {
			continue
		} else if s[i] != ' ' {
			res++
		} else {
			return res
		}
	}

	return res
}

// 187. 重复的DNA序列
// 所有 DNA 都由一系列缩写为 'A'，'C'，'G' 和 'T' 的核苷酸组成，例如："ACGAATTCCG"。
// 在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。
// 编写一个函数来找出所有目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。
func findRepeatedDnaSequences(s string) []string {
	const l = 10
	bin := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

	n := len(s)
	if n <= l {
		return nil
	}

	x := 0
	for _, ch := range s[:l-1] {
		x = x<<2 | bin[byte(ch)]
	}
	cnt := map[int]int{}
	var ans []string
	for i := 0; i <= n-l; i++ {
		x = (x<<2 | bin[s[i+l-1]]) & (1<<(l*2) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			ans = append(ans, s[i:i+l])
		}
	}
	return ans
}

// 415. 字符串相加
func AddStrings(num1 string, num2 string) string {
	var add int
	var ans string

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		res := x + y + add
		add = res / 10
		ans = strconv.Itoa(res%10) + ans
	}

	return ans
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

// 1446. 连续字符 (easy)
// 给你一个字符串s, 字符串的「能量」定义为: 只包含一种字符的最长非空子字符串的长度. 请你返回字符串的能量.
// 输入: s = "leetcode"
// 输出: 2
// 解释: 子字符串 "ee" 长度为 2 ，只包含字符 'e'
func maxPower(s string) int {
	ans, cnt := 1, 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
			if cnt > ans {
				ans = cnt
			}
		} else {
			cnt = 1
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

// Interpret 1678. 设计 Goal 解析器
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

// 剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	curr := []byte(s)

	var res []byte
	res = append(res, curr[n:]...)
	res = append(res, curr[:n]...)

	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
