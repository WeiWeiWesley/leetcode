package findindexoffirstoccurrence

func strStr(haystack string, needle string) int {

	x := len(needle)

	if x > len(haystack) {
		return -1
	}

	for i := range haystack {

		if x+i > len(haystack) {
			return -1
		}

		if haystack[i:i+x] == needle {
			return i
		}

	}

	return -1
}
