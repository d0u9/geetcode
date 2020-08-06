/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func solution1(root *Node, ret *[]int) {
	if root == nil {
		return
	}

	*ret = append(*ret, root.Val)
	for _, n := range(root.Children) {
		solution1(n, ret)
	}
}

func preorder(root *Node) []int {
	ret := []int{}
	solution1(root, &ret)
	return ret
}
// @lc code=end

