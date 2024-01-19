// Accepted  3 ms  3.47 MB

func sortedArrayToBST(nums []int) *TreeNode {
	var solve func(left int, right int) *TreeNode

	solve = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}

		p := (left + right) / 2

		root := TreeNode{Val: nums[p]}
		root.Left = solve(left, p-1)
		root.Right = solve(p+1, right)

		return &root
	}

	return solve(0, len(nums)-1)
}