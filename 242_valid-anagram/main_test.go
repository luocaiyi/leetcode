package _42_valid_anagram

import (
	"log"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s1, t1 := "anagram", "nagaram"
	s2, t2 := "rat", "car"
	if !isAnagram(s1, t1) || isAnagram(s2, t2) {
		log.Fatal("Err")
	}
}
