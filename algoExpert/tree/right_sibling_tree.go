package tree

func RightSiblingTree(root *BinaryTree) *BinaryTree {
	mutate(root, nil, false)
	return root
}

func mutate(node, parent *BinaryTree, isLeftChild bool) {
	if node == nil {
		return
	}

	left, right := node.Left, node.Right
	mutate(left, node, true)
	if parent == nil {
		node.Right = nil
	} else if isLeftChild {
		node.Right = parent.Right
	} else {
		if parent.Right == nil {
			node.Right = nil
		} else {
			node.Right = parent.Right.Left
		}
	}

	mutate(right, node, false)
}
