func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	first, second, next := 0, 1, 1
	for i := 1; i < n; i++ {
		next = first + second
		first, second = second, next
	}

	return next
}