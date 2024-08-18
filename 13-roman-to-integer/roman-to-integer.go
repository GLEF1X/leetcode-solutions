func romanToInt(s string) int {
    m := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    sum := 0
    var prev rune
    for _, char := range s {
        valueOfAChar := m[char]
        if prev == 'I' && (char == 'V' || char == 'X') {
            valueOfAChar--
            valueOfAChar -= m[prev]
        } else if prev == 'X' && (char == 'L' || char == 'C') {
            valueOfAChar -= 10
            valueOfAChar -= m[prev]
        } else if prev == 'C' && (char == 'D' || char == 'M') {
            valueOfAChar -= 100
            valueOfAChar -= m[prev]
        }

        sum += valueOfAChar
        prev = char
    }

    return sum
}