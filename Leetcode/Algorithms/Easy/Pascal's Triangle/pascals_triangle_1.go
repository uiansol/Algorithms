// Accepted  1 ms  2.46 MB

func generate(numRows int) [][]int {
	var output [][]int
	row := []int{1}
	output = append(output, row)

	for i := 2; i <= numRows; i++ {
		newRow := []int{1}
		for j := 1; j < i-1; j++ {
			newRow = append(newRow, row[j-1]+row[j])
		}
		newRow = append(newRow, 1)

		output = append(output, newRow)
		row = newRow
	}

	return output
}