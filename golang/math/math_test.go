package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseComplement(t *testing.T) {
	res := BitwiseComplement(5)
	assert.Equal(t, res, 2)
}

func TestCheckPerfectNumber(t *testing.T) {
	// assert.Equal(t, checkPerfectNumber(28), true)
}

func TestPlusOne(t *testing.T) {
	assert.Equal(t, plusOne([]int{1, 2, 3}), []int{1, 2, 4})
}

func TestPivotInteger(t *testing.T) {
	res := pivotInteger(8)
	t.Log(res)
}
