package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModifyString(t *testing.T) {
	assert.Equal(t, ModifyString("j?qg??b"), "jaqgacb")
}

func TestInterpret(t *testing.T) {

}
