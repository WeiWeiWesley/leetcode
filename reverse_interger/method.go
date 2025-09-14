package reverseinterger

import "math"

// reverse 反轉一個 32-bit 整數的數字順序
// 如果反轉後的數字超過 32-bit 整數範圍 [-2^31, 2^31-1]，則回傳 0
// 時間複雜度: O(log n)，其中 n 是輸入數字的位數
// 空間複雜度: O(1)
func reverse(x int) int {
	// 使用 math 套件來取得 32-bit 整數範圍
	maxInt32 := math.MaxInt32 // 2147483647 (2^31 - 1)
	minInt32 := math.MinInt32 // -2147483648 (-2^31)

	// 記錄原始數字是否為負數
	signed := false

	// 處理負數：先轉為正數處理，最後再轉回負數
	if x < 0 {
		signed = true
		x = x * -1 // 轉為正數
	}

	// 用來儲存反轉後的結果
	y := 0

	// 反轉數字：使用餘數運算從個位數開始取
	for x > 0 {
		// 預防溢出檢查：在執行 y*10 + x%10 之前先檢查是否會溢出
		// 檢查條件：
		// 1. y > maxInt32/10：如果 y 超過 maxInt32/10，那麼 y*10 就會溢出
		// 2. y == maxInt32/10 && x%10 > maxInt32%10：如果 y 等於 maxInt32/10，
		//    那麼需要檢查加上 x%10 後是否會超過 maxInt32
		if y > maxInt32/10 || (y == maxInt32/10 && x%10 > maxInt32%10) {
			return 0 // 會溢出，回傳 0
		}

		// 將當前數字的最後一位加到結果的最前面
		y = y*10 + x%10
		// 移除已處理的最後一位
		x = x / 10
	}

	// 如果原始數字是負數，將結果轉為負數
	if signed {
		y = y * -1
	}

	// 最終檢查：確保結果在 32-bit 整數範圍內
	// 這個檢查是雙重保險，雖然前面的檢查應該已經涵蓋了所有情況
	if y > maxInt32 || y < minInt32 {
		return 0
	}

	return y
}
