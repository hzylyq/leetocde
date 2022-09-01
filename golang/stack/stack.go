package stack

import "strings"

// 71. 简化路径
// 给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为更加简洁的规范路径。
// 在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。任意多个连续的斜杠（即，'//'）都被视为单个斜杠 '/' 。 对于此问题，任何其他格式的点（例如，'...'）均被视为文件/目录名称。
// 请注意，返回的 规范路径 必须遵循下述格式：
// 始终以斜杠 '/' 开头。
// 两个目录名之间必须只有一个斜杠 '/' 。
// 最后一个目录名（如果存在）不能 以 '/' 结尾。
// 此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。
// 返回简化后得到的 规范路径 。
func simplifyPath(path string) string {
	stack := make([]string, 0)
	
	for _, name := range strings.Split(path, "/") {
		if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if name != "." && name != "" {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}

// 394. 字符串解码
// 给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
func decodeString(s string) string {
	var stk []string
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '0' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			var sub []string
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
			
			}
		}
	}
	
	return ""
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	
	var ans int
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}
	
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 739. 每日温度
// 请根据每日 气温 列表 temperatures ，请计算在每一天需要等几天才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 单调栈
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	
	stack := make([]int, 0, len(temperatures))
	
	for i := 0; i < len(temperatures); i++ {
		temp := temperatures[i]
		
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			preIdx := stack[len(stack)-1]
			
			stack = stack[:len(stack)-1]
			
			res[preIdx] = i - preIdx
		}
		stack = append(stack, i)
	}
	return res
}

// 946
func validateStackSequences(pushed []int, popped []int) bool {
	
	var st []int
	var j int
	for _, val := range pushed {
		st = append(st, val)
		
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1]
			j++
		}
	}
	
	return len(st) == 0
}

// RemoveOuterParentheses 1021. 删除最外层的括号
func RemoveOuterParentheses(s string) string {
	var res []rune
	level := 0
	
	for _, ch := range s {
		if ch == ')' {
			level--
		}
		
		if level > 0 {
			res = append(res, ch)
		}
		
		if ch == '(' {
			level++
		}
	}
	
	return string(res)
}
