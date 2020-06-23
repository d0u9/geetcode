/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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

func recur(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil  {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	if r := recur(left.Left, right.Right); r == false {
		return false
	}
	if r := recur(left.Right, right.Left); r == false  {
		return false
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}
// @lc code=end

