// Accepted  47 ms  7.25 MB

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		n1 := int64(arr[i])
		b1 := fmt.Sprint(strconv.FormatInt(n1, 2))
		c1 := strings.Count(b1, "1")

		n2 := int64(arr[j])
		b2 := fmt.Sprint(strconv.FormatInt(n2, 2))
		c2 := strings.Count(b2, "1")

		if c1 < c2 {
			return true
		} else if c2 < c1 {
			return false
		} else {
			return n1 < n2
		}
	})

	return arr
}