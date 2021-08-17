package base

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoNumberSum(t *testing.T) {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	res := TwoNumberSum(array, target)
	t.Log(res)
}

func TestTwoNumberSum2(t *testing.T) {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	res := TwoNumberSum2(array, target)
	t.Log(res)
}

func TestTwoNumberSum3(t *testing.T) {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	res := TwoNumberSum3(array, target)
	t.Log(res)
}

func TestIsValidSubsequence(t *testing.T) {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	t.Log(IsValidSubsequence2(array, sequence))
}

func TestBinarySearch(t *testing.T) {
	array := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}

	t.Log(BinarySearch(array, 33))
}

func TestGetNthFib(t *testing.T) {
	t.Log(GetNthFib(2))
	t.Log(GetNthFib(3))
	t.Log(GetNthFib(6))
}

func TestGetNthFib2(t *testing.T) {
	t.Log(GetNthFib2(6))
}

func TestFindThreeLargestNumbers(t *testing.T) {
	t.Log(FindThreeLargestNumbers([]int{1, 2, 3, 5, 1, 23}))
}

func TestHammingDistance(t *testing.T) {
	assert.Equal(t, hammingDistance(1, 4), 2)
}
