package array

func SpiralTraverse(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}

	var res []int

	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1

	for startRow <= endRow && startCol <= endCol {
		for col := startCol; col <= endCol; col++ {
			res = append(res, array[startRow][col])
		}

		for row := startRow + 1; row <= endRow; row++ {
			res = append(res, array[row][endCol])
		}

		for col := endCol - 1; col >= startCol; col-- {
			if startRow == endRow {
				break
			}
			res = append(res, array[endRow][col])
		}

		for row := endRow - 1; row > startRow; row-- {
			if startCol == endCol {
				break
			}
			res = append(res, array[row][startCol])
		}

		startRow++
		endRow--
		startCol++
		endCol--
	}

	return res
}
