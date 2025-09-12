package string

import "slices"

// 541. 反转字符串 II
func reverseStr(s string, k int) string {
	arr := []byte(s)
	n := len(s)
	for i := 0; i < n; i += 2 * k {
		slices.Reverse(arr[i:min(n, i+k)])
	}
	return string(arr)
}
