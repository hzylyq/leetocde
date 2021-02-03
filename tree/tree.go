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

// 129. 求根到叶子节点数字之和
func SumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return helper(0, root)
}

func helper(num int, node *TreeNode) int {
	var res int
	num = num*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return num
	}
	if node.Left != nil {
		res += helper(num, node.Left)
	}
	if node.Right != nil {
		res += helper(num, node.Right)
	}
	return res
}

// 二叉树的层序遍历
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	res := make([][]int, 0)
	for len(queue) > 0 {
		n := len(queue)
		r := make([]int, 0)
		for i := 0; i < n; i++ {
			r = append(r, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		res = append(res, r)
	}
	return res
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

// 429. N 叉树的层序遍历
func ChildLevelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	res := make([][]int, 0)

	for len(queue) > 0 {
		n := len(queue)
		r := make([]int, 0)
		for i := 0; i < n; i++ {
			r = append(r, queue[0].Val)
			queue = append(queue, queue[0].Children...)
			queue = queue[1:]
		}
		res = append(res, r)
	}
	return res
}

// 面试题 04.03. 特定深度节点链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}

	res := make([]*ListNode, 0)
	list := make([]*TreeNode, 0)
	list = append(list, tree)

	for len(list) > 0 {
		size := len(list)

		var node, head *ListNode
		for i := 0; i < size; i++ {
			if i == 0 {
				node = &ListNode{
					Val: list[i].Val,
				}
				head = node
			} else {
				node.Next = &ListNode{
					Val: list[i].Val,
				}
				node = node.Next
			}
			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
		}
		if head != nil {
			res = append(res, head)
		}
		list = list[size:]
	}
	return res
}

// 589. N叉树的前序遍历
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	res = append(res, root.Val)
	for _, child := range root.Children {
		res = append(res, preorder(child)...)
	}
	return res
}
