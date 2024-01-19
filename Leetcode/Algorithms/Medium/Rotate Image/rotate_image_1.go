// Accepted  1 ms  2.3 MB

func rotate(matrix [][]int) {
	length := len(matrix)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if j > i {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}

		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}