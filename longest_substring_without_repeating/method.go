package lswrc

// Sliding Window
func lengthOfLongestSubstring(s string) int {

	n := len(s)
	result := 0
	start := 0

	charIndexMap := make(map[uint8]int)

    //起點終點都由最左動態至最右
	for end := 0; end < n; end++ {

		// 檢查是否已出現重複
		duplicateIndex, isDuplicate := charIndexMap[s[end]]

		if isDuplicate {
			// 更新當前結果
			result = max(result, end-start)

			// 清除已出現的部分
			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}

			// 移動起點至我們發現重複的下一個字元
			start = duplicateIndex + 1
		}

		// 記錄不重複結果
		charIndexMap[s[end]] = end
	}
	// Update the result for last window
	// For a input like "abc" the execution will not enter the above if statement for checking duplicates
	result = max(result, n-start)

	return result
}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}