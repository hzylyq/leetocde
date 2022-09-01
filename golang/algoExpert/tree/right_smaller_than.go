package tree

func RightSmallerThan(array []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(array); i++ {
		count := 0
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				count++
			}
		}
		res = append(res, count)
	}

	return res
}
