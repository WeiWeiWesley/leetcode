package longestpalindromicsubstring

// 方法一：暴力搜尋法 (Brute Force) - O(n³)
// 雙層迴圈遍歷所有可能的子字串，然後檢查是否為回文
func longestPalindromeBruteForce(s string) string {
	n := len(s)

	switch n {
	case 0, 1:
		return s
	}

	res := ""

	// 雙層迴圈，左右移動指標
	for i := 0; i < n; i++ {
		for j := n - 1; i <= j; j-- {

			// 檢查左右一致
			if s[i] != s[j] {
				continue
			}

			// 檢查最大回文
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

// 方法二：中心擴展法 (Expand Around Centers) - O(n²)
// 以每個字符（或字符間隙）為中心，向兩邊擴展尋找回文
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return  ""
	}

	res := ""


	for i:=0; i<len(s); i++ {

		currMaxlen := 0

		// odd
		lenO := expend(s, i, i)
		// even
		lenE := expend(s, i, i+1)

		if lenE > lenO {
			currMaxlen = lenE
		} else {
			currMaxlen = lenO
		}

		if currMaxlen > len(res) {
			start := i - (currMaxlen-1)/2
			res = s[start:start+currMaxlen]
		}
	}

	return res
}

func expend(s string, left, right int) int {
	for left>=0 && right<len(s) && s[left] == s[right] {
		left--
		right++
	}

	// 左右邊界各超過 1 = -2
	// 長度計算 + 1
	return right - left - 2 + 1
}


func isPalindrome(s string, left, right int) bool {
	// 左右同時向內縮
	for i, j := left, right; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
