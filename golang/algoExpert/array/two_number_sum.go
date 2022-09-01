package array

import "sort"

// O(n^2) time
// O(1) space
func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	for i, val1 := range array {
		for j := i + 1; j < len(array); j++ {
			val2 := array[j]
			if val1+val2 == target {
				return []int{val1, val2}
			}
		}
	}
	return nil
}

func TwoNumberSum2(array []int, target int) []int {
	numMap := make(map[int]bool)
	// Write your code here.
	for _, val1 := range array {
		val2 := target - val1
		if _, ok := numMap[val2]; ok {
			return []int{val1, val2}
		}
		numMap[val1] = true
	}
	return nil
}

func TwoNumberSum3(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		sum := array[left] + array[right]
		if sum == target {
			return []int{array[left], array[right]}
		} else if sum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}

