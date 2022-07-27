package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOuterParentheses(t *testing.T) {
	res := RemoveOuterParentheses("(()())(()))")
	assert.Equal(t, res, "()()()")
}

func TestDailyTemperatures(t *testing.T) {
	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
}
