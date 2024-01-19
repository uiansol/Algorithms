// Accepted  12 ms  5.4 MB

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func isValidBST(root *TreeNode) bool {
	return solve(root, MinInt, MaxInt)
}

func solve(node *TreeNode, low int, high int) bool {
	if node == nil {
		return true
	}

	if node.Val <= low || node.Val >= high {
		return false
	}

	return solve(node.Left, low, node.Val) && solve(node.Right, node.Val, high)
}