func fib(n int) int {
	var recFib func(int) int
	cache := map[int]int{}
	cache[0] = 0
	cache[1] = 1

	recFib = func(v int) int {
		if fib, ok := cache[v]; ok {
			return fib
		}

		return recFib(v-1) + recFib(v-2)
	}

	return recFib(n)
}