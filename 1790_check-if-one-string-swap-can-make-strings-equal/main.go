package _790_check_if_one_string_swap_can_make_strings_equal

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	cnt := 0
	for i, diff := 0, 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			switch cnt {
			case 0:
				diff = i
			case 1:
				if s1[diff] != s2[i] || s2[diff] != s1[i] {
					return false
				}
			default:
				return false
			}
			cnt++
		}
	}
	if cnt == 2 {
		return true
	}
	return false
}
