/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
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
func recur(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}

	rsum := recur(root.Right, sum)
	root.Val = rsum + root.Val
	lsum := recur(root.Left, root.Val)
	return lsum
}

func yuanjing(root *TreeNode, pre *TreeNode) *TreeNode {
	if root == nil {
		return pre
	}

	r := yuanjing(root.Right, pre)
	if r != nil {
		root.Val = root.Val + r.Val
	}
	l := yuanjing(root.Left, root)

	return l
}

func convertBST(root *TreeNode) *TreeNode {
	//recur(root, 0)
	yuanjing(root, nil)
	return root
}
// @lc code=end

