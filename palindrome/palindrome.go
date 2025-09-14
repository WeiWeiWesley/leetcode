package palindrome


// isPalindrome 判斷一個整數是否為回文數
// 使用反轉數字的一半來避免整數溢位問題
// 時間複雜度：O(log n)，空間複雜度：O(1)
func isPalindrome(x int) bool {
	// 邊界條件檢查：
	// 1. 負數不可能是回文數（如 -121 反轉後變成 121-）
	// 2. 以 0 結尾的數字（除了 0 本身）不可能是回文數
	//    因為回文數不能以 0 開頭（如 10 反轉後變成 01）
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 初始化反轉後的數字
	revertedNumber := 0

	// 核心演算法：只反轉數字的一半
	// 當 x > revertedNumber 時繼續反轉，確保只處理一半的位數
	// 這樣可以避免整數溢位問題，因為我們不需要完全反轉整個數字
	for x > revertedNumber {
		// 將 x 的最後一位數字加到 revertedNumber 的末尾
		// 例如：x=123, revertedNumber=0 -> revertedNumber=3
		//      x=12,  revertedNumber=3  -> revertedNumber=32
		revertedNumber = revertedNumber*10 + x%10

		// 移除 x 的最後一位數字
		// 例如：x=123 -> x=12
		x /= 10
	}

	// 比較結果：
	// 1. 偶數位數的數字：x == revertedNumber
	//    例如：1221 -> x=12, revertedNumber=12 -> true
	// 2. 奇數位數的數字：x == revertedNumber/10
	//    例如：12321 -> x=12, revertedNumber=123 -> x == revertedNumber/10 (12) -> true
	//    中間的數字不需要比較，因為它會與自己相等
	return x == revertedNumber || x == revertedNumber/10
}

