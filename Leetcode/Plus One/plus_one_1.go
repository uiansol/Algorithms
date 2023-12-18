// Accepted  2 ms  2.1 MB

func plusOne(digits []int) []int {
	length := len(digits)
	carry := 0

	for i := length - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			carry = 0
			break
		} else {
			digits[i] = 0
			carry = 1
		}
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}