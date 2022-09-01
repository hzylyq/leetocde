package main

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	// Write your code here.
	idx1 := len(array1)
	idx2 := len(array2)
	var res []int
	smallest, current := math.MaxInt32, math.MinInt32
	i, j := 0, 0
	for i < idx1 && j < idx2 {
		first, second := array1[i], array2[j]
		if first < second {
			current = second - first
			i++
		} else if first > second {
			current = first - second
			j++
		} else {
			return []int{first, second}
		}

		if current < smallest {
			smallest = current
			res = []int{first, second}
		}

	}

	return res
}
