package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	assert.Equal(t, MinCostClimbingStairs([]int{10, 15, 20}), 15)
}
