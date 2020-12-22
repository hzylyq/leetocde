package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseComplement(t *testing.T) {
	res := BitwiseComplement(5)
	assert.Equal(t, res, 2)
}
