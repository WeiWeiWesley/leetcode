package removeelement

// 這題除了回傳值外，也會檢查 nums 最終內容，所以需要改變 nums(same address)
func removeElement(nums []int, val int) int {
    slow := 0 // 指向下一個要放置非val元素的位置

    for fast := 0; fast < len(nums); fast++ {
        if nums[fast] != val {
            nums[slow] = nums[fast]
            slow++
        }
    }

    return slow // 回傳移除後陣列的長度
}
