package _85_mini_parser

import (
	"testing"
)

func TestDeserialize(t *testing.T) {
	deserialize("[123,[456,[789]]]")
	deserialize("234")
}
