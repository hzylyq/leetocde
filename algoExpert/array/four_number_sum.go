package array

func FourNumberSum(array []int, target int) [][]int {
	allPairSum := make(map[int][][]int)
	quadruplets := make([][]int, 0)

	for i := 1; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			currentSum := array[i] + array[j]
			difference := target - currentSum
			if pairs, ok := allPairSum[difference]; ok {
				for _, pair := range pairs {
					newQuadruplet := append(pair, array[i], array[j])
					quadruplets = append(quadruplets, newQuadruplet)
				}
			}
		}

		for k := 0; k < i; k++ {
			currSum := array[i] + array[k]
			allPairSum[currSum] = append(allPairSum[currSum], []int{array[k], array[i]})
		}
	}

	return quadruplets
}
