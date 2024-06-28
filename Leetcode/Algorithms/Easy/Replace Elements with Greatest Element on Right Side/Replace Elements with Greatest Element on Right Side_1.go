func replaceElements(arr []int) []int {
	aux, higher := 0, -1

	for i := len(arr) - 1; i >= 0; i-- {
		aux = arr[i]
		arr[i] = higher
		if aux > higher {
			higher = aux
		}
	}

	return arr
}