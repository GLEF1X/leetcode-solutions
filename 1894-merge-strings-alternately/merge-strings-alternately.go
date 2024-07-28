import "strings"

func mergeAlternately(word1 string, word2 string) string {
    var sb strings.Builder

    var ptr1, ptr2 int
    var word1Turn bool = true
    for ptr1 < len(word1) || ptr2 < len(word2) {
        if word1Turn {
            if ptr1 < len(word1) {
                sb.WriteByte(word1[ptr1])
                ptr1++
            }
            word1Turn = false
        } else {
            if ptr2 < len(word2) {
                sb.WriteByte(word2[ptr2])
                ptr2++
            }
            word1Turn = true
        }
    }

    return sb.String()
}