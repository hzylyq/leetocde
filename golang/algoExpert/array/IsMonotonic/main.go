package main

func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}

	direction := array[1] - array[0]
	for i := 2; i < len(array); i++ {
		if direction == 0 {
			direction = array[i] - array[i-1]
			continue
		}

		difference := array[i] - array[i-1]
		if !isSameDirection(direction, difference) {
			return false
		}
	}

	return true
}

func isSameDirection(direction, difference int) bool {
	if direction > 0 {
		return difference >= 0
	} else {
		return difference <= 0
	}
}
