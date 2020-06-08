/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum = sum - root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}

	if root.Left != nil && root.Right != nil {
		return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
	}

	if root.Left != nil && root.Right == nil {
		return hasPathSum(root.Left, sum)

	}

	if root.Left == nil && root.Right != nil {
		return hasPathSum(root.Right, sum)
	}

	return false
}
