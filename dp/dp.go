package dp

import (
	"math"
)

// 53. 最大子序和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 62. 不同路径
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// 63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 1
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if i == 0 && j == 0 {
					dp[i][j] = 1
				} else if i == 0 {
					dp[i][j] = dp[i][j-1]
				} else if j == 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}

	return dp[m-1][n-1]
}

// 64. 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func MinCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

// 70. 爬楼梯
func climbStairs(n int) int {
	if n < 2 {
		return n
	}

	p, q, r := 1, 1, 0
	for i := 2; i <= n; i++ {
		r = p + q
		p = q
		q = r
	}

	return r
}

// LCP 22. 黑白方格画
func paintingPlan(n int, k int) int {
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

// 78. 子集
func subsets(nums []int) [][]int {
	var set []int
	var ans [][]int

	var dfs func(int)
	dfs = func(i int) {
		// 遍历结束
		if i == len(nums) {
			ans = append(ans, append([]int{}, set...))
			return
		}

		// 选择当前位置
		set = append(set, nums[i])
		dfs(i + 1)
		set = set[:len(set)-1]
		dfs(i + 1)
	}

	dfs(0)
	return ans
}

// Generate 118. 杨辉三角
func Generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 {
				res[i][j] = 1
			} else if j == i {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}

	return res
}

// 122. 买卖股票的最佳时机 II
func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// 不持有股票的利润，持有股票的利润
	v1, v2 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		currV1 := max(v1, v2+prices[i])
		currV2 := max(v2, v1-prices[i])

		v1 = currV1
		v2 = currV2
	}

	return v1
}

// 139. 单词拆分
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool, len(wordDict))
	for i := 0; i < len(wordDict); i++ {
		wordSet[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

// 152. 乘积最大子数组
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF

		maxF = max(max(mx*nums[i], mn*nums[i]), nums[i])
		minF = min(min(mx*nums[i], mn*nums[i]), nums[i])
		ans = max(ans, maxF)
	}

	return ans
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

// 221. 最大正方形
// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0

	for i, row := range matrix {
		dp[i] = make([]int, len(matrix[i]))
		for j, col := range row {
			dp[i][j] = int(col - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

// 300. 最长递增子序列
func lengthOfLIS(nums []int) int {

}

// 309. 最佳买卖股票时机含冷冻期
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		newF0 := max(f0, f2-prices[i])
		newF1 := f0 + prices[i]
		newF2 := max(f2, f1)

		f0 = newF0
		f1 = newF1
		f2 = newF2
	}

	return max(max(f0, f1), f2)
}

// 416. 分割等和子集
func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	sum, max := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	if max > target {
		return false
	}

	return false
}

// 473. 火柴拼正方形
func makeSquare(matchsticks []int) bool {
	// 首先判断火柴总数能不能被4整除
	// var total int
	// for _, stick := range matchsticks {
	//	total += stick
	// }
	//
	// if total%4 != 0 {
	//	return false
	// }
	//
	// // 每条边的长度
	// tLen := total / 4
	// dp := make([]int, 1<<len(matchsticks))
	// for i := 1; i < len(dp); i++ {
	//	dp[i] = -1
	// }
	//
	// for s := 1; s < len(dp); s++ {
	//
	// }

	return false
}

// 494. 目标和
// 回溯 2^n
func findTargetSumWays(nums []int, target int) int {
	var count int

	var backtrack func(int, int)
	backtrack = func(index int, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}

		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}

	backtrack(0, 0)
	return count
}

// 动态规划O(n) 背包问题
func findTargetSumWays2(nums []int, target int) int {

	return 0
}

// 509. 斐波那契数
// F(0) = 0，F(1) = 1
// F(n) = F(n - 1) + F(n - 2)，其中 n > 1
func fib(n int) int {
	dp := make([]int, 0, n)
	dp = append(dp, 0, 1)

	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}

// 647. 回文子串
func countSubstrings(s string) int {
	size := len(s)
	ans := 0

	for i := 0; i < 2*size-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < size && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}

	return ans
}

// 1262. 可被三整除的最大和
func maxSumDivThree(nums []int) int {
	dp := [3]int{}

	for i := 0; i < len(nums); i++ {
		a, b, c := dp[0]+nums[i], dp[1]+nums[i], dp[2]+nums[i]
		dp[a%3] = max(dp[a%3], a)
		dp[b%3] = max(dp[b%3], b)
		dp[c%3] = max(dp[c%3], c)
	}

	return dp[0]
}

// 剑指 Offer 10- I. 斐波那契数列
func fib2(n int) int {
	dp := make([]int, 0, n)
	dp = append(dp, 0, 1)

	for i := 2; i <= n; i++ {
		dp = append(dp, (dp[i-1]+dp[i-2])%(1e9+7))
	}
	return dp[n]
}

// O(n)
func fib3(n int) int {
	if n < 2 {
		return n
	}

	p, q, r := 0, 0, 1

	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % (1e9 + 7)
	}
	return r
}

// 剑指 Offer 10- II. 青蛙跳台阶问题
func numWays(n int) int {
	if n < 2 {
		return 1
	}

	p, q, r := 1, 1, 0
	for i := 2; i <= n; i++ {
		r = (p + q) % (1e9 + 7)
		p = q % (1e9 + 7)
		q = r % (1e9 + 7)
	}

	return r % (1e9 + 7)
}

// 剑指 Offer II 091. 粉刷房子
func minCost(costs [][]int) int {
	dp := costs[0]

	for _, cost := range costs[1:] {
		dpNew := make([]int, 3)

		for j, c := range cost {
			dpNew[j] = min(dp[(j+1)%3], dp[(j+2)%3]) + c
		}

		dp = dpNew
	}

	return min(min(dp[0], dp[1]), dp[2])
}

// 剑指 Offer 47. 礼物的最大价值
func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = max(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

// 剑指 Offer 63. 股票的最大利润
func maxProfit(prices []int) int {
	var miniVal = math.MaxInt32
	var profit int

	for _, price := range prices {
		if price < miniVal {
			miniVal = price
		} else if price-miniVal > profit {
			profit = price - miniVal
		}
	}

	return profit
}

// 剑指 Offer II 088. 爬楼梯的最少成本
func minCostClimbingStairs(cost []int) int {
	// dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])

	n := len(cost)
	var (
		prev int
		curr int
	)

	for i := 2; i <= n; i++ {
		next := min(prev+cost[i-2], curr+cost[i-1])
		prev = curr
		curr = next
	}

	return curr
}

// 剑指 Offer II 089. 房屋偷盗
func robI(nums []int) int {
	// dp[n] = max(dp[n-2]+nums[n], dp[n-1])
	if len(nums) < 2 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	prev := nums[0]
	curr := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		next := max(prev+nums[i], curr)
		prev = curr
		curr = next
	}

	return curr
}

//
// func updateMatrix(mat [][]int) [][]int {
// 	for i := 0; i < len(mat); i++ {
// 		for j := 0; j < len(mat[i]); j++ {
// 			if mat[]
// 		}
// 	}

// 	return mat
// }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
