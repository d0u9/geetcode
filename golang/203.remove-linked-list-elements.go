/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var newH, last, cur *ListNode

	cur = head
	newH = head
	for cur != nil {
		if cur.Val == val {
			if last == nil {
				cur = cur.Next
				newH = cur
			} else {
				cur = cur.Next
				last.Next = cur
			}
		} else {
			last = cur
			cur = cur.Next
		}
	}

	return newH
}
// @lc code=end

