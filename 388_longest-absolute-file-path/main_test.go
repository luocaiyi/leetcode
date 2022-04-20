package _88_longest_absolute_file_path

import "testing"

func TestLengthLongestPath(t *testing.T) {
	input := "dir\\n\\tsubdir1\\n\\t\\tfile1.ext\\n\\t\\tsubsubdir1\\n\\tsubdir2\\n\\t\\tsubsubdir2\\n\\t\\t\\tfile2.ext"
	println(lengthLongestPath(input))
}
