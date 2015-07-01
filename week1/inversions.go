package week1

func FindInversions(arr []int) ([]int, int) {
	allInversions := 0
	if len(arr) <= 1 {
		return arr, allInversions
	}

	// split into smaller pieces
	leftPart, rightPart := arr[0:len(arr)/2], arr[len(arr)/2:]

	leftPart, leftInversions := FindInversions(leftPart)
	rightPart, rightInversions := FindInversions(rightPart)
	wholeArray, splitInversions := sortAndFindSplitInversions(leftPart, rightPart)

	// merge all inversions into one map
	allInversions = leftInversions + rightInversions + splitInversions

	return wholeArray, allInversions
}

func sortAndFindSplitInversions(leftPart, rightPart []int) ([]int, int) {
	resultArray := make([]int, len(leftPart)+len(rightPart))
	inversions := 0

	i, j := 0, 0
	for k := range resultArray {
		// handle out of bound cases
		if i >= len(leftPart) {
			resultArray[k] = rightPart[j]
			j++
			continue
		} else if j >= len(rightPart) {
			resultArray[k] = leftPart[i]
			i++
			continue
		}

		// default condition
		if leftPart[i] > rightPart[j] {
			resultArray[k] = rightPart[j]
			inversions += len(leftPart) - i
			j++
		} else {
			resultArray[k] = leftPart[i]
			i++
		}
	}
	return resultArray, inversions
}
