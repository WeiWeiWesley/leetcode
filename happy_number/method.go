package happynumber

func isHappy(n int) bool {
	// 使用一個 set 來儲存出現過的數字
	seen := make(map[int]bool)

	// 不斷進行數字計算，直到得到 1 或者進入循環
	for n != 1 {
		// 如果數字 n 已經出現過，代表進入了循環，不是 Happy Number
		if seen[n] {
			return false
		}
		seen[n] = true

		// 計算新的數字，將 n 的每個位數平方後相加
		sum := 0
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}

		n = sum
	}

	// 最終得到 1，是 Happy Number
	return true
}
