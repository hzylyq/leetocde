package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{4, 3, 2, 5, 1, 9}
	BubbleSort(arr)
	t.Log(arr)
}

func TestMaxArea(t *testing.T) {
	MaxArea([][]int{})
}

func TestSearch(t *testing.T) {
	search([]int{-1, 0, 3, 5, 9, 12}, 9)
}
