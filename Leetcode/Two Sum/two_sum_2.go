// Accepted  4 ms  4.4 MB

func twoSum(nums []int, target int) []int {
	store := make(map[int]int)

	for i, v := range nums {
		if val, ok := store[target-v]; ok {
			return []int{i, val}
		}
		store[v] = i
	}

	return nil
}