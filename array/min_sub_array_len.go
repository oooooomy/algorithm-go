package array

// 209. 长度最小的子数组
// 标准的零分
// func minSubArrayLenError(target int, nums []int) int {
// 	slices.Sort(nums)
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
	return 0
}
