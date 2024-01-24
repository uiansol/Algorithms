func isHappy(n int) bool {
	results := make(map[int]bool)

	for {
		if n == 1 {
			return true
		}

		nextN := 0
		for n > 0 {
			rest := n % 10
			nextN += rest * rest
			n /= 10
		}

		if _, ok := results[nextN]; ok {
			return false
		}

		results[nextN] = true
		n = nextN
	}
}