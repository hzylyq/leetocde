package backtrack

// 46. 全排列
func permute(nums []int) [][]int {
	var ans [][]int
	var curr []int

	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			ans = append(ans, curr)
			return
		}


	}

	dfs(0)

	return ans
}
