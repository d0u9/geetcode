/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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

func recur(root *TreeNode, prefix string, ret *[]string) bool {
	if root == nil {
		return true
	}

	p := fmt.Sprintf("%s%d", prefix, root.Val)
	l := recur(root.Left, p + "->", ret)
	r := recur(root.Right, p + "->", ret)

	if l == true && r == true {
		*ret = append(*ret, p)
	}

	return false
}

func binaryTreePaths(root *TreeNode) []string {
	ret := []string{}
	recur(root, "", &ret)
	return ret
}
// @lc code=end

