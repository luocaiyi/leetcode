package _24_goat_latin

import "strings"

func toGoatLatin(sentence string) string {
	vowel := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	ans := &strings.Builder{}
	for i, cnt, n := 0, 1, len(sentence); i < n; i++ {
		if cnt > 1 {
			ans.WriteByte(' ')
		}
		start := i
		for i++; i < n && sentence[i] != ' '; i++ {
		}
		cnt++
		if vowel[sentence[start]] {
			ans.WriteString(sentence[start:i])
		} else {
			ans.WriteString(sentence[start+1 : i])
			ans.WriteByte(sentence[start])
		}
		ans.WriteByte('m')
		ans.WriteString(strings.Repeat("a", cnt))
	}
	return ans.String()
}
