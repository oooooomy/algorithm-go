package string

import "testing"

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	if string(s) != "olleh" {
		t.Errorf("\nexpected: %s, output: %s\n", "olleh", s)
	}
}

func TestReverseStr(t *testing.T) {
	s := reverseStr("a", 2)
	t.Log(s)
}
