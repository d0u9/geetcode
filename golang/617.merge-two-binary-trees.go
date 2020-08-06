/*
 * @lc app=leetcode id=617 lang=golang
 *
 * [617] Merge Two Binary Trees
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
func recur(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	this := TreeNode {
		Val: 0,
		Left: nil,
		Right: nil,
	}

	var l1, l2, r1, r2 *TreeNode
	if t1 != nil {
		this.Val = this.Val + t1.Val
		l1 = t1.Left
		r1 = t1.Right
	}
	if t2 != nil {
		this.Val = this.Val + t2.Val
		l2 = t2.Left
		r2 = t2.Right
	}

	this.Left = recur(l1, l2)
	this.Right = recur(r1, r2)

	return &this
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    return recur(t1, t2)
}
// @lc code=end

