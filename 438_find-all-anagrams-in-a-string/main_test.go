package _38_find_all_anagrams_in_a_string

import (
	"fmt"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	res := findAnagrams(s, p)
	fmt.Printf("%v", res)
}
