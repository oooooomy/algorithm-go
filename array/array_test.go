package array

import (
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(res)
}
