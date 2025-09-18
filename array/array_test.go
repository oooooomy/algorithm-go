package array

import (
	"testing"
)

func TestFourSum(t *testing.T) {
	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	t.Log(res)
}

func TestRemoveElement(t *testing.T) {
	res := removeElement([]int{3, 2, 2, 3}, 3)
	t.Log("res: ", res)
	if res != 2 {
		t.Fail()
	}
	res = removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	t.Log("res: ", res)
	if res != 5 {
		t.Fail()
	}
	res = removeElementTwo([]int{3, 2, 2, 3}, 3)
	t.Log("res: ", res)
	if res != 2 {
		t.Fail()
	}
	res = removeElementTwo([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	t.Log("res: ", res)
	if res != 5 {
		t.Fail()
	}
}

func TestSortedSquares(t *testing.T) {
	res := sortedSquares([]int{-7, -3, 2, 3, 11})
	t.Log("res: ", res)
	a := []int{4, 9, 9, 49, 121}
	for i, v := range res {
		if v != a[i] {
			t.Fail()
		}
	}
}

func TestMinSubArrayLen(t *testing.T) {
	res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	t.Log("res: ", res)
	if res != 2 {
		t.Errorf("\nexpected: %d, output: %d\n", 2, res)
	}
	res = minSubArrayLen(11, []int{1, 2, 3, 4, 5})
	t.Log("res: ", res)
	if res != 3 {
		t.Errorf("\nexpected: %d, output: %d\n", 3, res)
	}
}

func TestSearch(t *testing.T) {
	ans := 0
	ans = search([]int{-1, 0, 3, 5, 9, 12}, 9)
	if ans != 4 {
		t.Errorf("\nexpected: %d, output: %d\n", 4, ans)
	}
	ans = search([]int{-1, 0, 3, 5, 9, 12}, 2)
	if ans != -1 {
		t.Errorf("\nexpected: %d, output: %d\n", -1, ans)
	}
	ans = search2([]int{-1, 0, 3, 5, 9, 12}, 9)
	if ans != 4 {
		t.Errorf("\nexpected: %d, output: %d\n", 4, ans)
	}
	ans = search2([]int{-1, 0, 3, 5, 9, 12}, 2)
	if ans != -1 {
		t.Errorf("\nexpected: %d, output: %d\n", -1, ans)
	}
}
