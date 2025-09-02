package twosum

// hash table version
func twoSum(nums []int, target int) []int {
	// 定位
	hashMap := map[int]int{}
	for i, v := range nums {
		// 檢查當前數字是否為之前某個數字的補數
		if index, ok := hashMap[v]; ok {
			return  []int{index, i}
		}

		// 記錄當前數字對應的補數索引（如果還沒記錄過）
		if _, ok := hashMap[target-v]; !ok {
			hashMap[target-v] = i
		}
	}

	return []int{}
}
