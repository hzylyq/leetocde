package SameBsts

import "testing"

func TestSameBsts(t *testing.T) {
	res := SameBsts([]int{10, 15, 8, 12, 94, 81, 5, 2, 11},
		[]int{10, 8, 5, 15, 2, 12, 11, 94, 81})
	t.Log(res)
}
