package searchinsertposition

// O(log n) 有限制搜尋的時間
func SearchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r-l)/2 // 尋找中間點

		if target > nums[m] {
			l = m + 1 // 向右邊區塊尋找，變更左界
		} else if target < nums[m] {
			r = m - 1 // 向左邊區塊尋找，變更右界
		} else {
			return m // target == nums[m] // 找到相同的數值
		}
	}

	// 當迴圈結束時，l 就是插入位置
	return l
}
