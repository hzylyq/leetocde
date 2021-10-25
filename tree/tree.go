package tree

import (
	"math"
	"strconv"
	"strings"
)

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

// 94. Binary Tree Inorder Traversal
func inorderTraversal(root *TreeNode) []int {
	array := make([]int, 0)

	if root == nil {
		return array
	}
	array = append(array, inorderTraversal(root.Left)...)
	array = append(array, root.Val)
	array = append(array, inorderTraversal(root.Right)...)

	return array
}

// 95. 不同的二叉搜索树 II
func GenerateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}

	return dfs(1, n)
}

// NumTrees 96. 不同的二叉搜索树
// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
func NumTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
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

// isValidBST 98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return isHelperValidBST(root, math.MinInt32, math.MaxInt32)
}

func isHelperValidBST(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}

	return isHelperValidBST(node.Left, lower, node.Val) && isHelperValidBST(node.Right, node.Val, upper)
}

// 100. 相同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

// 101. 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type helperFunc func(*TreeNode, *TreeNode) bool
	var helper helperFunc
	helper = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p != nil && q != nil {
			return p.Val == q.Val && helper(p.Left, q.Right) && helper(p.Right, q.Left)
		}
		return false
	}
	return helper(root.Left, root.Right)
}

// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	var index int
	// 记录下第一个val的idx
	for idx, val := range inorder {
		if val == root.Val {
			index = idx
			break
		}
	}
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}

// 非递归
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	stack := []*TreeNode{root}
	var inorderIndex int

	for i := 1; i < len(preorder); i++ {
		preOrderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{
				Val:   preOrderVal,
				Left:  nil,
				Right: nil,
			}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{
				preOrderVal, nil, nil,
			}
			stack = append(stack, node.Right)
		}
	}
	return root
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

// flatten 114. 二叉树展开为链表
func flatten(root *TreeNode) {
	list := helperFlatten(root)

	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func helperFlatten(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	res := make([]*TreeNode, 0)
	res = append(res, root)
	res = append(res, helperFlatten(root.Left)...)
	res = append(res, helperFlatten(root.Right)...)
	return res
}

// MaxPathSum 124. 二叉树中的最大路径和
// 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
// 路径和 是路径中各节点值的总和。
// 给你一个二叉树的根节点root, 返回其最大路径和.
func MaxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var helperFunc func(root *TreeNode) int
	helperFunc = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := max(helperFunc(node.Left), 0)
		rightMax := max(helperFunc(node.Right), 0)

		newMaxVal := leftMax + node.Val + rightMax
		maxSum = max(newMaxVal, maxSum)

		return node.Val + max(leftMax, rightMax)
	}

	helperFunc(root)

	return maxSum
}

// SumNumbers 129. 求根到叶子节点数字之和
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

// 211. 添加与搜索单词 - 数据结构设计
// 请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
// 实现词典类 WordDictionary ：
// WordDictionary() 初始化词典对象
// void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
// bool search(word) 如果数据结构中存在字符串与word 匹配，则返回 true；否则，返回 false 。word 中可能包含一些 '.' ，每个. 都可以表示任何一个字母。
type trieNode struct {
	children [26]*trieNode
	isEnd    bool
}

func (t *trieNode) Insert(word string) {

}

type wordDictionary struct {
	children [26]
}

func NewWordDictionary() wordDictionary {

}

func (w *wordDictionary) AddWord(word string) {

}

func (w *wordDictionary) Search(word string) bool {

}

// 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 235. 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res := root
	for {
		if p.Val < res.Val && q.Val < res.Val {
			res = res.Left
		} else if p.Val > res.Val && q.Val > res.Val {
			res = res.Right
		} else {
			return res
		}
	}
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 层序遍历
	stack := []*TreeNode{root}
	var res string
	for len(stack) > 0 {
		length := len(stack)

		for _, node := range stack {
			if node == nil {
				res += "null,"

			} else {
				res += strconv.Itoa(node.Val)
			}

			if node.Left == nil {
				res += "null"
			} else {
				stack = append(stack, node.Left)
			}

			if node.Right == nil {
				res += "null"
			} else {
				stack = append(stack, node.Right)
			}
		}
		stack = stack[length:]
	}

	res += "]"
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	if len(sp) == 0 {
		return nil
	}
	var helper func() *TreeNode
	helper = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}

		val, _ := strconv.Atoi(sp[0])

		return &TreeNode{
			Val:   val,
			Left:  helper(),
			Right: helper(),
		}
	}

	return helper()
}

type Node struct {
	Val      int
	Children []*Node
}

// ChildLevelOrder 429. N 叉树的层序遍历
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

