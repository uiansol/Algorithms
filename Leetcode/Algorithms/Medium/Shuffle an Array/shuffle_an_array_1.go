// Accepted  17 ms  7.39 MB

type Solution struct {
	result   []int
	original []int
}

func Constructor(nums []int) Solution {
	copyArr := make([]int, len(nums))

	copy(copyArr, nums)

	return Solution{
		result:   nums,
		original: copyArr,
	}
}

func (this *Solution) Reset() []int {
	return this.original
}

func (this *Solution) Shuffle() []int {
	rand.Shuffle(len(this.original), func(i int, j int) {
		this.result[i], this.result[j] = this.result[j], this.result[i]
	})

	return this.result
}