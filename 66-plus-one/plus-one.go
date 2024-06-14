func plusOne(digits []int) []int {
    digitsLength := len(digits)
    lastNumber := digits[digitsLength - 1]
    if lastNumber == 9 {
        for i := digitsLength - 1; i >= 0; i-- {
            currentNum := digits[i]
            if currentNum == 9 {
                digits[i] = 0

                if i == 0 {
                    digits = append([]int{1}, digits...)
                }
            } else {
                digits[i]++
                break;
            }
            
        }
    } else {
        digits[digitsLength - 1]++
    }

    return digits
}