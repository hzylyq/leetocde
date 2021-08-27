package stack

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	var ans int
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 739. 每日温度
// 请根据每日 气温 列表 temperatures ，请计算在每一天需要等几天才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 1 暴力
func dailyTemperatures(temperatures []int) []int {
	// res := make([]int, len(temperatures))
	//
	// for i := 0; i<len(temperatures); i++ {
	// 	for j := i+1; j<len(temperatures);j++ {
	// 		if
	// 	}
	// }
}
