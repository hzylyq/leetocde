package main

import "sort"

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)

	res := 0
	for idx, query := range queries {
		res += query *(len(queries)-idx-1)
	}
	return res
}
