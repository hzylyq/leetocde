package backtrack

import "sort"

// LetterCombinations 17. 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。
// 所有组合 回溯
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var phoneMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var combs []string

	var backtrack func(digits string, index int, comb string)

	backtrack = func(digits string, index int, comb string) {
		if index == len(digits) {
			combs = append(combs, comb)
			return
		}

		digit := string(digits[index])
		letters := phoneMap[digit]

		for i := 0; i < len(letters); i++ {
			backtrack(digits, index+1, comb+string(letters[i]))
		}
	}

	backtrack(digits, 0, "")

	return combs
}

// 22. 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
func generateParenthesis(n int) []string {
	var res []string
	var tmp []byte

	var backoff func(int, int)
	backoff = func(open, close int) {
		if len(tmp) == n*2 {
			res = append(res, string(append([]byte(nil), tmp...)))
			return
		}

		if open < n {
			tmp = append(tmp, '(')
			backoff(open+1, close)
			tmp = tmp[:len(tmp)-1]
		}
		if close < open {
			tmp = append(tmp, ')')
			backoff(open, close+1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	backoff(0, 0)

	return res
}

// 39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var comb []int

	var backoff func(target int, idx int)
	backoff = func(target int, idx int) {
		if idx == len(candidates) {
			return
		}

		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		backoff(target, idx+1)
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			backoff(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}

	backoff(target, 0)

	return ans
}

// 40. 组合总和 II
func combinationSum2(candidates []int, target int) [][]int {

	return nil
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

// 47. 全排列 II
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	var set []int
	vis := make([]bool, len(nums))
	var backoff func(int)
	backoff = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int(nil), set...))
			return
		}

		for i, v := range nums {
			if vis[i] || (i > 0 && !vis[i-1] && v == nums[i-1]) {
				continue
			}

			set = append(set, v)
			vis[i] = true
			backoff(cur + 1)
			vis[i] = false
			set = set[:len(set)-1]
		}
	}

	backoff(0)
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

// 79. 单词搜索
func exist(board [][]byte, word string) bool {
	type pair struct {
		x int
		y int
	}

	var directions = []pair{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}

	}

	for i, row := range board {
		for j := range row {
			check(i, j)
		}
	}

	return false
}

// 90. 子集 II
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var set []int

	sort.Ints(nums)

	var backoff func(bool, int)
	backoff = func(choosePre bool, curr int) {
		if curr == len(nums) {
			res = append(res, append([]int(nil), set...))
			return
		}

		backoff(false, curr+1)

		if !choosePre && curr > 0 && nums[curr] == nums[curr-1] {
			return
		}

		set = append(set, nums[curr])
		backoff(true, curr+1)
		set = set[:len(set)-1]
	}

	backoff(false, 0)
	return res
}
