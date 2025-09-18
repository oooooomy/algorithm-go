package array

import (
	"math"
)

// 209. 长度最小的子数组
// 标准的零分
// func minSubArrayLenError(target int, nums []int) int {
// 	slices.Sort(nums) //要求子数组，排序就不是了
// 	if nums[0] > target {
// 		return 0
// 	}
// 	sum, res := 0, 0
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		sum += nums[i]
// 		res++
// 		if sum == target {
// 			return res
// 		}
// 	}
// 	return 0
// }

func minSubArrayLen(target int, nums []int) int {
	slow, sum, ans := 0, 0, math.MaxInt
	for fast, v := range nums {
		sum += v
		for sum >= target {
			ans = min(ans, fast-slow+1)
			sum -= nums[slow]
			slow++
		}
	}
	if ans != math.MaxInt {
		return ans
	}
	return 0
}
