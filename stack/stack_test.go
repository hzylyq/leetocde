package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOuterParentheses(t *testing.T) {
	res := RemoveOuterParentheses("(()())(()))")
	assert.Equal(t, res, "()()()")
}



