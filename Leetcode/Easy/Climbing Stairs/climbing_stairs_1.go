// Time Limit Exceeded

func climbStairs(n int) int {
	return climb(0, n)
}

func climb(i int, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}

	return climb(i+1, n) + climb(i+2, n)
}