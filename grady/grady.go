package grady

import "math"

// 55. 跳跃游戏
func canJump(nums []int) bool {
	var rightMax int

	for i := 0; i < len(nums); i++ {
		if i <= rightMax {
			rightMax = max(rightMax, i+nums[i])
			if rightMax >= len(nums)-1 {
				return true
			}
		}
	}

	return false
}

// 334. 递增的三元子序列
// 给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
// 如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first, second := nums[0], math.MaxInt32
	for i := 1; i < len(nums); i++ {
		if nums[i] > second {
			return true
		}

		if nums[i] > first {
			second = nums[i]
		} else {
			first = nums[i]
		}
	}

	return false
}

// 807. 保持城市天际线
func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)

	rowMax := make([]int, n)
	colMax := make([]int, n)

	for i, row := range grid {
		for j, col := range row {
			rowMax[i] = max(rowMax[i], col)
			colMax[j] = max(colMax[j], col)
		}
	}

	ans := 0

	for i, row := range grid {
		for j, col := range row {
			ans += min(rowMax[i], colMax[j]) - col
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
