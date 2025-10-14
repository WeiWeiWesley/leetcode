package minsizesubarrsum

// 解法：滑動窗口（Sliding Window）
// 思路：維護一個窗口 [head, k]，動態調整窗口大小
// 1. 右邊界(k)不斷向右擴展，累加 sum
// 2. 當 sum >= target 時，嘗試收縮左邊界(head)來找最小長度
// 3. 因為都是正整數，移除左邊元素會讓 sum 變小，具有單調性
// 時間複雜度：O(n) - 每個元素最多被訪問兩次（加入和移除）
// 空間複雜度：O(1) - 只使用固定變數
func minSubArrayLen(target int, nums []int) int {

	result := len(nums) + 1 // 預設最大長度 + 1，用來判斷是否找到解
	sum := 0                // 當前窗口的總和
	head := 0               // 窗口左邊界

	// 持續累加直到超過目標（右邊界 k 向右擴展）
	for k, v := range nums {

		sum += v // 擴展窗口：加入右邊界元素

		// 當窗口和 >= target 時，嘗試收縮左邊界找最小長度
		for sum >= target {

			// 確認當前窗口長度
			currentLen := k - head + 1

			// 更新最小長度
			if result > currentLen {
				result = currentLen
			}

			// 收縮窗口：移除左邊界元素
			sum -= nums[head]
			head++ // 左邊界右移
		}
	}

	// 如果 result 沒有被更新過，表示沒有符合條件的子陣列
	if result == len(nums)+1 {
		return 0
	}

	return result
}
