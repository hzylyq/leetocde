package tree

// 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tree *BST) InOrderTraverse(array []int) []int {
	if tree == nil {
		return array
	}
	array = tree.Left.InOrderTraverse(array)
	array = append(array, tree.Value)
	array = tree.Right.InOrderTraverse(array)
	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	if tree == nil {
		return array
	}
	array = append(array, tree.Value)
	array = tree.Left.PreOrderTraverse(array)
	array = tree.Right.PreOrderTraverse(array)

	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	if tree == nil {
		return array
	}
	array = tree.Left.PostOrderTraverse(array)
	array = tree.Right.PostOrderTraverse(array)
	array = append(array, tree.Value)

	return array
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

// 剑指 Offer 32 - II. 从上到下打印二叉树 II
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var list []*TreeNode
	list = append(list, root)

	res := make([][]int, 0)
	for len(list) > 0 {
		size := len(list)
		var data []int
		for i := 0; i < size; i++ {
			data = append(data, list[i].Val)

			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
		}
		res = append(res, data)
		list = list[size:]
	}
	return res
}

// 897. 递增顺序查找树
func IncreasingBST(root *TreeNode) *TreeNode {
	list := inOrder(root)
	res := new(TreeNode)
	curr := new(TreeNode)
	curr = res
	for _, val := range list {
		curr.Right = new(TreeNode)
		curr.Right.Val = val
		curr = curr.Right
	}
	return res.Right
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	res = append(res, inOrder(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inOrder(root.Right)...)
	return res
}

// 剑指 Offer 54. 二叉搜索树的第k大节点
func KthLargest(root *TreeNode, k int) int {
	var ans, count int

	type helperFunc func(*TreeNode, int)
	var helper helperFunc
	helper = func(root *TreeNode, k int) {
		if root.Right != nil {
			helper(root.Right, k)
		}
		count++

		if count == k {
			ans = root.Val
			return
		}
		if root.Left != nil {
			helper(root.Left, k)
		}

		return
	}
	helper(root, k)
	return ans
}

// find closest val in BST
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	if tree == nil {
		return -1
	}

	return tree.findCloseValue(target, tree.Value)
}

func (tree *BST) findCloseValue(target, closest int) int {
	currentNode := tree

	for currentNode != nil {
		if abs(target, closest) > abs(target, currentNode.Value) {
			closest = currentNode.Value
		}
		if target < currentNode.Value {
			currentNode = currentNode.Left
		} else if target > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			break
		}
	}

	return closest
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	if root == nil {
		return -1
	}

	var sum, depth int
	queue := []*BinaryTree{
		root,
	}

	for len(queue) > 0 {
		length := len(queue)
		sum += len(queue) * depth
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]
		depth++
	}

	// Write your code here.
	return sum
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	var sums []int
	calculationBranchSums(root, 0, &sums)
	return sums
}

func calculationBranchSums(node *BinaryTree, runningSum int, sums *[]int) {
	if node == nil {
		return
	}

	newRunningSum := runningSum + node.Value
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, newRunningSum)
	}
	calculationBranchSums(node.Left, newRunningSum, sums)
	calculationBranchSums(node.Right, newRunningSum, sums)
}

type ChildNode struct {
	Name     string
	Children []*ChildNode
}

func (n *ChildNode) DepthFirstSearch(array []string) []string {
	// Write your code here.
	array = append(array, n.Name)
	for _, node := range n.Children {
		array = append(array, node.DepthFirstSearch(array)...)
	}
	return nil
}
