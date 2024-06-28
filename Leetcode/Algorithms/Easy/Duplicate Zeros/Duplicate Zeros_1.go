func duplicateZeros(arr []int) {
	zeros, idx, last := 0, -1, -1

	for idx < len(arr)-1 {
		idx++
		last++
		if arr[last] == 0 {
			idx++
			if idx < len(arr) {
				zeros++
			}
		}
	}

	lastZero, count := 0, 0
	for i, v := range arr {
		if v == 0 && count < zeros {
			lastZero = i
			count++
		}
	}

	idx = len(arr) - 1
	for idx >= 0 && last >= 0 {
		arr[idx] = arr[last]
		idx--
		if arr[last] == 0 && zeros > 0 && last <= lastZero {
			arr[idx] = arr[last]
			idx--
			zeros--
		}
		last--
	}
}