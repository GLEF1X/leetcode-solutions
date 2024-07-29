function compress(chars) {
    if (chars.length === 1) {
        return 1
    }

    let windowStartPointer = 0, windowEndPointer = 0, currentChar = chars[0], result = 1;

    // ["a","a","b","b","c"]
    while (windowEndPointer < chars.length) {
        const char = chars[windowEndPointer]
        if (char === currentChar) {
            windowEndPointer++
        }
        if (char !== currentChar || windowEndPointer === chars.length) {
            const charCount = windowEndPointer - windowStartPointer
            if (charCount > 1) {
                if (charCount >= 10) {
                    let digits = [], num = charCount
                    while (num !== 0) {
                        const currentDigit = num % 10
                        digits.unshift(String(currentDigit))
                        num = Math.floor(num / 10)
                    }

                    chars.splice(windowStartPointer + 1, charCount - 1, ...digits)
                    result += 1 + digits.length
                    windowStartPointer += digits.length + 1
                    windowEndPointer = windowStartPointer
                } else {
                    result += 2
                    chars.splice(windowStartPointer + 1, charCount - 1, String(charCount))
                    windowStartPointer += 2
                    windowEndPointer = windowStartPointer
                }
            } else {
                result += 1

                windowStartPointer = windowEndPointer
                windowEndPointer++
            }
        }
        currentChar = char
    }
    return result
};