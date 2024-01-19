// Accepted  15 ms  6 MB

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	quo, _ := math.Modf(math.Log10(float64(n)) / math.Log10(3))

	return n == int(math.Pow(3, quo))
}