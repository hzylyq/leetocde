package ThreeNumSum

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	res := make([][]int, 0)
	for i := 0; i < len(array)-2; i++ {
		left, right := i+1, len(array)-1
		for left < right {
			sum := array[i] + array[left] + array[right]
			if target == sum {
				res = append(res, []int{array[i], array[left], array[right]})
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return res
}
