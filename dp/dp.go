package dp

func GenerateParenthesis(n int) []string {
	var res []string

	return nil
}

func helpGenerateParenthesis(idx, n, sum, add int, str string, res []string) {
	sum += add
	if sum < 0 {
		return
	}

	if add == 1 {
		str += "("
	} else {
		str += ")"
	}

	if idx == n*2 {

	}
}

func MinCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// LCP 22. 黑白方格画
func PaintingPlan(n int, k int) int {
	if k == n*n {
		return 1
	}
	if k == 0 {
		return 1
	}
	if k < n {
		return 0
	}

	res := 0

	// 2. 普通情况 : 假设涂了i行,j列 那么一共涂的方块数应该是 in + j(n - i), 那么k = in + j(n - i) 可以得到 j = (k-in)/(n - i)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if n*i+j*(n-i) == k {
				res += C(n, i) * C(n, j)
			}
		}
	}

	return res
}

func C(base, up int) int {
	a := 1
	for i := 0; i < up; i++ {
		a *= base - i
	}

	b := 1
	for i := 1; i <= up; i++ {
		b *= i
	}
	return a / b
}

// 198. 打家劫舍 middle
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
// 链接：https://leetcode-cn.com/problems/house-robber
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// 状态专业方程 dp[n] = max(dp[n-2]+nums[n], dp[n-1])
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
