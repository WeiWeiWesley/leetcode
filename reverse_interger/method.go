package reverseinterger

import (
	"math"
)

//善用餘數
func reverse(x int) int {

	sign := 1

	//負數判斷
	if x < 0 {
		sign = -1
		x = sign * x
	}

	var result int
	for x > 0 {
		// 取出最後一位，並將前次結果乘上倍數
		result = result*10 + x%10

		// 題目範圍限制
		if sign*result >= math.MaxInt32 || sign*result <= math.MinInt32 {
			return 0
		}

		// 將已加入新結果的部分移至小數點後
		x = x / 10
	}

	// + or - 校正
	return sign * result
}
