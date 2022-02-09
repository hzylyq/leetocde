package tree

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func FindNodesDistanceK(tree *BinaryTree, target int, k int) []int {
	nodesToParents := make(map[int]*BinaryTree)
	populateNodeToParents(tree, nodesToParents, nil)
	targetNode := getNodeFromValue(target, tree, nodesToParents)

	return BFSForNodeK(targetNode, nodesToParents, k)
}

func BFSForNodeK(targetNode *BinaryTree, nodeToParents map[int]*BinaryTree, k int) []int {
	type item struct {
		node     *BinaryTree
		distance int
	}

	queue := []item{{node: targetNode, distance: 0}}
	seen := map[int]bool{targetNode.Value: true}
	var currItem item
	for len(queue) > 0 {
		currItem, queue = queue[0], queue[1:]
		currentNode, distanceFromTarget := currItem.node, currItem.distance

		if distanceFromTarget == k {
			nodeDistanceK := make([]int, 0)
			for _, node := range queue {
				nodeDistanceK = append(nodeDistanceK, node.node.Value)
			}
			nodeDistanceK = append(nodeDistanceK, currentNode.Value)
			return nodeDistanceK
		}

		connectNodes := []*BinaryTree{currentNode.Left, currentNode.Right, nodeToParents[currentNode.Value]}
		for _, node := range connectNodes {
			if node == nil {
				continue
			}
			if seen[node.Value] {
				continue
			}

			seen[node.Value] = true
			queue = append(queue, item{
				node:     node,
				distance: distanceFromTarget + 1,
			})
		}
	}

	return []int{}
}

func populateNodeToParents(node *BinaryTree, nodeToParents map[int]*BinaryTree, parent *BinaryTree) {
	if node != nil {
		nodeToParents[node.Value] = parent
		populateNodeToParents(node.Left, nodeToParents, node)
		populateNodeToParents(node.Right, nodeToParents, node)
	}
}

func getNodeFromValue(value int, tree *BinaryTree, nodeToParents map[int]*BinaryTree) *BinaryTree {
	if tree.Value == value {
		return tree
	}

	nodeParent := nodeToParents[value]
	if nodeParent.Left != nil && nodeParent.Left.Value == value {
		return nodeParent.Left
	}

	return nodeParent.Right
}
