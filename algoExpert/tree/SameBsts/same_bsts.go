package SameBsts

import "math"

func SameBsts(arrayOne, arrayTwo []int) bool {
	if len(arrayOne) != len(arrayTwo) {
		return false
	}
	if len(arrayOne) == 0 && len(arrayTwo) == 0 {
		return true
	}
	if arrayOne[0] != arrayTwo[0] {
		return false
	}
	smallerOne := GetSmaller(arrayOne)
	smallerTwo := GetSmaller(arrayTwo)

	largerOne := GetLagerOrEqual(arrayOne)
	largerTwo := GetLagerOrEqual(arrayTwo)

	return SameBsts(smallerOne, smallerTwo) && SameBsts(largerOne, largerTwo)
}

func GetSmaller(array []int) []int {
	var res []int
	for i := 1; i < len(array); i++ {
		if array[i] < array[0] {
			res = append(res, array[i])
		}
	}
	return res
}

func GetLagerOrEqual(array []int) []int {
	var res []int
	for i := 1; i < len(array); i++ {
		if array[i] >= array[0] {
			res = append(res, array[i])
		}
	}
	return res
}

func SameBsts2(arrayOne, arrayTwo []int) bool {
	return areSameBsts(arrayOne, arrayTwo, 0, 0, math.MinInt32, math.MaxInt32)
}

func areSameBsts(arrayOne, arrayTwo []int, rootIdxOne, rootIdxTwo int, minVal, maxVal int) bool {
	if rootIdxOne == -1 || rootIdxTwo == -1 {
		return rootIdxOne == rootIdxTwo
	}
	if arrayOne[rootIdxOne] != arrayTwo[rootIdxTwo] {
		return false
	}
	leftRootIdxOne := getIdxOfFirstSmaller(arrayOne, rootIdxOne, minVal)
	leftRootIdxTwo := getIdxOfFirstSmaller(arrayTwo, rootIdxTwo, minVal)
	rightRootIdxOne := getIdxOfFirstBiggerOrEqual(arrayOne, rootIdxOne, maxVal)
	rightRootIdxTwo := getIdxOfFirstBiggerOrEqual(arrayTwo, rootIdxTwo, maxVal)

	currentVal := arrayOne[rootIdxOne]
	leftAreSame := areSameBsts(arrayOne, arrayTwo, leftRootIdxOne, leftRootIdxTwo, minVal, currentVal)
	rightAreSame := areSameBsts(arrayOne, arrayTwo, rightRootIdxOne, rightRootIdxTwo, currentVal, maxVal)
	return leftAreSame && rightAreSame
}

func getIdxOfFirstSmaller(array []int, startingIdx, minVal int) int {
	for i := startingIdx + 1; i < len(array); i++ {
		if array[i] < array[startingIdx] && array[i] >= minVal {
			return i
		}
	}
	return -1
}

func getIdxOfFirstBiggerOrEqual(array []int, startingIdx, maxVal int) int {
	for i := startingIdx + 1; i < len(array); i++ {
		if array[i] >= array[startingIdx] && array[i] < maxVal {
			return i
		}
	}
	return -1
}
