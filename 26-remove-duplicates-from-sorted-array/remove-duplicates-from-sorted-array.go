func removeDuplicates(nums []int) int {
	var uniqueCount, prevNumber int
	numsLength, insertionIdx := len(nums), -1

	for i, num := range nums {
		if num != prevNumber || i == 0 {
			uniqueCount++
		}

		if newNumber := num != prevNumber; newNumber && insertionIdx != -1 {
			nums[insertionIdx] = num
			insertionIdx++
		}

		if i+1 < numsLength && nums[i+1] == num && insertionIdx == -1 {
			insertionIdx = i + 1
		}

		prevNumber = num
	}

	return uniqueCount
}