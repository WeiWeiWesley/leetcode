package minsizesubarrsum

func minSubArrayLen(target int, nums []int) int {

	result := len(nums) + 1 //預設最大長度 + 1，反向排除
	sum := 0
	head := 0

	//持續累加直到超過目標
	for k, v := range nums {

		sum += v

		//超過後調整起點位置
		for sum >= target {

			//確認當前長度
			currentLen := k - head + 1

			if result > currentLen {
				result = currentLen
			}

			sum -= nums[head]
			head++
		}
	}

	if result == len(nums) + 1 {
		return 0
	}

	return result
}
