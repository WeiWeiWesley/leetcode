package zigzagconversion

/*
Z字型轉換 - 規律解法

核心思路：觀察每一行字符的索引規律，直接計算出每個位置的字符，無需創建矩陣

範例：s = "PAYPALISHIRING", numRows = 4
視覺化排列：
索引:  0     6       12      (第 0 行) - 間隔: 6

	1   5 7    11 13      (第 1 行) - 間隔: 4, 2 交替
	2 4   8 10    14      (第 2 行) - 間隔: 2, 4 交替
	3     9       15      (第 3 行) - 間隔: 6

規律：
- cycleLen = 2 * numRows - 2 (一個完整Z字週期的長度)
- 第 0 行和最後一行：字符索引間隔固定為 cycleLen
- 中間行 i：有兩個間隔交替出現
  - 間隔1：cycleLen - 2*i (斜向上的間隔)
  - 間隔2：2*i (向下的間隔)
*/
func convert(s string, numRows int) string {
	// 邊界情況：字符串過短或只有一行時，直接返回原字符串
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	n := len(s)
	// 計算一個完整Z字週期的長度
	// 例如：numRows=4 時，週期為 0→1→2→3→2→1，長度為 6
	cycleLen := 2*numRows - 2
	result := make([]byte, 0, n)

	// 遍歷每一行
	for row := 0; row < numRows; row++ {
		// 當前行的第一個字符索引就是行號
		index := row

		// 遍歷當前行的所有字符
		for index < n {
			// 添加當前索引的字符
			result = append(result, s[index])

			// 計算下一個字符的索引
			if row == 0 || row == numRows-1 {
				// 第一行和最後一行：間隔固定為 cycleLen
				index += cycleLen
			} else {
				// 中間行：需要處理兩個交替的間隔

				// 第一個間隔：斜向上的距離 (到下一個"谷底"的距離)
				gap1 := cycleLen - 2*row
				index += gap1

				// 如果還有字符，添加斜線上的字符
				if index < n {
					result = append(result, s[index])
					// 第二個間隔：向下的距離 (到下一個"波峰"的距離)
					gap2 := 2 * row
					index += gap2
				}
			}
		}
	}

	return string(result)
}
