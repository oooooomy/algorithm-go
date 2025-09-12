package string

// 344. 反转字符串
func reverseString(s []byte) {
	n := len(s)
	if n < 2 {
		return
	}
	l, r := 0, n-1
	for l < r {
		t := s[l]
		s[l] = s[r]
		s[r] = t
		l++
		r--
	}
}
