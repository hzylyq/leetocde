package FindSuccessor

import (
	"testing"
)

func TestFindSuccessor(t *testing.T) {
	node := &BinaryTree{
		Value:  5,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}

	tree := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value:  6,
					Left:   nil,
					Right:  nil,
					Parent: nil,
				},
				Right:  nil,
				Parent: nil,
			},
			Right:  node,
			Parent: nil,
		},
		Right: &BinaryTree{
			Value:  3,
			Left:   nil,
			Right:  nil,
			Parent: nil,
		},
		Parent: nil,
	}

	FindSuccessor(tree, node)
}

func TestFindSuccessor2(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	root.Left.Parent = root
	root.Right = &BinaryTree{Value: 3}
	root.Right.Parent = root
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Left.Parent = root.Left
	root.Left.Right = &BinaryTree{Value: 5}
	root.Left.Right.Parent = root.Left
	root.Left.Right.Left = &BinaryTree{Value: 6}
	root.Left.Right.Left.Parent = root.Left.Right
	root.Left.Right.Right = &BinaryTree{Value: 7}
	root.Left.Right.Right.Parent = root.Left.Right
	root.Left.Right.Right.Left = &BinaryTree{Value: 8}
	root.Left.Right.Right.Left.Parent = root.Left.Right.Right

	node := root.Left.Right

	actual := FindSuccessor2(root, node)
	t.Log(actual)

}
