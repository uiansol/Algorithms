func getRow(numRows int) []int {
	row := []int{1}

	for i := 2; i <= numRows+1; i++ {
		newRow := []int{1}
		for j := 1; j < i-1; j++ {
			newRow = append(newRow, row[j-1]+row[j])
		}
		newRow = append(newRow, 1)

		row = newRow
	}

	return row
}