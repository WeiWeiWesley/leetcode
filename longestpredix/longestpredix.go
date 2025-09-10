package longestpredix

func longestCommonPrefix(strs []string) string {
	// 邊界條件：如果輸入為空，直接返回空字串
	if len(strs) == 0 {
		return ""
	}

	// 以第一個字串作為初始前綴（最長可能的前綴）
	prefix := strs[0]

	// 從最長的前綴開始，逐步縮短
	for len(prefix) > 0 {
		allMatch := true

		// 檢查所有字串是否都有這個前綴
		for _, str := range strs {
			// 如果字串長度不足，或者前綴不匹配
			if len(str) < len(prefix) || str[:len(prefix)] != prefix {
				allMatch = false
				break // 一旦發現不匹配，立即跳出
			}
		}

		// 如果所有字串都匹配這個前綴，返回結果
		if allMatch {
			return prefix
		}

		// 縮短前綴：移除最後一個字符
		prefix = prefix[:len(prefix)-1]
	}

	// 如果沒有找到任何共同前綴，返回空字串
	return ""
}
