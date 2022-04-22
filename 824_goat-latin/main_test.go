package _24_goat_latin

import "testing"

func TestToGoatLatin(t *testing.T) {
	sentence := "The quick brown fox jumped over the lazy dog"
	println(toGoatLatin(sentence))
}
