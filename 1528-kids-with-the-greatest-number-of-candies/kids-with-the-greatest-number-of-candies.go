func kidsWithCandies(candies []int, extraCandies int) []bool {
    result := make([]bool, 0, len(candies))
    for i := 0; i < len(candies); i++ {
        candiesCountWithExtra := candies[i] + extraCandies
        var isGreatest bool = true
        for j := 0; j < len(candies); j++ {
            if j == i {
                continue
            }

            if candiesCountWithExtra < candies[j] {
                isGreatest = false
            }
        }

        result = append(result, isGreatest)
    }

    return result
}