package tree

type BinaryTree2 struct {
	Value int

	Left   *BinaryTree2
	Right  *BinaryTree2
	Parent *BinaryTree2
}

func (tree *BinaryTree2) IterativeInOrderTraversal(callback func(int)) {
	var previous, next *BinaryTree2
	current := tree
	for current != nil {
		if previous == nil || previous == current.Parent {
			if current.Left != nil {
				next = current.Left
			} else {
				callback(current.Value)
				if current.Right != nil {
					next = current.Right
				} else {
					next = current.Parent
				}
			}
		} else if previous == current.Left {
			callback(current.Value)
			if current.Right != nil {
				next = current.Right
			} else {
				next = current.Parent
			}
		} else {
			next = current.Parent
		}

		previous, current = current, next
	}
}
