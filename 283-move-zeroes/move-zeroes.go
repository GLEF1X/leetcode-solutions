func moveZeroes(nums []int) {
	zeroPointer, numPointer := -1, -1
	for i := 0; i < len(nums); i++ {
		number := nums[i]
		if number == 0 && zeroPointer == -1 {
			zeroPointer = i
		} else if number != 0 && numPointer == -1 {
			numPointer = i
		}

		if zeroPointer != -1 && numPointer != -1 {
			if zeroPointer > numPointer {
				numPointer = -1
				continue
			}

			nums[zeroPointer] = nums[numPointer]
			nums[numPointer] = 0
			zeroPointer++
			numPointer = -1
		}
	}
}
