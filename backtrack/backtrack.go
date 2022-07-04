package backtrack

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
