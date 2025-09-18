package tip

import "testing"

func TestMajorityElement(t *testing.T) {
	ans := majorityElement([]int{3, 2, 3})
	if ans != 3 {
		t.Errorf("\nexpected: %d, output: %d\n", 3, ans)
	}
	ans = majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	if ans != 2 {
		t.Errorf("\nexpected: %d, output: %d\n", 2, ans)
	}
	ans = majorityElement([]int{0, 0, 5, 3, 5, 1, 5, 7})
	t.Logf("test ans: %d\n", ans)
}
