package _42_valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	tab := map[rune]int{}
//	for _, ch := range s {
//		tab[ch]++
//	}
//	for _, ch := range t {
//		tab[ch]--
//		if tab[ch] < 0 {
//			return false
//		}
//	}
//	return true
//}

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	s1, s2 := []byte(s), []byte(t)
//	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j]} )
//	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j]} )
//	return string(s1) == string(s2)
//}
