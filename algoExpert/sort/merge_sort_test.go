package sort

import "testing"

func TestMergeSort(t *testing.T) {
	res := []int{8, 5, 2, 9, 5, 6, 3}
	MergeSort(res)
}
