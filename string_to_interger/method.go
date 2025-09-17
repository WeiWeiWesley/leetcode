package stringtointerger

// myAtoi 將字串轉換為 32 位有符號整數
// 實現類似 C/C++ 的 atoi 函數功能
// 演算法步驟：
// 1. 忽略前導空白字符
// 2. 檢查正負號（如果有的話）
// 3. 讀取數字直到遇到非數字字符
// 4. 將數字轉換為整數並處理溢出
func myAtoi(s string) int {
	res := 0          // 結果值
	negative := false // 是否為負數
	count := 0        // 已讀取的數字個數

	// 32位有符號整數的範圍限制
	const (
		MaxInt32 = 2147483647  // 2^31 - 1
		MinInt32 = -2147483648 // -2^31
	)

	// 遍歷字串中的每個字符
	for i := 0; i < len(s); i++ {
		// 處理空白字符
		if s[i] == ' ' {
			// 如果已經開始讀取數字，遇到空格就停止
			if count > 0 {
				break
			}
			continue // 跳過前導空白
		}

		// 處理負號（只能在開頭且未讀取數字時）
		if s[i] == '-' && count == 0 {
			negative = true
			continue
		}

		// 處理正號（只能在開頭且未讀取數字時）
		if s[i] == '+' && count == 0 {
			continue
		}

		// 嘗試將字符轉換為數字
		num, ok := isNum(s[i])
		if ok {
			// 檢查是否會導致 32 位整數溢出
			// 條件：res > MaxInt32/10 或 (res == MaxInt32/10 且 num > MaxInt32%10)
			if res > MaxInt32/10 || (res == MaxInt32/10 && num > MaxInt32%10) {
				if negative {
					return MinInt32 // 負數溢出，返回最小值
				}
				return MaxInt32 // 正數溢出，返回最大值
			}
			// 安全地添加數字
			res = res*10 + num
			count++
		} else {
			// 遇到非數字字符，停止讀取
			break
		}
	}

	// 根據正負號調整結果
	if negative {
		res *= -1
	}

	return res
}

// isNum 檢查字符是否為數字字符（0-9）
// 參數：s - 要檢查的字符
// 返回：num - 對應的數字值，res - 是否為有效數字字符
func isNum(s byte) (num int, res bool) {
	res = true

	// 使用 switch 語句將字符轉換為對應的數字
	switch s {
	case '0':
		num = 0
	case '1':
		num = 1
	case '2':
		num = 2
	case '3':
		num = 3
	case '4':
		num = 4
	case '5':
		num = 5
	case '6':
		num = 6
	case '7':
		num = 7
	case '8':
		num = 8
	case '9':
		num = 9
	default:
		res = false // 非數字字符
	}

	return
}
