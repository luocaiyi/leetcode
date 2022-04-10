package _0_valid_parentheses

import (
	"log"
	"testing"
)

var s1 = "()"
var s2 = "()[]{}"
var s3 = "{[]}"

var es1 = "(]"
var es2 = "([)]"

func TestIsValid(t *testing.T) {
	if !isValid(s1) || !isValid(s2) || !isValid(s3) {
		log.Fatal("Err")
	}
	if isValid(es1) || isValid(es2) {
		log.Fatal("Err2")
	}
}
