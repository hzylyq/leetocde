package breadthfirstsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZigzagLevelOrder(t *testing.T) {
	node := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	res := ZigzagLevelOrder(node)
	assert.Equal(t, res, [][]int{
		{3},
		{20, 9},
		{15, 7},
	})
}

func TestZigzagLevelOrder2(t *testing.T) {
	vals := []int{5, 1, 6, 8}
	for i, n := 0, len(vals); i < n/2; i++ {
		vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
	}

	assert.Equal(t, vals, []int{8, 6, 1, 5})
}
