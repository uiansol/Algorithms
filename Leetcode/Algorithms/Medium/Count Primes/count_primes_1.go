// Accepted  62 ms  12.53 MB

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	numbers := make([]bool, n)
	numbers[0], numbers[1] = true, true
	limit := int(math.Sqrt(float64(n)) + 1)

	for p := 2; p < limit; p++ {
		if !numbers[p] {
			for m := p * p; m < n; m += p {
				numbers[m] = true
			}
		}
	}

	count := 0
	for _, v := range numbers {
		if !v {
			count++
		}
	}

	return count
}