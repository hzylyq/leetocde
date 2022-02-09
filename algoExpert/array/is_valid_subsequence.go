package array

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	var seqIndex, arrIndex int
	for arrIndex < len(array) && seqIndex < len(sequence) {
		if array[arrIndex] == sequence[seqIndex] {
			seqIndex++
			arrIndex++
		} else {
			arrIndex++
		}
	}

	return seqIndex == len(sequence)
}

func IsValidSubsequence2(array []int, sequence []int) bool {
	// Write your code here.
	var seqIndex int
	for _, val := range array {
		if seqIndex == len(sequence) {
			break
		}
		if val == sequence[seqIndex] {
			seqIndex++
		}
	}

	return seqIndex == len(sequence)
}
