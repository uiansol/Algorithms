// Accepted  0 ms  3.9 MB

func levelOrder(root *TreeNode) [][]int {
	var solve func(node *TreeNode, level int)
	levels := [][]int{}
	if root == nil {
		return levels
	}

	solve = func(node *TreeNode, level int) {
		if len(levels) == level {
			levels = append(levels, []int{})
		}

		levels[level] = append(levels[level], node.Val)

		if node.Left != nil {
			solve(node.Left, level+1)
		}
		if node.Right != nil {
			solve(node.Right, level+1)
		}
	}

	solve(root, 0)

	return levels
}