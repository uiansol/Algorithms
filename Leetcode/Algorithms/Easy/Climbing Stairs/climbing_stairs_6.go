// Accepted  0 ms  1.92 MB

func climbStairs(n int) int {
	sqrt5 := 2.2360679775
	phi := (1 + sqrt5) / 2
	psi := (1 - sqrt5) / 2

	return int(((math.Pow(phi, float64(n+1)) - math.Pow(psi, float64(n+1))) / sqrt5))
}