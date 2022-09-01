package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallerNumbersThanCurrent(t *testing.T) {
	assert.Equal(t, SmallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}), []int{4, 0, 1, 1, 3})
}

func TestContainsNearbyDuplicate(t *testing.T) {
	containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)
}
