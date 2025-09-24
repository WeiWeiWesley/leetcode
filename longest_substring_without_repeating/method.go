package lswrc

// Sliding Window - 滑動窗口演算法
// 核心思路：
// 1. 使用兩個指針 left 和 right 維護一個滑動窗口 [left, right]
// 2. 窗口內保證沒有重複字符
// 3. 當 right 遇到重複字符時，將 left 移動到重複字符的下一個位置
// 4. 持續記錄窗口的最大長度
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// charMap 記錄每個字符最後出現的位置
	charMap := make(map[byte]int)
	maxLen := 0
	left := 0 // 窗口左邊界

	for right := 0; right < n; right++ {
		// 如果當前字符已存在且位置在當前窗口內，需要移動左指針
		if pos, exists := charMap[s[right]]; exists && pos >= left {
			// 將 left 移動到重複字符的下一個位置
			// 這樣可以確保窗口內沒有重複字符
			left = pos + 1
		}

		// 更新當前字符的最新位置
		charMap[s[right]] = right

		// 計算當前窗口長度並更新最大值
		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}
