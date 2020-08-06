/*
 * @lc app=leetcode id=590 lang=golang
 *
 * [590] N-ary Tree Postorder Traversal
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

	for _, n := range(root.Children) {
		solution1(n, ret)
	}
	*ret = append(*ret, root.Val)
}

func postorder(root *Node) []int {
	ret := []int{}
	solution1(root, &ret)

	return ret
}
// @lc code=end

