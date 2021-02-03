package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTrees(t *testing.T) {
	res := GenerateTrees(3)
	t.Log(res)
}

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
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

	res := LevelOrder(root)
	t.Log(res)
}

func TestSumNumbers(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}
	res := SumNumbers(root)
	assert.Equal(t, res, 25)
}
