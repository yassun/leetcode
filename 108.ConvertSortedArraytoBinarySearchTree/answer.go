/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	c := len(nums) / 2
	return &TreeNode{
		Val:   nums[c],
		Left:  sortedArrayToBST(nums[:c]),
		Right: sortedArrayToBST(nums[c+1:]),
	}
}
