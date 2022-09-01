package tree

func AllKindsOfNodeDepths(root *BinaryTree) int {
	// Write your code here.
	return helper(root, 0)
}

func helper(root *BinaryTree, depth int) int {
	if root == nil {
		return 0
	}

	depthSum := (depth * (depth + 1)) / 2
	return depthSum + helper(root.Left, depth+1) + helper(root.Right, depth+1)
}
