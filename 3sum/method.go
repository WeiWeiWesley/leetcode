package threesum

import (
	"sort"
)

// ThreeSum 找到所有不重複的三元組，使得三個數的和為 0
// 使用排序 + 雙指針的方法
// 時間複雜度: O(n²), 空間複雜度: O(1)
func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)

	// 邊界條件：至少需要 3 個元素
	if n < 3 {
		return res
	}

	// 1. 先排序陣列，讓重複元素聚集在一起，方便去重
	sort.Ints(nums)

	// 2. 固定第一個數，用雙指針找另外兩個數
	for i := 0; i < n-2; i++ {
		// 去重：如果當前數和前一個數相同，跳過
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1  // 左指針
		right := n - 1 // 右指針

		// 3. 雙指針尋找另外兩個數
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// 找到符合條件的三元組
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// 去重：跳過相同的左指針值
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 去重：跳過相同的右指針值
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				// 總和太小，左指針向右移動（增加總和）
				left++
			} else {
				// 總和太大，右指針向左移動（減少總和）
				right--
			}
		}
	}

	return res
}
