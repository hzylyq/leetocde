package tree

func FlattenBinaryTree(root *BinaryTree) *BinaryTree {
	inOrderNodes := getNodeInOrder(root)
	for i := 0; i < len(inOrderNodes)-1; i++ {
		leftNode := inOrderNodes[i]
		rightNode := inOrderNodes[i+1]
		leftNode.Right = rightNode
		rightNode.Left = leftNode
	}

	return inOrderNodes[0]
}

func getNodeInOrder(tree *BinaryTree) []*BinaryTree {
	inOrderNodes := make([]*BinaryTree, 0)

	if tree != nil {
		inOrderNodes = append(inOrderNodes, getNodeInOrder(tree.Left)...)
		inOrderNodes = append(inOrderNodes, tree)
		inOrderNodes = append(inOrderNodes, getNodeInOrder(tree.Right)...)
	}

	return inOrderNodes
}
