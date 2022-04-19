package _21_shortest_distance_to_a_character

import (
	"fmt"
	"testing"
)

func TestShortestToChar(t *testing.T) {
	s := "loveleetcode"
	c := byte('e')
	fmt.Printf("%+v\n", shortestToChar(s, c))
}
