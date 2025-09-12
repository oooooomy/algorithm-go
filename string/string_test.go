package string

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s)
}

func TestReverseStr(t *testing.T) {
	s := reverseStr("a", 2)
	fmt.Println(s)
}
