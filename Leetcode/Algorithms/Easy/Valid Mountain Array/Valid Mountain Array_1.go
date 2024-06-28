func validMountainArray(arr []int) bool {
	if len(arr) <= 2 {
		return false
	}

	if arr[0] >= arr[1] || arr[len(arr)-2] <= arr[len(arr)-1] {
		return false
	}

	asc := true
	for i := 0; i < len(arr)-1; i++ {
		if asc {
			if arr[i] < arr[i+1] {
				continue
			} else if arr[i] == arr[i+1] {
				return false
			} else {
				asc = false
			}
		} else {
			if arr[i] > arr[i+1] {
				continue
			} else {
				return false
			}
		}
	}

	return true
}