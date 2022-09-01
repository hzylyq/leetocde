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

func TestValidateStackSequences(t *testing.T) {
	validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})
}
