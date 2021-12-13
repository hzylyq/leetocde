package grady

// 807. 保持城市天际线
func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)

	rowMax := make([]int, n)
	colMax := make([]int, n)

	for i, row := range grid {
		for j, col := range row {
			rowMax[i] = max(rowMax[i], col)
			colMax[j] = max(colMax[j], col)
		}
	}

	ans := 0

	for i, row := range grid {
		for j, col := range row {
			ans += min(rowMax[i], colMax[j]) - col
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
