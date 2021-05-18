package ValidateThreeNode

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func ValidateThreeNodes(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	if isDescendant(nodeTwo, nodeOne) {
		return isDescendant(nodeThree, nodeTwo)
	}
	if isDescendant(nodeTwo, nodeThree) {
		return isDescendant(nodeOne, nodeTwo)
	}
	return false
}

// isDescendant 判断target是否为node的后继节点
func isDescendant(node, target *BST) bool {
	if node == nil {
		return false
	}
	if node == target {
		return true
	}

	if node.Value > target.Value {
		return isDescendant(node.Left, target)
	}
	return isDescendant(node.Right, target)
}
