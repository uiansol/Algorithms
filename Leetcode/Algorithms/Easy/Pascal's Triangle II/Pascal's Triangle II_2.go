func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	rowBefore := getRow(rowIndex - 1)
	var row []int

	row = append(row, 1)
	for i := 0; i < len(rowBefore)-1; i++ {
		row = append(row, rowBefore[i]+rowBefore[i+1])
	}
	row = append(row, 1)

	return row
}
