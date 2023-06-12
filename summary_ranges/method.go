package summaryranges

import "strconv"

func summaryRanges(nums []int) []string {
	// 建立一個空陣列來存放結果
	result := []string{}

	// 如果輸入陣列為空，直接返回空結果陣列
	if len(nums) == 0 {
		return result
	}

	// 初始化起始範圍的起點和終點
	start := nums[0]
	end := nums[0]

	// 遍歷整個陣列
	for i := 1; i < len(nums); i++ {
		// 如果當前元素與前一個元素連續，更新終點
		if nums[i] == nums[i-1]+1 {
			end = nums[i]
		} else {
			// 如果當前元素不與前一個元素連續，表示範圍中斷
			// 由於題目指定輸出，僅需知道 start -> end 即可組成字串
			result = append(result, formatRange(start, end))

			//將範圍加入結果陣列並更新起點和終點
			start = nums[i]
			end = nums[i]
		}
	}

	// 將最後一個範圍加入結果陣列
	result = append(result, formatRange(start, end))

	return result
}

// 格式化範圍字串的輔助函數 ex 2->4
func formatRange(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	}

	return strconv.Itoa(start) + "->" + strconv.Itoa(end)
}
