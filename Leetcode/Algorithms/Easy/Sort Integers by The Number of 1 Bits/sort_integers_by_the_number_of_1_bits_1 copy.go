// Accepted  8 ms  3.40 MB

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		c1 := hammingWeight(arr[i])
		c2 := hammingWeight(arr[j])

		if c1 == c2 {
			return arr[i] < arr[j]
		}
		return c1 < c2
	})

	return arr
}

func hammingWeight(n int) int {
	weight := 0

	for n > 0 {
		weight++
		n &= (n - 1)
	}

	return weight
}