// Accepted  12 ms  6.26 MB

func majorityElement(nums []int) int {
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	res, max := 0, 0
	for k, v := range freq {
		if v > max {
			max = v
			res = k
		}
	}

	return res
}