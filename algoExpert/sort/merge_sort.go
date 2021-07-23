package sort

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	midIdx := len(arr) / 2
	leftArr := MergeSort(arr[:midIdx])
	rightArr := MergeSort(arr[midIdx:])
	return MergeSortArr(leftArr, rightArr)
}

func MergeSortArr(left, right []int) []int {
	res := make([]int, len(left)+len(right))

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		res[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		res[k] = right[j]
		j++
		k++
	}

	return res
}
