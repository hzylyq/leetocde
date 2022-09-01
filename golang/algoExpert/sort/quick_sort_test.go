package sort

import (
	"log"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5,4,3,2,1}
	arr = QuickSort(arr)
	log.Print(arr)
}
