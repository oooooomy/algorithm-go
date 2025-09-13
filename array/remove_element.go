package array

// 27. 移除元素
func removeElement(nums []int, val int) int {
	k := 0
	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}
	return k
}

func removeElementTwo(nums []int, val int) int {
	fast, slow := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
