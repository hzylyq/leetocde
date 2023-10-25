package bfs

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ZigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	
	if root == nil {
		return nil
	}
	
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		var vals []int
		q := queue
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		
		// 奇数层的元素翻转
		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		ans = append(ans, vals)
	}
	return ans
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 117. 填充每个节点的下一个右侧节点指针 II
func connect(root *Node) *Node {
	
	return root
}

// 200. 岛屿数量
func numIslands(grid [][]byte) int {
	
	return 0
}

// 733. 图像渲染
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	currColor := image[sr][sc]
	if currColor == color {
		return image
	}
	
	return image
}

// 797. 所有可能的路径
func allPathsSourceTarget(graph [][]int) [][]int {
	return nil
}

// 1091. 二进制矩阵中的最短路径
func shortestPathBinaryMatrix(grid [][]int) int {
	return 0
}

// 2698. 求一个整数的惩罚数
func punishmentNumber(n int) int {
	var dfs func(string, int, int, int) bool
	dfs = func(s string, pos int, tot int, target int) bool {
		if pos == len(s) {
			return tot == target
		}
		sum := 0
		for i := pos; i < len(s); i++ {
			sum = sum*10 + int(s[i]-'0')
			if sum+tot > target {
				break
			}
			
			if dfs(s, i+1, sum+tot, target) {
				return true
			}
		}
		return false
	}
	
	res := 0
	
	for i := 0; i <= n; i++ {
		if dfs(strconv.Itoa(i*i), 0, 0, i) {
			res += i * i
		}
	}
	
	return res
}
