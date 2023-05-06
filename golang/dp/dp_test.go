package dp

import (
	"log"
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

func TestMaxSumDivThree(t *testing.T) {
	maxSumDivThree([]int{3, 6, 5, 1, 8})
}

func TestCoinChange(t *testing.T) {
	print(coinChange([]int{1, 2, 5}, 11))
}

func TestMinimumTotal(t *testing.T) {
	minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
}

func TestNumSquares(t *testing.T) {
	res := numSquares(13)
	log.Print(res)
}

func TestIsInterleave(t *testing.T) {
	res := isInterleave("", "", "")
	assert.Equal(t, true, res)
}
