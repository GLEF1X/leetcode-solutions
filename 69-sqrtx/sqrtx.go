func mySqrt(x int) int {
	left, right := 0, x

	var res int
	for left <= right {
		possibleSquareRootResult := left + ((right - left) / 2)
		currentNumberSampledSquared := possibleSquareRootResult * possibleSquareRootResult

		if currentNumberSampledSquared > x {
			right = possibleSquareRootResult - 1
		} else if currentNumberSampledSquared < x {
			left = possibleSquareRootResult + 1
			res = possibleSquareRootResult
		} else {
			return possibleSquareRootResult
		}
	}

	return res
}