package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	assert.Equal(t, MinCostClimbingStairs([]int{10, 15, 20}), 15)
}

func TestPaintingPlan(t *testing.T) {
	assert.Equal(t, paintingPlan(3, 8), 9)
}

func TestUniquePaths(t *testing.T) {
	assert.Equal(t, uniquePaths(3, 7), 28)
}

func TestGenerate(t *testing.T) {
	Generate(5)
}

func TestWordBreak(t *testing.T) {
	wordBreak("leetcode", []string{"leet", "code"})
}
