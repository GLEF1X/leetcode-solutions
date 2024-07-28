import "strings"

func reverseVowels(s string) string {
	// leetcode
	// [e e o e]
	// [e o e e]
	start, end := 0, len(s)-1
	var startReadyForSwap, endReadyForSwap bool
	// var sb strings.Builder
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
			charFromStart := strings.ToLower(string(b[start]))
			if charFromStart == "a" || charFromStart == "e" || charFromStart == "i" || charFromStart == "o" || charFromStart == "u" {
				startReadyForSwap = true
			} else {
				start++
			}
			continue
		}

		if !endReadyForSwap {
			charFromEnd := strings.ToLower(string(b[end]))
			if charFromEnd == "a" || charFromEnd == "e" || charFromEnd == "i" || charFromEnd == "o" || charFromEnd == "u" {
				endReadyForSwap = true
			} else {
				end--
			}
			continue
		}

	}
	return string(b)
}
