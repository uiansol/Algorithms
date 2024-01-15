// Accepted  3 ms  3 MB

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		if !isRowValid(&board, i) {
			return false
		}

		if !isColValid(&board, i) {
			return false
		}

		if !isSquareValid(&board, i) {
			return false
		}
	}

	return true
}

func isRowValid(board *[][]byte, row int) bool {
	chars := make(map[byte]int)

	for i := 0; i < 9; i++ {
		if (*board)[row][i] != '.' {
			chars[(*board)[row][i]]++
		}
	}

	for _, v := range chars {
		if v > 1 {
			return false
		}
	}

	return true
}

func isColValid(board *[][]byte, col int) bool {
	chars := make(map[byte]int)

	for i := 0; i < 9; i++ {
		if (*board)[i][col] != '.' {
			chars[(*board)[i][col]]++
		}
	}

	for _, v := range chars {
		if v > 1 {
			return false
		}
	}

	return true
}

func isSquareValid(board *[][]byte, square int) bool {
	chars := make(map[byte]int)
	si, sj := 0, 0

	switch square {
	case 0:
		si, sj = 0, 0
	case 1:
		si, sj = 0, 3
	case 2:
		si, sj = 0, 6
	case 3:
		si, sj = 3, 0
	case 4:
		si, sj = 3, 3
	case 5:
		si, sj = 3, 6
	case 6:
		si, sj = 6, 0
	case 7:
		si, sj = 6, 3
	case 8:
		si, sj = 6, 6
	}

	for i := si; i < si+3; i++ {
		for j := sj; j < sj+3; j++ {
			if (*board)[i][j] != '.' {
				chars[(*board)[i][j]]++
			}
		}
	}

	for _, v := range chars {
		if v > 1 {
			return false
		}
	}

	return true
}