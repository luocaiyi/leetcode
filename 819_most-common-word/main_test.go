package _19_most_common_word

import "testing"

func TestMostCommonWord(t *testing.T) {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	println(mostCommonWord(paragraph, banned))
}
