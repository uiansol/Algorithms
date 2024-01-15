// Accepted  2 ms  2.24 MB

func merge(nums1 []int, m int, nums2 []int, n int) {
	place, idx1, idx2 := m+n-1, m-1, n-1

	for idx1 >= 0 && idx2 >= 0 {
		if nums1[idx1] >= nums2[idx2] {
			nums1[place] = nums1[idx1]
			idx1--
		} else {
			nums1[place] = nums2[idx2]
			idx2--
		}
		place--
	}

	for idx1 >= 0 {
		nums1[place] = nums1[idx1]
		idx1--
		place--
	}
	for idx2 >= 0 {
		nums1[place] = nums2[idx2]
		idx2--
		place--
	}
}