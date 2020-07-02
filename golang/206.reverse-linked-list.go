/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var cur *ListNode = nil

	for head != nil {
		next := head.Next
		head.Next = cur
		cur = head
		head = next
	}

	return cur
}
// @lc code=end

