func lengthOfLastWord(s string) int {
    var counter int

    for i := len(s) - 1; i >= 0; i-- {
        asciiChar := s[i]
        if counter != 0 && asciiChar == ' ' {
            break
        } else if asciiChar == ' ' {
            continue
        }

        counter++
    }
    return counter
}