package backtrack

// 22. 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
func generateParenthesis(n int) []string {
	var res []string

	var backoff func([]byte, int, int)
	backoff = func(tmp []byte, open, close int) {
		if len(tmp) == n*2 {
			curr := make([]byte, len(tmp))
			copy(curr, []byte(tmp))
			res = append(res, string(curr))
			return
		}

		if open < n {
			tmp = append(tmp, '(')
			backoff(tmp, open+1, close)
			tmp = tmp[:len(tmp)-1]
		}
		if close < open {
			tmp = append(tmp, ')')
			backoff(tmp, open, close+1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	backoff(nil, 0, 0)

	return res
}

// 46. 全排列
func permute(nums []int) [][]int {
	var res [][]int

	var backoff func(int, []int)

	backoff = func(idx int, tmp []int) {
		if idx == len(nums)-1 {
			curr := make([]int, len(tmp))
			copy(curr, tmp)
			res = append(res, curr)
			return
		}

		for i := idx; i < len(tmp); i++ {
			tmp[i], tmp[idx] = tmp[idx], tmp[i]
			backoff(idx+1, tmp)
			tmp[i], tmp[idx] = tmp[idx], tmp[i]
		}
	}

	backoff(0, nums)
	return res
}

// 77. 组合
func combine(n int, k int) [][]int {
	var res [][]int

	var backoff func(int, []int)

	backoff = func(idx int, tmp []int) {
		if len(tmp) == k {
			curr := make([]int, len(tmp))
			copy(curr, tmp)
			res = append(res, curr)
			return
		}

		for i := idx + 1; i <= n; i++ {
			tmp = append(tmp, i)
			backoff(i, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	backoff(0, nil)

	return res
}
