package searchinsertposition

func searchInsert(nums []int, target int) int {

	previous := 0

	if nums[len(nums)-1] < target {
		return len(nums)
	}

	for i, v := range nums {
		if v == target {
			return i
		}

		if previous < target && v > target {
			return i
		}

		previous = v
	}


	return 0
}