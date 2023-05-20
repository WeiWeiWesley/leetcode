package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	// array tail 的 pointer
	p1 := m - 1
	p2 := n - 1

	// 合併後的 tail pointer
	p := m + n - 1

	// 由後向前比较 nums1 和 nums2 元素，將較大的 val 取代在 num1 的位置
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// nums2 的值可能沒取完，需把剩餘 val 加到 num1
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
