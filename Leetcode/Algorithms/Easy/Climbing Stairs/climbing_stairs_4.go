// Accepted  1 ms  1.94 MB

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	first, second, third := 1, 2, 0

	for i := 3; i <= n; i++ {
		third = first + second
		first = second
		second = third
	}

	return second
}