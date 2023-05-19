package searchinsertposition

// nums 是排序過陣列，可以使用 Binary Search 加速搜尋
func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	//Binary Search
	for l <= r {
		m := l + (r-l)/2

		v := nums[m]

		switch {
		case v < target:
			l = m + 1
		case v > target:
			r = m - 1
		default:
			return m // v==target
		}
	}

	return l //目標不在區間內，回傳移動後的結果
}
