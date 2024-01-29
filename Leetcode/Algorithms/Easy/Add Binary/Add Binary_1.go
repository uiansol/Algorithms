func addBinary(a string, b string) string {
	la, lb, l := len(a)-1, len(b)-1, max(len(a), len(b))
	var output []rune
	c := 0

	for i := 0; i < l; i++ {
		sum := 0

		if la-i >= 0 && a[la-i] == '1' {
			sum += 1
		}

		if lb-i >= 0 && b[lb-i] == '1' {
			sum += 1
		}
		sum += c

		switch sum {
		case 0:
			output = append(output, '0')
			c = 0
		case 1:
			output = append(output, '1')
			c = 0
		case 2:
			output = append(output, '0')
			c = 1
		case 3:
			output = append(output, '1')
			c = 1
		}
	}

	if c == 1 {
		output = append(output, '1')
	}

	slices.Reverse(output)
	return string(output)
}