// PathSum 437. 路径总和 III
// 给定一个二叉树的根节点 root，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的路径的数目。
// 路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
func PathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return PathSum(root.Left, targetSum) + pathSumHelperFunc(root, targetSum) + PathSum(root.Right, targetSum)
}

func pathSumHelperFunc(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}

	var sum int
	if node.Val == targetSum {
		sum++
	}
	sum += pathSumHelperFunc(node.Left, targetSum-node.Val)
	sum += pathSumHelperFunc(node.Right, targetSum-node.Val)

	return sum
}

// 538. 把二叉搜索树转换为累加树
// 给出二叉搜索树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点node的新值等于原树中大于或等于node.val的值之和。
// 提醒一下，二叉搜索树满足下列约束条件：
// 节点的左子树仅包含键小于节点键的节点。
// 节点的右子树仅包含键大于节点键的节点。
// 左右子树也必须是二叉搜索树。
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	// 后序遍历
	var helperFunc func(node *TreeNode)
	helperFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		helperFunc(node.Right)
		node.Val += sum
		sum = node.Val
		helperFunc(node.Left)
	}
	helperFunc(root)

	return root
}

// 671. 二叉树中第二小的节点
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	ans := -1
	rootVal := root.Val

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || ans != -1 && node.Val >= ans {
			return
		}
		if node.Val > rootVal {
			ans = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
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

// 501. 二叉搜索树中的众数
// todo
func findMode(root *TreeNode) []int {
	var base, count, maxCount int
	var res []int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			res = append(res, base)
		} else if count > maxCount {
			maxCount = count
			res = []int{base}
		}
	}

	curr := root
	for curr != nil {
		if curr.Left == nil {
			update(curr.Val)
			curr = curr.Right
			continue
		}
		pre := curr.Left
		if pre.Right != nil && pre.Right != curr {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = curr
			curr = curr.Right
		} else {
			pre.Right = nil
			update(curr.Val)
			curr = curr.Right
		}
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

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
func (tree *BST) Insert(value int) *BST {
	if tree.Contains(value) {
		return tree
	}

	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}

func (tree *BST) Contains(value int) bool {
	if tree == nil {
		return false
	}

	if tree.Value == value {
		return true
	}

	if tree.Value > value {
		return tree.Left.Contains(value)
	} else {
		return tree.Right.Contains(value)
	}
}

func (tree *BST) Remove(value int) *BST {
	if !tree.Contains(value) {
		return tree
	}

	tree.remove(value, nil)
	return tree
}

func (tree *BST) remove(val int, parent *BST) {
	if tree.Value > val {
		if tree.Left != nil {
			tree.Left.remove(val, tree)
		}
	} else if tree.Value < val {
		if tree.Right != nil {
			tree.Right.remove(val, tree)
		}
	} else {
		if tree.Left != nil && tree.Right != nil {
			tree.Value = tree.Right.getMinValue()
			tree.Right.remove(tree.Value, tree)
		} else if parent == nil {
			if tree.Left != nil {
				tree.Value = tree.Left.Value
				tree.Left = tree.Left.Left
				tree.Right = tree.Left.Right
			} else if tree.Right != nil {
				tree.Value = tree.Right.Value
				tree.Left = tree.Right.Left
				tree.Right = tree.Right.Right
			}
		} else if parent.Left == tree {
			if tree.Left != nil {
				parent.Left = tree.Left
			} else {
				parent.Left = tree.Right
			}
		} else if parent.Right == tree {
			if tree.Left != nil {
				parent.Right = tree.Left
			} else {
				parent.Right = tree.Right
			}
		}
	}
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue()
}

func RangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	var sum int
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	if root.Val > low {
		sum += RangeSumBST(root.Left, low, high)
	}
	if root.Val < high {
		sum += RangeSumBST(root.Right, low, high)
	}

	return sum
}

func MinDiffInBST(root *TreeNode) int {
	list := inOrder(root)

	res := math.MaxInt32
	for i := 1; i < len(list); i++ {
		mini := list[i] - list[i-1]
		if mini < res {
			res = mini
		}
	}
	return res
}

// 剑指 Offer 55 - I. 二叉树的深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	var res []float64
	queue := []*TreeNode{root}

	currNum, nextLevelNum, count, sum := 1, 0, 1, 0

	for len(queue) > 0 {
		if currNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			currNum--
			sum += node.Val
			queue = queue[1:]
		}
		if currNum == 0 {
			res = append(res, float64(sum)/float64(count))
			currNum, count, sum, nextLevelNum = nextLevelNum, nextLevelNum, 0, 0
		}
	}

	return res
}

