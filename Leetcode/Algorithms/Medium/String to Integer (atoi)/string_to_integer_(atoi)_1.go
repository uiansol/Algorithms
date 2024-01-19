// Accepted  0 ms  2.21 MB

func myAtoi(s string) int {
	sign, result, index, n := 1, 0, 0, len(s)

	INT_MAX := int(math.Pow(2, 31) - 1)
	INT_MIN := int(-math.Pow(2, 31))

	for index < n && s[index] == ' ' {
		index++
	}

	if index < n && s[index] == '+' {
		sign = 1
		index++
	} else if index < n && s[index] == '-' {
		sign = -1
		index++
	}

	for index < n && unicode.IsDigit(rune(s[index])) { //s[index] >= '0' && s[index] <= '9' {
		digit := int(s[index]) - '0'

		if (result > INT_MAX/10) || (result == INT_MAX/10 && digit > INT_MAX%10) {
			if sign == 1 {
				return INT_MAX
			} else {
				return INT_MIN
			}
		}

		result = 10*result + digit
		index++
	}

	return sign * result
}