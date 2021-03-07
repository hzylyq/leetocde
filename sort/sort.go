package sort

func BubbleSort(array []int) []int {
	// Write your code here.
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}

// func InsertionSort(array []int) []int {
// 	// Write your code here.
// 	for i := 0; i < len(array); i++ {
// 		for j :=
//
// 	}
// 	return nil
// }

func MaxArea(num [][]int) int {
	if len(num) == 0 {
		return 0
	}

	m, n := len(num), len(num[0])
	left := make([][]int, m)

	for i, row := range num {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == 0 {
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
	for i, row := range num {
		for j, v := range row {
			if v == 0 {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width,left[k][j])
				area = max(area,(i-k+1)*width/2)
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
