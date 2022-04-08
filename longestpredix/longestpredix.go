package longestpredix

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	longest := ""

	for i := 0; i <= len(strs[0]); i++ {

		tmp := strs[0][0:i]
		check := 0

		for _, v := range strs {
			if i > len(v) {
				break
			}

			if tmp == v[0:i] {
				check++
			}
		}

		if check == len(strs) {
			longest = tmp
		}
	}

	return longest
}
