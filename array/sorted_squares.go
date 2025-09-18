package array

// 977. 有序数组的平方
// 小于0先现排序，绝对值的复数不行
// 平方最大的元素一定是在数组的两侧!!!
// 两次双指针平方和比较插入新数组尾部
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	//这里容易忽略'='号,优化做法
	// l, r, p := 0, n-1, n-1
	// for l <= r {
	// 	if (nums[l] * nums[l]) < (nums[r] * nums[r]) {
	// 		res[p] = nums[r] * nums[r]
	// 		r--
	// 	} else {
	// 		res[p] = nums[l] * nums[l]
	// 		l++
	// 	}
	// 	p--
	// }

	l, r := 0, n-1
	for i := n - 1; i >= 0; i-- {
		x, y := nums[l]*nums[l], nums[r]*nums[r]
		if x < y {
			res[i] = y
			r--
		} else {
			res[i] = x
			l++
		}
	}

	return res
}
