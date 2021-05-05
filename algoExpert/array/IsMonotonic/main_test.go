package main

import "testing"

func TestIsMonotonic(t *testing.T) {
	array := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	t.Log(IsMonotonic(array))
}
