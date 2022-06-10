package lengthoflastword

func lengthOfLastWord(s string) int {

	if len(s) == 0 {
		return 0
	}

	count := 0

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == ' ' {
			if count != 0 {
				return count
			}
		} else {
			count++
		}

	}

	return count
}
