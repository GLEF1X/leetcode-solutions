func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums) - 1

    for left <= right {
        // remember to rewrite into left + ((right - left) / 2) to prevent overflow
        mid := (left + right) / 2
        currentNum := nums[mid]

        if currentNum == target {
            return mid   
        } else if currentNum < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return left
}