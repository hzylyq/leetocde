package array

func ArrayOfProducts(array []int) []int {
	product := make([]int, len(array))

	for i := range array {
		runningProduct := 1
		for j := range array {
			if i != j {
				runningProduct *= array[j]
			}
		}
		product[i] = runningProduct
	}
	return product
}

func ArrayOfProduct2(array []int) []int {
	product := make([]int, len(array))
	leftProduct := make([]int, len(array))
	rightProduct := make([]int, len(array))

	for i := range array {
		product[i] = 1
		leftProduct[i] = 1
		rightProduct[i] = 1
	}
	leftRunningProduct := 1
	for i := range array {
		leftProduct[i] = leftRunningProduct
		leftRunningProduct *= array[i]
	}

	rightRunningProduct := 1
	for i := len(array) - 1; i >= 0; i-- {
		rightProduct[i] = rightRunningProduct
		rightRunningProduct *= array[i]
	}

	for i := range array {
		product[i] = leftProduct[i] * rightProduct[i]
	}
	return product
}

func ArrayOfProducts3(array []int) []int {
	product := make([]int, len(array))

	for i := range array {
		product[i] = 1
	}

	leftRunningProduct := 1
	for i := range array {
		product[i] = leftRunningProduct
		leftRunningProduct *= array[i]
	}

	rightRunningProduct := 1
	for i := len(array) - 1; i >= 0; i-- {
		product[i] *= rightRunningProduct
		rightRunningProduct *= array[i]
	}

	return product
}
