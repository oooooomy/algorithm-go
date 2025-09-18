package array

// 704. 二分查找
// [L,R]
func search(nums []int, target int) int {
	left, right, mid := 0, len(nums), 0
	for left <= right {
		mid = (left + right) / 2
		if mid < 0 || mid >= len(nums) {
			break
		}
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// [L,R)
func search2(nums []int, target int) int {
	left, right, mid := 0, len(nums), 0
	for left < right {
		mid = (left + right) / 2
		if mid < 0 || mid >= len(nums) {
			break
		}
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return -1
}
