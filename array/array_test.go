package array

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func TestCheckSubarraySum(t *testing.T) {
	res := checkSubarraySum([]int{0, 0}, 6)
	assert.Equal(t, res, true)
}
