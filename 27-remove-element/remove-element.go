func removeElement(nums []int, val int) int {
    var removedIdx int
    var swapping bool
    var occurencesCount int

    for i, num := range nums {
        if num == val {
            if !swapping {
                removedIdx = i
                swapping = true
            }

            continue
        } else {
            occurencesCount++
        }

        if swapping {
            nums[removedIdx] = num
            removedIdx++
        }
    }

    return occurencesCount
}