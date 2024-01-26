func spiralOrder(matrix [][]int) []int {
	x1, x2, y1, y2 := 0, len(matrix), 0, len(matrix[0])
	l := x2 * y2

	inside := func(a, b int) bool {
		return a >= x1 && a < x2 && b >= y1 && b < y2
	}

	i, j := 0, 0
	direction := 0
	var result []int

	for l > 0 {
		result = append(result, matrix[i][j])

		if direction == 0 {
			if inside(i, j+1) {
				j++
			} else {
				direction = 1
				x1++
			}
		}

		if direction == 1 {
			if inside(i+1, j) {
				i++
			} else {
				direction = 2
				y2--
			}
		}

		if direction == 2 {
			if inside(i, j-1) {
				j--
			} else {
				direction = 3
				x2--
			}
		}

		if direction == 3 {
			if inside(i-1, j) {
				i--
			} else {
				direction = 0
				y1++
				j++
			}
		}
		l--
	}

	return result
}