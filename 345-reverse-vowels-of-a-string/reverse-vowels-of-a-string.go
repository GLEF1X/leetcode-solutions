import "strings"

func reverseVowels(s string) string {
	start, end := 0, len(s)-1
	var startReadyForSwap, endReadyForSwap bool
	b := []byte(s)

	for start < end {
		if startReadyForSwap && endReadyForSwap {
			t := b[start]
			b[start] = b[end]
			b[end] = t
			startReadyForSwap, endReadyForSwap = false, false
			start++
			end--
			continue
		}

		if !startReadyForSwap {
			if isVowel(b[start]){
				startReadyForSwap = true
			} else {
				start++
			}
			continue
		}

		if !endReadyForSwap {
			if isVowel(b[end]) {
				endReadyForSwap = true
			} else {
				end--
			}
			continue
		}

	}
	return string(b)
}


func isVowel(b byte) bool {
     char := strings.ToLower(string(b))
     return char == "a" || char == "e" || char == "i" || char == "o" || char == "u"
}