package twosum

//simple loop version
// func twoSum(nums []int, target int) []int {

// 	for i := range nums {
// 		for j := range nums {
// 			if i < j && nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}

// 	return []int{}
// }

//hash table version
func twoSum(nums []int, target int) []int {

	table := map[int]int{}

	for i, n := range nums {
		table[n] = i
	}

	res := []int{}

	l := len(nums)

	for i := 0; i < l; i++ {

		if j, ok := table[target-nums[i]]; ok && j != i {
			return []int{i, j}
		}
		
	}

	return res
}
