// Accepted  1 ms  2.03 MB

func climbStairs(n int) int {
	var climb func(i int, n int) int
	memo := make([]int, n+1)

	climb = func(i int, n int) int {
		if i > n {
			return 0
		}

		if i == n {
			return 1
		}

		if memo[i] > 0 {
			return memo[i]
		}

		memo[i] = climb(i+1, n) + climb(i+2, n)

		return memo[i]
	}

	return climb(0, n)
}