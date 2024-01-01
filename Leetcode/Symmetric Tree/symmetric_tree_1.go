// Accepted  3 ms  2.78 MB

func isSymmetric(root *TreeNode) bool {
	return solve(root, root)
}

func solve(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return (node1.Val == node2.Val) && solve(node1.Right, node2.Left) && solve(node1.Left, node2.Right)
}