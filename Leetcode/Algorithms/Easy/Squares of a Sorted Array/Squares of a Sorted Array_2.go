func sortedSquares(nums []int) []int {
	var temp []int
	idx, nIdx := -1, 0

	for _, v := range nums {
		squared := v * v
		if v < 0 {
			temp = append(temp, squared)
			idx += 1
		} else {
			for idx >= 0 && squared > temp[idx] {
				nums[nIdx] = temp[idx]
				nIdx++
				idx--
			}
			nums[nIdx] = squared
			nIdx++
		}
	}

	for idx >= 0 {
		nums[nIdx] = temp[idx]
		nIdx++
		idx--
	}

	return nums
}