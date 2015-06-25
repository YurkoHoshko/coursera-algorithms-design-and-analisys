package week1

func FindInversions(arr []int) ([]int, map[int][]int) {
	allInversions := make(map[int][]int)
	if len(arr) <= 1 {
		return arr, allInversions
	}

	// split into smaller pieces
	leftPart, rightPart := arr[0:len(arr)/2], arr[len(arr)/2:]

	leftPart, leftInversions := FindInversions(leftPart)
	rightPart, rightInversions := FindInversions(rightPart)
	wholeArray, splitInversions := sortAndFindSplitInversions(leftPart, rightPart)

	// merge all inversions into one map
	for _, elem := range arr {
		allInversions[elem] = append(leftInversions[elem], rightInversions[elem]...)
		allInversions[elem] = append(allInversions[elem], splitInversions[elem]...)
	}
	return wholeArray, allInversions
}

func sortAndFindSplitInversions(leftPart, rightPart []int) ([]int, map[int][]int) {
	resultArray := make([]int, len(leftPart)+len(rightPart))
	inversions := make(map[int][]int)

	i, j := 0, 0
	for k := range resultArray {
		// handle out of bound cases
		if i >= len(leftPart) {
			resultArray[k] = rightPart[j]
			j++
			continue
		} else if j >= len(rightPart) {
			resultArray[k] = leftPart[i]
			_, ok := inversions[leftPart[i]]
			// add inversions only if they were not added yet
			if !ok {
				inversions[leftPart[i]] = append(inversions[leftPart[i]], rightPart...)
			}
			i++
			continue
		}

		// default condition
		if leftPart[i] > rightPart[j] {
			resultArray[k] = rightPart[j]
			inversions[leftPart[i]] = append(inversions[leftPart[i]], rightPart[j])
			j++
		} else {
			resultArray[k] = leftPart[i]
			i++
		}
	}
	return resultArray, inversions
}
