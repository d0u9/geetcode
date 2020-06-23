/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
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
func recv(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	if l := recv(a.Left, b.Left); l == false {
		return false
	}
	if r := recv(a.Right, b.Right); r == false {
		return false
	}

	return true
}

func iter(p *TreeNode, q *TreeNode) bool {
	var stack = make([][]*TreeNode, 0)
	slen := 0
	for true {
		if p == nil && q == nil {
			if slen == 0{
				break
			}
			p, q = stack[slen - 1][0], stack[slen - 1][1]
			p, q = p.Right, q.Right
			stack = stack[:(slen-1)]
			slen--
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		z := []*TreeNode {p, q,}
		stack = append(stack, z)
		slen++
		p, q = p.Left, q.Left
	}
	return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// return recv(p, q)
	return iter(p, q)
}
// @lc code=end

