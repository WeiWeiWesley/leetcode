package palindrome

func isPalindrome(x int) bool {

	if x < 0 && int(x/10) <= 0 {
		return false
	}

	sl := []int{}

	for {
		a := x % 10

		x = int(x / 10)

		sl = append(sl, a)

		if x <= 0 {
			break
		}
	}

	lenX := len(sl)
	maxIndex := lenX - 1

	if lenX == 1 {
		return true
	}

	if lenX%2 == 0 {

		center := lenX / 2

		for i := 0; i < center; i++ {
			if sl[i] != sl[maxIndex-i] {
				return false
			}
		}
	} else {

		center := (lenX - 1) / 2

		for i := 0; i < center; i++ {
			if sl[i] != sl[maxIndex-i] {
				return false
			}
		}
	}

	return true
}
