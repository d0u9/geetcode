/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recr(root *TreeNode, acc, sum int) bool {
	if root == nil {
		return false
	}
	v := root.Val + acc
	if v == sum && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		if l := recr(root.Left, v, sum); l == true {
			return true
		}
	}
	if root.Right != nil {
		if r := recr(root.Right, v, sum); r == true {
			return true
		}
	}

	return false
}
func hasPathSum(root *TreeNode, sum int) bool {
	return recr(root, 0, sum)
}
// @lc code=end

