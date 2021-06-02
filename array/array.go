package array

import (
	"sort"
)

// 27. 移除元素
func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}

// 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	
	// 二分查找
	m := len(matrix)
	n := len(matrix[0])
	
	i := sort.Search(m*n, func(i int) bool {
		return matrix[i/n][i%n] >= target
	})
	return i < m*n && matrix[i/n][i%n] == target
}

// 80. 删除有序数组中的重复项 II
func removeDuplicates(nums []int) int {
	// 快慢指针
	length := len(nums)
	if length <= 2 {
		return length
	}
	
	fast, slow := 2, 2
	for fast < length {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// 88. 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p, q := 0, 0
	for {
		if p == m {
			sorted = append(sorted, nums2[q:]...)
			break
		}
		if q == n {
			sorted = append(sorted, nums1[p:]...)
			break
		}
		
		if nums1[p] < nums2[q] {
			sorted = append(sorted, nums1[p])
			p++
		} else {
			sorted = append(sorted, nums2[q])
			q++
		}
	}
	
	copy(nums1, sorted)
}

// 217. 存在重复元素
func containsDuplicate(nums []int) bool {
	set := make(map[int]bool, len(nums))
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = true
	}
	return false
}

// 523. 连续的子数组和 前缀和加求余
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	
	mp := map[int]int{0: -1}
	remainer := 0
	
	for i, num := range nums {
		remainer = (remainer + num) % k
		if prevIdx, ok := mp[remainer]; ok {
			if i-prevIdx >= 2 {
				return true
			}
		} else {
			mp[remainer] = i
		}
	}
	return false
}

// 1480. Running Sum of 1d Array
func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	var total int
	for k, v := range nums {
		res[k] = total + v
		total += v
	}
	return res
}
