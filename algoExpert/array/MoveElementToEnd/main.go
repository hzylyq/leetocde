package main

func MoveElementToEnd(array []int, toMove int) []int {
	i, j := 0, len(array)-1
	for i < j {
		for i < j && array[j] == toMove {
			j--
		}
		if array[i] == toMove {
			array[i], array[j] = array[j], array[i]
		}
		i++
	}
	return array
}
