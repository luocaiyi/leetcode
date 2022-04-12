package _9_group_anagrams

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	res := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Printf("%v", res)
}
