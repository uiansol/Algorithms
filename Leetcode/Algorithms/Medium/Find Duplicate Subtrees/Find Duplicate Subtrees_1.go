/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	table := make(map[string][]*TreeNode)

	treeRepresentation(root, table)

	var result []*TreeNode
	for _, value := range table {
		if len(value) >= 2 {
			result = append(result, value[0])
		}
	}

	return result
}

func treeRepresentation(node *TreeNode, table map[string][]*TreeNode) string {
	if node == nil {
		return "null"
	}

	rep := "[" + strconv.Itoa(node.Val) + "]" + treeRepresentation(node.Left, table) + treeRepresentation(node.Right, table)
	table[rep] = append(table[rep], node)

	return rep
}