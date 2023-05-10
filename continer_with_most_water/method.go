package continerwithmostwater

// O(n^2)
// func maxArea(height []int) int {

// 	n := len(height)
// 	max := 0
// 	right := n - 1

// 	for i := 0; i < right; i++ {
// 		for j := right; j > i; j-- {

// 			h := height[i]

// 			if height[i] > height[j] {
// 				h = height[j]
// 			}

// 			water := (j - i) * h

// 			if water > max {
// 				max = water
// 			}

// 		}
// 	}

// 	return max
// }

// O(n)
func maxArea(height []int) int {
	n := len(height)

	result := 0
	l := 0
	r := n - 1

	for l < r {
		currArea := min(height[l], height[r]) * (r - l)

		if result < currArea {
			result = currArea
		}

		// 如果左高低於右高，由於下個區間寬度減少，無須動右高，移動左高找出最大值
		// 移動較低的一邊
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
