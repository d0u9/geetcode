/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func recur(root *TreeNode) (int, int) {
	if root == nil {
		return 0, -1
	}

	lmaxdia, lmaxpath := recur(root.Left)
	rmaxdia, rmaxpath := recur(root.Right)

	maxdia := lmaxpath + rmaxpath + 2

	maxpath := MaxInt(lmaxpath, rmaxpath)
	//fmt.Println(root.Val, maxpath)
	if maxpath == -1 {
		maxpath = 0
	} else {
		maxpath++
	}

	a ,b := MaxInt(maxdia, MaxInt(lmaxdia, rmaxdia)), maxpath

	//fmt.Println(root.Val, a, b)
	return a, b
}

func diameterOfBinaryTree(root *TreeNode) int {
	d, _ := recur(root)
	return d
}
// @lc code=end

