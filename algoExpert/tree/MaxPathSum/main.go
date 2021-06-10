package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func MaxPathSum(tree *BinaryTree) int {
	// Write your code here.
	return -1
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
