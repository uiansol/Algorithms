// Accepted  0 ms  4.5 MB

func maxDepth(root *TreeNode) int {
	return solve(root, 1)
}

func solve(node *TreeNode, deep int) int {
	deepLeft, deepRight := 0, 0

	if node == nil {
		return 0
	}

	deepLeft = 1 + solve(node.Left, deep)
	deepRight = 1 + solve(node.Right, deep)

	return max(deepLeft, deepRight)
}