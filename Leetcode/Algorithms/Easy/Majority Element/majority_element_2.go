// Accepted  8 ms  6.43 MB

func majorityElement(nums []int) int {
	l := len(nums)
	count := l / 2

	for true {
		randomIndex := rand.Intn(l)
		randomElement := nums[randomIndex]
		sum := 0
		for i := 0; i < l; i++ {
			if nums[i] == randomElement {
				sum++
			}
		}
		if sum > count {
			return randomElement
		}
	}

	return 0
}