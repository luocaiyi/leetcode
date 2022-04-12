package _9_group_anagrams

func groupAnagrams(strs []string) (res [][]string) {
	if len(strs) < 2 {
		return [][]string{strs}
	}
	m := make(map[[26]uint8][]string)
	for _, str := range strs {
		var ch [26]uint8
		for _, c := range str {
			ch[c-'a']++
		}
		m[ch] = append(m[ch], str)
	}
	for _, strings := range m {
		res = append(res, strings)
	}
	return
}
