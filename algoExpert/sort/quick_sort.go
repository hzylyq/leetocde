package sort


func QuickSort(array []int) []int {
	// Write your code here.
	return helper(array, 0, len(array)-1)
}

func helper(array []int, start, end int) []int {
	if start >= end {
		return array
	}

	pivot := start
	left := start + 1
	right := end

	for right >= left {
		if array[left] > array[pivot] && array[right] < array[pivot] {
			array[left], array[right] = array[right], array[left]
		}
		if array[left] <= array[pivot] {
			left += 1
		}
		if array[right] >= array[pivot] {
			right -= 1
		}
	}

	array[pivot], array[right] = array[right], array[pivot]
	if right-1-start < end-(right+1) {
		helper(array, start, right-1)
		helper(array, right+1, end)
	} else {
		helper(array, right+1, end)
		helper(array, start, right-1)
	}

	return array
}

