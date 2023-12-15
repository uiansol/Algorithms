// Accepted  4 ms  3.5 MB

func intersect(nums1 []int, nums2 []int) []int {
	var result []int

	map1 := make(map[int]int)
	map2 := make(map[int]int)

	for _, v := range nums1 {
		map1[v]++
	}

	for _, v := range nums2 {
		map2[v]++
	}

	for k, v1 := range map1 {
		if v2, ok := map2[k]; ok {
			for i := 0; i < min(v1, v2); i++ {
				result = append(result, k)
			}
		}
	}

	return result
}