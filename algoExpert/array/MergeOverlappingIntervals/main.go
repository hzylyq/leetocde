package MergeOverlappingIntervals

import "sort"

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	// first sort interval
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedInterval := make([][]int, 0)
	currentInterval := intervals[0]
	mergedInterval = append(mergedInterval, currentInterval)

	for _, nextInterval := range intervals {
		currentIntervalEnd := currentInterval[1]
		nextIntervalStart, nextIntervalEnd := nextInterval[0], nextInterval[1]
		if currentIntervalEnd >= nextIntervalStart {
			currentInterval[1] = max(currentIntervalEnd, nextIntervalEnd)
		} else {
			currentInterval = nextInterval
			mergedInterval = append(mergedInterval, currentInterval)
		}
	}

	return mergedInterval
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
