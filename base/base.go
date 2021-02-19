package base

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

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	var seqIndex, arrIndex int
	for arrIndex < len(array) && seqIndex < len(sequence) {
		if array[arrIndex] == sequence[seqIndex] {
			seqIndex++
			arrIndex++
		} else {
			arrIndex++
		}
	}

	return seqIndex == len(sequence)
}

func IsValidSubsequence2(array []int, sequence []int) bool {
	// Write your code here.
	var seqIndex int
	for _, val := range array {
		if seqIndex == len(sequence) {
			break
		}
		if val == sequence[seqIndex] {
			seqIndex++
		}
	}

	return seqIndex == len(sequence)
}

func BinarySearch(array []int, target int) int {
	// Write your code here.
	low, high := 0, len(array)-1
	for low < high {
		mid := low + (high-low)/2
		numMid := array[mid]
		if target == numMid {
			return mid
		}

		if target < numMid {
			high = mid
		} else {
			low = mid
		}
	}

	return -1
}

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Ints(coins)
	currentChangeCreated := 0
	for _, coin := range coins {
		if coin > currentChangeCreated+1 {
			return currentChangeCreated + 1
		}
		currentChangeCreated += coin
	}

	return currentChangeCreated + 1
}
