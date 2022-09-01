package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalancedBrackets(t *testing.T) {
	res := BalancedBrackets("{}gawgw()")
	assert.Equal(t, res, true)
}
