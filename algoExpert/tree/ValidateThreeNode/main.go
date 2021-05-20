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

// isDescendant 判断target是否为node的后继节点
func isDescendant2(node, target *BST) bool {
	currentNode := node
	for currentNode != nil {
		if currentNode == target {
			return true
		}
		if currentNode.Value > target.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	
	return false
}

func ValidateThreeNodes3(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	searchOne := nodeOne
	searchTwo := nodeThree
	
	for {
		foundThreeFromOne := searchOne == nodeThree
		foundOneFromThree := searchTwo == nodeOne
		foundNodeTwo := searchOne == nodeTwo || searchTwo == nodeTwo
		finishedSearching := searchOne == nil && searchTwo == nil
		if foundThreeFromOne || foundOneFromThree || foundNodeTwo || finishedSearching {
			break
		}
		
		if searchOne != nil {
			if searchOne.Value > nodeTwo.Value {
				searchOne = searchOne.Left
			} else {
				searchOne = searchOne.Right
			}
		}
		if searchTwo != nil {
			if searchTwo.Value > nodeTwo.Value {
				searchTwo = searchTwo.Left
			} else {
				searchTwo = searchTwo.Right
			}
		}
	}
	
	foundNodeFromOther := searchOne == nodeThree || searchTwo == nodeOne
	foundNodeTwo := searchOne == nodeTwo || searchTwo == nodeOne
	if !foundNodeTwo || foundNodeFromOther {
		return false
	}
	
	if searchOne == nodeTwo {
		return searchForTarget(nodeTwo, nodeThree)
	}
	return searchForTarget(nodeTwo, nodeOne)
}

func searchForTarget(node, target *BST) bool {
	currentNode := node
	for currentNode != nil {
		if currentNode == target {
			return true
		}
		if currentNode.Value > target.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	
	return false
}