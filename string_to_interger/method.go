package stringtointerger

import (
	"math"
	"strings"
)

// 48 ~ 57 => 0 ~ 9
func myAtoi(s string) int {

	s = strings.Trim(s, " ")

	n := len(s)

	if n == 0 {
		return 0
	}

	res := 0
	start := 0
	sign := 1

	switch string(s[0]) {
	case "-":
		sign = -1
		start = 1
	case "+":
		sign = 1
		start = 1
	}

	for i := start; i < n; i++ {

		// '' 代表 byte
		if s[i] < '0' || s[i] > '9' {
			return res * sign
		}

		// 將原有結果進位 + 新個位數
		res = res*10 + int(s[i]-'0')

		//須於回圈內檢查，避免溢位判斷錯誤
		if res*sign > math.MaxInt32 {
			return math.MaxInt32
		}

		if res*sign < math.MinInt32 {
			return math.MinInt32
		}

	}

	res *= sign

	return res
}
