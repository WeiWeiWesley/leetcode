package removeduplicates

func removeDuplicates(nums []int) int {
	slow := 1 // 第一個一定是 unique

	for fast := 1; fast < len(nums); fast++ {
		if nums[slow-1] != nums[fast] { // 如果相同留在原地等待
			nums[slow] = nums[fast] // 向前覆蓋
			slow++                  // 不同在繼續向前
		}
	}

	return slow
}
