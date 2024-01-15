// Accepted  3 ms  4.4 MB

func maxDepth(root *TreeNode) int {
	deepLeft, deepRight := 0, 0

	if root == nil {
		return 0
	}

	deepLeft = maxDepth(root.Left)
	deepRight = maxDepth(root.Right)

	return max(deepLeft, deepRight) + 1
}