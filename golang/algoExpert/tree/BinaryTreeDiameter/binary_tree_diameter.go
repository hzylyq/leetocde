package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

type TreeInfo struct {
	Height   int
	Diameter int
}

// 对于任意一个节点来说，他的最长路劲等于子树的深度相加
func BinaryTreeDiameter(tree *BinaryTree) int {
	return getTreeInfo(tree).Diameter
}

func getTreeInfo(node *BinaryTree) TreeInfo {
	if node == nil {
		return TreeInfo{}
	}

	leftTreeInfo := getTreeInfo(node.Left)
	rightTreeInfo := getTreeInfo(node.Right)

	longPath := leftTreeInfo.Height + rightTreeInfo.Height
	maxDiameter := max(leftTreeInfo.Diameter, rightTreeInfo.Diameter)
	currentMaxDiameter := max(maxDiameter, longPath)
	currentHeight := max(leftTreeInfo.Height, rightTreeInfo.Height) + 1

	return TreeInfo{Diameter: currentMaxDiameter, Height: currentHeight}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
