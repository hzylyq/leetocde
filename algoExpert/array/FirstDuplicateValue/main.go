package main

func FirstDuplicateValue(array []int) int {
	seen := make(map[int]bool, len(array))
	for _, val := range array {
		if seen[val] {
			return val
		}
		seen[val] = true
	}
	return -1
}

func FirstDuplicateValue2(array []int) int {
	for _, val := range array {
		absVal := abs(val - 1)
		if array[absVal] < 0 {
			return val
		}
		array[absVal] *= -1
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
