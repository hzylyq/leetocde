package sort

import "sort"

// BubbleSort 冒泡排序
func BubbleSort(array []int) []int {
	// Write your code here.
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}

// InsertionSort 插入排序
func InsertionSort(array []int) []int {
	for i := range array {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}

// SelectionSort 选择排序
func SelectionSort(array []int64) []int64 {
	var smallest int
	for i := 0; i < len(array); i++ {
		smallest = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				smallest = j
			}
		}

		if array[i] != array[smallest] {
			array[i], array[smallest] = array[smallest], array[i]
		}
	}
	return array
}

// QuickSort 快速排序
func QuickSort(array []int64) []int64 {

	return nil
}

func helperQuickSort(array []int64, low int, high int) {

}

func MaxArea(num [][]int) int {
	if len(num) == 0 {
		return 0
	}

	m, n := len(num), len(num[0])
	left := make([][]int, m)

	for i, row := range num {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == 0 {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	var ans int
	for i, row := range num {
		for j, v := range row {
			if v == 0 {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width/2)
			}
			ans = max(ans, area)
		}
	}

	return ans
}

// 240. 搜索二维矩阵 II
// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
// O(mn)
func searchMatrix1(matrix [][]int, target int) bool {
	for _, col := range matrix {
		for _, row := range col {
			if row == target {
				return true
			}
		}
	}
	return false
}

// O(mlogn)
func searchMatrix2(matrix [][]int, target int) bool {
	for _, col := range matrix {
		i := sort.SearchInts(col, target)
		if i < len(col) && target == col[i] {
			return true
		}
	}
	return false
}

// O(m+n)
func searchMatrix3(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

// 704. 二分查找
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid + 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
