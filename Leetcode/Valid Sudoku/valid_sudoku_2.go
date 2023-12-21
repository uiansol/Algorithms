// Accepted  1 ms  2.62 MB

func isValidSudoku(board [][]byte) bool {
	rows := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	cols := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	boxes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			pos := int(board[r][c]) - 1

			if rows[r]&(1<<pos) != 0 {
				return false
			}
			rows[r] |= (1 << pos)

			if cols[c]&(1<<pos) != 0 {
				return false
			}
			cols[c] |= (1 << pos)

			idx := (r/3)*3 + c/3
			if boxes[idx]&(1<<pos) != 0 {
				return false
			}
			boxes[idx] |= (1 << pos)
		}
	}

	return true
}