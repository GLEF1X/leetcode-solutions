func longestCommonPrefix(strs []string) string {
	var iterationChar byte
	var currentIdx int
	for {
		for i, str := range strs {
			if currentIdx >= len(str) {
				return str[:currentIdx]
			}
			if i == 0 {
				iterationChar = str[currentIdx]
			}

			char := str[currentIdx]
			if iterationChar != 0 && char != iterationChar {
				return str[:currentIdx]
			}

		}

		currentIdx++
	}

	return ""
}