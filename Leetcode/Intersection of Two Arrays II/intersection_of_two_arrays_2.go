// Accepted  4 ms  2.6 MB

func intersect(nums1 []int, nums2 []int) []int {
	var result []int

	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	i, j := 0, 0

	for i < length1 && j < length2 {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else {
			if nums1[i] < nums2[j] {
				i++
			} else {
				j++
			}
		}
	}

	return result
}