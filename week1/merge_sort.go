package week1

// working horse
func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	// split into smaller pieces
	firstPart, secondPart := a[0:len(a)/2], a[len(a)/2:]

	firstPart = MergeSort(firstPart)
	secondPart = MergeSort(secondPart)
	return merge(firstPart, secondPart)
}

// merge sorted arrays
func merge(firstPart, secondPart []int) []int {
	result := make([]int, len(firstPart)+len(secondPart))

	i, j := 0, 0
	for k := range result {
		// handle out of bound cases
		if i >= len(firstPart) {
			result[k] = secondPart[j]
			j++
			continue
		} else if j >= len(secondPart) {
			result[k] = firstPart[i]
			i++
			continue
		}

		// default condition
		if firstPart[i] > secondPart[j] {
			result[k] = secondPart[j]
			j++
		} else {
			result[k] = firstPart[i]
			i++
		}
	}
	return result
}
