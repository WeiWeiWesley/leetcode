package lswrc

func lengthOfLongestSubstring(s string) int {

	m := make(map[byte]int)
	res := 0
	left, right := 0, 0

	for right < len(s) {
		rch := s[right]

		if count, ok := m[rch]; !ok || (ok && count == 0) {
			m[rch] = 1

			if right-left+1 > res {
				res = right - left + 1
			}

			// fmt.Println("right", right, "rch", string(rch))
			right++

			continue
		}

		lch := s[left]

		left++

		if count, ok := m[lch]; ok {
			m[lch] = count - 1
			// fmt.Println("left", left, "lch", string(lch))
		}
	}

	return res
}
