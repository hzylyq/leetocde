package search

import "sort"

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	var (
		l = 0
		r = len(nums) - 1
	)

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			if nums[0] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[length-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	leftMost := sort.SearchInts(nums, target)
	if leftMost == len(nums) || nums[leftMost] != target {
		return []int{-1, -1}
	}
	rightMost := sort.SearchInts(nums, target+1) - 1
	return []int{leftMost, rightMost}
}

// 35. 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
func searchInsert(nums []int, target int) int {
	ans := len(nums)

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

// 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	// 二分查找

	var (
		m = len(matrix)
		n = len(matrix[0])
	)

	left := 0
	right := m*n - 1
	for left <= right {
		mid := (right-left)/2 + left
		x := matrix[mid/n][mid%n]
		if x < target {
			left = mid + 1
		} else if x > target {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}
