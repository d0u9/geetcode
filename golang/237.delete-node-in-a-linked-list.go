/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	c, n := node, node.Next
	for n != nil {
		c.Val = n.Val
		if n.Next != nil {
			c = n
		} else {
			c.Next = nil
		}
		n = n.Next
	}
}
// @lc code=end

