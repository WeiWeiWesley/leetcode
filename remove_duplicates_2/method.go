package removeduplicates2

func removeDuplicates(nums []int) int {
	// 檢查輸入陣列是否為空
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return 1
	}

	// 定義兩個指針：slow 和 fast
	slow := 2 // slow 指向下一個要填入元素的位置，初始化為 2

	// 遍歷整個陣列
	for fast:=2; fast < len(nums); fast++ {
		// 比較 nums[fast] 和 nums[slow-2] 的值
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast] // 如果不相等，表示遇到了一個新的元素，將其放入 slow 指向的位置
			slow++                  // slow 往後移動一格
		}
	}

	return slow // 返回 slow 的值，即為去重後的陣列長度
}
