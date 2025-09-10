package lengthoflastword

func lengthOfLastWord(s string) int {
	res := 0        // 紀錄結果
	i := len(s) - 1 // 從尾部開始

	// 從尾部開始跳過空白
	for i >= 0 && s[i] == ' ' { //寫法特性滿則條件時迴圈才會繼續
		i--
	}

	// 從尾部找到的第一個字串開始計算長度
	for i >= 0 && s[i] != ' ' { //寫法特性滿則條件時迴圈才會繼續
		res++
		i--
	}

	return res
}
