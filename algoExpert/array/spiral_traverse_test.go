package array

import "testing"

func TestSpiralTraverse(t *testing.T) {
	array := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7}}
	res := SpiralTraverse(array)
	array = [][]int{{1}}
	res = SpiralTraverse(array)

	array = [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}}
	res = SpiralTraverse(array)

	t.Log(res)
}
