package tries

func MultiStringSearch(bigString string, smallStrings []string) []bool {
	outPut := make([]bool, 0, len(smallStrings))

	for _, smallStr := range smallStrings {
		outPut = append(outPut, isInBigString(bigString, smallStr))
	}

	return outPut
}

func isInBigString(bigString, smallString string) bool {
	for i := range bigString {
		if i+len(smallString) > len(bigString) {
			break
		}
		if isInBigStringHelper(bigString, smallString, i) {
			return true
		}
	}

	return false
}

func isInBigStringHelper(bigStr, smallStr string, idx int) bool {
	leftBigIdx := idx
	rightBigIdx := idx + len(smallStr) - 1
	leftSmallIdx := 0
	rightSmallIdx := len(smallStr) - 1

	for leftBigIdx <= rightBigIdx {
		if bigStr[leftBigIdx] != smallStr[leftSmallIdx] || bigStr[rightBigIdx] != smallStr[rightSmallIdx] {
			return false
		}

		leftBigIdx += 1
		rightBigIdx -= 1
		leftSmallIdx += 1
		rightSmallIdx -= 1
	}

	return true
}
