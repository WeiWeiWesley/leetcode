package binarysearch

func search(nums []int, target int) int {
	// binary search with two-pointers
	left, right := 0, len(nums)-1

	// binary search core
	for left <= right {

		mid := left + (right-left)/2 //避免溢位不寫 (left+right)/2

		midNum := nums[mid]

		if target == midNum {
			return mid
		} else if target > midNum {
			// 目標大於 mid，左半邊可不看，移動左界至 mid+1
			left = mid + 1
		} else {
			// 目標小於 mid，右半邊可不看，移動右界至 mid-1
			right = mid - 1
		}

	}

	// target doesn't exist in input array
	return -1
}
