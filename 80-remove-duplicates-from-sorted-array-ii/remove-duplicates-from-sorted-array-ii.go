func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // This index will point to the position where the next unique element should be placed
    insertionIdx := 1
    count := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            count++
        } else {
            count = 1
        }

        if count <= 2 {
            nums[insertionIdx] = nums[i]
            insertionIdx++
        }
    }

    return insertionIdx
}