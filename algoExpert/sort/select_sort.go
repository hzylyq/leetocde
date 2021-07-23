package sort

func SelectionSort(arr []int) []int {
	for i := range arr {
		smallest := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[smallest] {
				smallest = j
			}
		}
		if arr[i] != arr[smallest] {
			arr[i], arr[smallest] = arr[smallest], arr[i]
		}
	}

	return arr
}
