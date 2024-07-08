func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	if n%2 != 0 {
		return x * myPow(x*x, (n-1)/2)
	}
	return myPow(x*x, n/2)
}