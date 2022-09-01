package sort

import (
	"math"
)

func RadixSort(array []int) []int {
	if len(array) <= 0 {
		return array
	}

	maxNumber := max(array)
	digit := 0

	for maxNumber/(int(math.Pow(float64(10), float64(digit)))) > 0 {
		countSortArray(array, digit)
		digit += 1
	}

	return array
}

func countSortArray(array []int, digit int) {
	sortedArray := make([]int, len(array))
	countArray := make([]int, 10)

	digitCol := int(math.Pow(10, float64(digit)))

	for _, num := range array {
		countIdx := (num / digitCol) % 10
		countArray[countIdx] += 1
	}

	for idx := 1; idx < 10; idx++ {
		countArray[idx] += countArray[idx-1]
	}

	for idx := len(array) - 1; idx >= 0; idx-- {
		countIdx := (array[idx] / digitCol) % 10
		countArray[countIdx] -= 1
		sortedIdx := countArray[countIdx]
		sortedArray[sortedIdx] = array[idx]
	}

	for idx := range array {
		array[idx] = sortedArray[idx]
	}
}

func max(array []int) int {
	currentMax := array[0]

	for _, val := range array {
		if val > currentMax {
			currentMax = val
		}
	}
	return currentMax
}
