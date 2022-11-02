package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	t.Log(res)
}

func TestMerge(t *testing.T) {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func TestCheckSubarraySum(t *testing.T) {
	res := checkSubarraySum([]int{0, 0}, 6)
	assert.Equal(t, res, true)
}

func TestLetterCombinations(t *testing.T) {
	res := LetterCombinations("")
	t.Log(res)
}

func TestTwoSum(t *testing.T) {
	twoSum([]int{2, 7, 11, 15}, 9)
}

func TestIsSubsequence(t *testing.T) {
	isSubsequence("axc", "ahbgdc")
}

func TestArrayRankTransform(t *testing.T) {
	arrayRankTransform([]int{40, 10, 20, 30})
}

func TestShuff(t *testing.T) {
	shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
}

func TestThreeSum2(t *testing.T) {
	threeSum2([]int{-1, 0, 1, 2, -1, -4})
}
