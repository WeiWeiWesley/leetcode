package longestpalindromicsubstring

// a
// aa
// babad -> bab
func longestPalindrome(s string) string {

	n := len(s)

	switch n {
	case 0, 1:
		return s
	}

	res := ""

	//雙層迴圈，左右移動指標
	for i := 0; i < n; i++ {
		for j := n - 1; i <= j; j-- {

			//檢查左右一致
			if s[i] != s[j] {
				continue
			}

			//檢查最大回文
			if isPalindrome(s, i, j) {

				if len(s[i:j+1]) > len(res) {
					res = s[i : j+1]
				}

				break
			}
		}
	}

	return res
}

func isPalindrome(s string, left, right int) bool {

	//左右同時向內縮
	for i, j := left, right; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
