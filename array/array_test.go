package array

import (
	"fmt"
	"log"
	"testing"
)

func TestFourSum(t *testing.T) {
	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(res)
}

func TestRemoveElement(t *testing.T) {
	res := removeElement([]int{3, 2, 2, 3}, 3)
	log.Println("res: ", res)
	if res != 2 {
		t.Fail()
	}
	res = removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	log.Println("res: ", res)
	if res != 5 {
		t.Fail()
	}
	res = removeElementTwo([]int{3, 2, 2, 3}, 3)
	log.Println("res: ", res)
	if res != 2 {
		t.Fail()
	}
	res = removeElementTwo([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	log.Println("res: ", res)
	if res != 5 {
		t.Fail()
	}
}
