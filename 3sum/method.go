package threesum

import (
	"sort"
)

// sort
// double loop
// 3 index pointer
func threeSum(nums []int) [][]int {

	res := [][]int{}

	n := len(nums)	

	//排序過可用來避免重複的組合
	sort.Ints(nums)

	//移動最左
	for indexA := 0; indexA < n-2; indexA++ {

		indexB := indexA + 1
		indexC := n - 1

		//避免重複的組合(左起)
		if indexA > 0 && nums[indexA] == nums[indexA-1] {
			continue
		}

		//移動內層左右
		for indexB < indexC {

			sum := nums[indexA] + nums[indexB] + nums[indexC]

			//避免重複的組合(右起)
			if indexC != n-1 && nums[indexC] == nums[indexC+1] {
				indexC--
				continue
			}

			if sum == 0 {
				res = append(res, []int{nums[indexA], nums[indexB], nums[indexC]})

				indexC--
				continue
			}

			if sum > 0 {
				indexC--
				continue
			}

			indexB++
		}
	}

	return res
}
