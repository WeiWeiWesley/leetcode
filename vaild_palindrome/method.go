package vaildpalindrome

import (
	"strings"
)

// isPalindrome 驗證字串是否為回文
// 解題思路：
// 1. 預處理：移除非字母數字字符，並轉換為小寫
// 2. 雙指標法：從兩端向中間比較字符
// 3. 時間複雜度：O(n)，空間複雜度：O(n)
func isPalindrome(s string) bool {
	// 邊界條件：空字串或單個字符都是回文
	if len(s) <= 1 {
		return true
	}

	// 步驟1：預處理 - 移除非字母數字字符並轉換為小寫
	// 這樣可以忽略大小寫和標點符號的差異
	newStr := ""
	for i := 0; i < len(s); i++ {
		if isAlphanumeric(s[i]) {
			newStr += strings.ToLower(string(s[i]))
		}
	}

	// 步驟2：雙指標法驗證回文
	// 從字串兩端開始，同時向中間移動並比較字符
	head := 0
	tail := len(newStr) - 1

	// 當兩個指標相遇時停止比較
	for head < tail {
		// 如果對應位置的字符不相等，則不是回文
		if newStr[head] != newStr[tail] {
			return false
		}
		// 移動指標向中間靠攏
		head++
		tail--
	}

	// 所有字符都匹配，是回文
	return true
}

// isAlphanumeric 檢查字符是否為字母或數字
// 使用ASCII值範圍判斷：
// - 'a'-'z': 小寫字母 (97-122)
// - 'A'-'Z': 大寫字母 (65-90)
// - '0'-'9': 數字 (48-57)
func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
