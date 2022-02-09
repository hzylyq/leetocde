package array

import "sort"

func SortedSquaredArray(array []int) []int {
	sort.Ints(array)
	res := make([]int, 0)
	for _, val := range array {
		res = append(res, val*val)
	}
	return res
}

func SortedSquaredArray2(array []int) []int {
	res := make([]int, 0)

	smallValIdx := 0
	largerValIdx := len(array) - 1

	for idx := len(array) - 1; idx >= 0; idx-- {
		smallVal := array[smallValIdx]
		largeVal := array[largerValIdx]

		if abs(smallVal) > abs(largeVal) {
			res[idx] = smallVal * smallVal
			smallValIdx++
		} else {
			res[idx] = largeVal * largeVal
			largerValIdx--
		}
	}

	return res
}
