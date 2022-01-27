package dp

// NumberOfWaysToTraverseGraph O(2^(n+m)) time / O(m+n)
func NumberOfWaysToTraverseGraph(width int, height int) int {
	if width == 1 || height == 1 {
		return 1
	}

	return NumberOfWaysToTraverseGraph(width-1, height) + NumberOfWaysToTraverseGraph(width, height-1)
}

func NumberOfWaysToTraverseGraph2(width int, height int) int {
	array := make([][]int, width+1)
	for i := range array {
		array[i] = make([]int, height+1)
	}

	for i := 1; i <= width; i++ {
		for j := 1; j <= height; j++ {
			if i == 1 || j == 1 {
				array[i][j] = 1
			} else {
				array[i][j] = array[i-1][j] + array[i][j-1]
			}
		}
	}

	return array[width][height]
}
