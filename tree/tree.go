package tree

// 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 树的遍历
func Traversal(root *TreeNode) []int {
	var res []int
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val) // 放在这是前序
		helper(node.Left)
		// res = append(res, node.Val) // 放在这是中序
		helper(node.Right)
		// res = append(res, node.Val) // 放在这是后序
	}
	helper(root)
	return res
}

// 前序遍历(非递归)
func preorderTraversal(root *TreeNode) []int {
	var res []int
	ptr := root
	stack := make([]*TreeNode, 0)
	for ptr != nil || len(stack) > 0 {
		if ptr != nil {
			res = append(res, ptr.Val)
			stack = append(stack, ptr) // 入栈
			ptr = ptr.Left             // 访问左子树
		} else {
			top := stack[len(stack)-1] // 出栈
			stack = stack[:len(stack)-1]
			ptr = top.Right // 访问右子树
		}
	}
	return res
}

// 95. 不同的二叉搜索树 II
func GenerateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}

	return dfs(1, n)
}

func dfs(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}

	res := make([]*TreeNode, 0)
	for i := left; i <= right; i++ {
		leftTrees := dfs(left, i-1)
		rightTrees := dfs(i+1, right)
		for _, lNode := range leftTrees {
			for _, rNode := range rightTrees {
				node := &TreeNode{
					Val:   i,
					Left:  lNode,
					Right: rNode,
				}
				res = append(res, node)
			}
		}
	}
	return res
}

// 108有序数组转化为二叉搜索数
func sortedArrayToBST(numList []int) *TreeNode {
	if len(numList) == 0 {
		return nil
	}

	mid := len(numList) / 2
	node := &TreeNode{
		Val:   numList[mid],
		Left:  sortedArrayToBST(numList[:mid]),
		Right: sortedArrayToBST(numList[mid+1:]),
	}
	return node
}
