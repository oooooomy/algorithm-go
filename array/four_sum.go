package array

import "slices"

// 18. 四数之和
func fourSum(nums []int, target int) (res [][]int) {
	slices.Sort(nums)
	for i := 0; i < len(nums)-3; i++ {
		if nums[i] > target && nums[i] > 0 && target > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if nums[i]+nums[j] > target && nums[i]+nums[j] > 0 && target > 0 {
				break
			}
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			x, y := nums[i], nums[j]
			l, r := j+1, len(nums)-1
			for l < r {
				if x+y+nums[l]+nums[r] < target {
					l++
				} else if x+y+nums[l]+nums[r] > target {
					r--
				} else {
					res = append(res, []int{x, y, nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				}
			}
		}
	}
	return
}
