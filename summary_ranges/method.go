package summaryranges

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var result []string
	start := 0

	for i := 1; i <= len(nums); i++ {
		// 當到達陣列末尾或發現不連續的數字時
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			if start == i-1 {
				// 單一數字
				result = append(result, fmt.Sprintf("%d", nums[start]))
			} else {
				// 範圍
				result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			}
			start = i
		}
	}

	return result
}
