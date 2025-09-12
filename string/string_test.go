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
