// Accepted  1 ms  2.22 MB

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2+n%2; i++ {
		for j := 0; j < n/2; j++ {
			tmp := matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]
			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}