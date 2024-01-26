func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	inside := func(a, b int) bool {
		return a >= 0 && a < m && b >= 0 && b < n
	}

	var result []int
	i, j := 0, 0
	up := true

	for {
		result = append(result, mat[i][j])

		if up {
			if inside(i-1, j+1) {
				i--
				j++
			} else if inside(i, j+1) {
				up = false
				j++
			} else if inside(i+1, j) {
				up = false
				i++
			} else {
				break
			}
		} else {
			if inside(i+1, j-1) {
				i++
				j--
			} else if inside(i+1, j) {
				up = true
				i++
			} else if inside(i, j+1) {
				up = true
				j++
			} else {
				break
			}
		}
	}

	return result
}
