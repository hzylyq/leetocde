package main

import "math"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func MaxPathSum(tree *BinaryTree) int {
	maxPathSum := math.MinInt32

	var maxGain func(binaryTree *BinaryTree) int

	maxGain = func(node *BinaryTree) int {
		if node == nil {
			return 0
		}

		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		currentSum := node.Value + leftGain + rightGain
		maxPathSum = max(currentSum, maxPathSum)

		return node.Value + max(leftGain, rightGain)
	}

	maxGain(tree)
	return maxPathSum
}

func helper(root *BinaryTree) int {
	if root == nil {
		return 0
	}

	// todo

	return root.Value + helper(root.Left) + helper(root.Right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
