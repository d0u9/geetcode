/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
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

func recur(root *TreeNode, left bool) int {
	if root == nil {
		return 0
	}

	l, r := recur(root.Left, true), recur(root.Right, false)
	if root.Left == nil && root.Right == nil {
		// current is a leaf node
		if left {
			//fmt.Println(root.Val)
			return root.Val
		} else {
			return 0
		}
	}
	return l + r
}

func sumOfLeftLeaves(root *TreeNode) int {
    return recur(root, false)
}
// @lc code=end