// 110. Balanced Binary Tree
func IsBalanced(root *TreeNode) bool {
	_, isBalanced := HeightAndIsBalanced(root)
	return isBalanced
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	miniDepth := math.MaxInt32
	if root.Left != nil {
		miniDepth = min(minDepth(root.Left), miniDepth)
	}
	if root.Right != nil {
		miniDepth = min(minDepth(root.Right), miniDepth)

	}

	return miniDepth + 1
}

func HeightAndIsBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	lf, bf := HeightAndIsBalanced(root.Left)
	if !bf {
		return 0, false
	}
	lr, br := HeightAndIsBalanced(root.Right)
	if !br {
		return 0, false
	}

	return max(lf, lr) + 1, abs(lf, lr) <= 1
}

// 111. Minimum Depth of Binary Tree
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}
	if root.Right == nil {
		return MinDepth(root.Left) + 1

	}

	return min(MinDepth(root.Left), MinDepth(root.Right)) + 1
}

// 230. 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	var count, res int

	type helperFunc func(root *TreeNode, k int)
	var helper helperFunc
	helper = func(root *TreeNode, k int) {
		if root == nil {
			return
		}

		helper(root.Left, k)

		count++
		if count == k {
			res = root.Val
			return
		}
		helper(root.Right, k)
	}
	helper(root, k)
	return res
}

// 404. Sum of Left Leaves
func sumOfLeftLeaves(root *TreeNode) int {
	queue := []*TreeNode{root}
	var res int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			if node.Left.Left == nil && node.Left.Right == nil {
				res += node.Left.Val
			}
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

func sumOfLeftLeaves2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			res += root.Left.Val
		} else {
			res += sumOfLeftLeaves2(root.Left)
		}
	}
	if root.Right != nil && !(root.Right.Left == nil && root.Right.Right == nil) {
		res += sumOfLeftLeaves2(root.Right)
	}

	return res
}

// 530. Minimum Absolute Difference in BST
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := inOrder(root)
	tmp := math.MaxInt32
	for i := 1; i <= len(res)-1; i++ {
		min := res[i] - res[i-1]
		if min < tmp {
			tmp = min
		}
	}
	return tmp
}

// 617. 合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	tree := &TreeNode{
		Val: root1.Val + root2.Val,
	}
	tree.Left = mergeTrees(root1.Left, root2.Left)
	tree.Right = mergeTrees(root1.Right, root2.Right)
	return tree
}

type TreeInfo struct {
	Height   int
	Diameter int
}

func diameterOfBinaryTree(root *TreeNode) int {
	return getTreeInfo(root).Diameter
}

func getTreeInfo(node *TreeNode) TreeInfo {
	if node == nil {
		return TreeInfo{}
	}

	leftTreeInfo := getTreeInfo(node.Left)
	rightTreeInfo := getTreeInfo(node.Right)

	longPath := leftTreeInfo.Height + rightTreeInfo.Height
	maxDiameter := max(leftTreeInfo.Diameter, rightTreeInfo.Diameter)
	currentMaxDiameter := max(maxDiameter, longPath)
	currentHeight := max(leftTreeInfo.Height, rightTreeInfo.Height) + 1

	return TreeInfo{Diameter: currentMaxDiameter, Height: currentHeight}
}

func diameterOfBinaryTree2(root *TreeNode) int {
	var ans int
	type helperFunc func(node *TreeNode) int
	var helper helperFunc

	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := helper(node.Left)
		r := helper(node.Right)

		ans = max(ans, l+r+1)

		return max(l, r) + 1
	}
	helper(root)

	return helper(root)
}

func minDiffInBST(root *TreeNode) int {
	// var mini = math.MaxInt32
	// var pre *TreeNode
	//
	//
	//
	// inOrder:= func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// }
	return 0
}

// 965. 单值二叉树
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Val != root.Left.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Val != root.Right.Val {
			return false
		}
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 993. 二叉树的堂兄弟节点
func IsCousins(root *TreeNode, x, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	var dfs func(node, parent *TreeNode, depth int)
	dfs = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}

		if xFound && yFound {
			return
		}
		dfs(node.Left, node, depth+1)
		if xFound && yFound {
			return
		}
		dfs(node.Right, node, depth+1)
	}
	dfs(root, nil, 0)
	return xDepth == yDepth && xParent != yParent
}

// 1022. 从根到叶的二进制数之和
func sumRootToLeaf(root *TreeNode) int {
	type helperFunc func(root *TreeNode, sum int) int
	var helper helperFunc

	helper = func(root *TreeNode, sum int) int {
		if root == nil {
			return 0
		}
		sum = sum<<1 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return helper(root.Left, sum) + helper(root.Right, sum)
	}

	return helper(root, 0)
}
