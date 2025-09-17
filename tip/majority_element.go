package tip

// 169. 多数元素
func majorityElement(nums []int) int {
	ans, vote := 0, 0
	for _, v := range nums {
		if vote == 0 {
			ans = v
		}
		if ans == v {
			vote++
		} else {
			vote--
		}
	}
	return ans
}
