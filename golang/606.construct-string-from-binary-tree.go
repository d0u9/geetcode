/*
 * @lc app=leetcode id=606 lang=golang
 *
 * [606] Construct String from Binary Tree
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

func recur(t *TreeNode, str *[]byte) bool {
	if t == nil {
		return false
	}

	v := []byte(strconv.Itoa(t.Val))
	*str = append(*str, v...)

	*str = append(*str, '(')
	l := recur(t.Left, str)
	*str = append(*str, ')')

	*str = append(*str, '(')
	r := recur(t.Right, str)
	if r {
		*str = append(*str, ')')
	} else {
		if l {
			*str = (*str)[:len(*str)-1]
		} else {
			*str = (*str)[:len(*str)-3]
		}
	}
	return true
}

func tree2str(t *TreeNode) string {
	ret := make([]byte, 0)
	recur(t, &ret)
	return string(ret)
}
// @lc code=end

