package backtrack

// 46. 全排列
func permute(nums []int) [][]int {
	var res [][]int

	var backoff func([]int, int)

	backoff = func(tmp []int, idx int) {
		if idx == len(nums)-1 {
			res = append(res, tmp)
			return
		}

		var before []int
		copy(before, tmp)

		for i := idx; i < len(tmp); i++ {
			before[i], before[idx] = before[idx], before[i]
			backoff(before, idx+1)
		}
	}

	backoff(nums, 0)
	return res
}
