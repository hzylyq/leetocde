package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
