package FindSuccessor

type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	// Write your code here.
	treeList := inOrder(tree)
	for idx, tree := range treeList {
		if tree == node && idx < len(treeList)-1 {
			return treeList[idx+1]
		}
	}

	return nil
}

func inOrder(tree *BinaryTree) []*BinaryTree {
	if tree == nil {
		return nil
	}

	var res []*BinaryTree
	res = append(res, inOrder(tree.Left)...)
	res = append(res, tree)
	res = append(res, inOrder(tree.Right)...)
	return res
}

func FindSuccessor2(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	if node.Right != nil {
		return getLeftMostChild(node.Right)
	}
	return getRightMostChild(node)
}

func getLeftMostChild(node *BinaryTree) *BinaryTree {
	currentNode := node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode
}

func getRightMostChild(node *BinaryTree) *BinaryTree {
	currentNode := node
	for currentNode.Parent != nil && currentNode.Parent.Right == currentNode {
		currentNode = currentNode.Parent
	}
	return currentNode.Parent
}
