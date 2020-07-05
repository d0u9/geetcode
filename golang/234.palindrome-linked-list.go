/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func solution1(head *ListNode) bool {
	dummy := ListNode{ Val: 0, Next: nil, }
	rh := &dummy

	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}

		next_slow := slow.Next
		slow.Next = rh.Next
		rh.Next = slow

		slow = next_slow
	}

	//fmt.Printf("fast = %v\n", fast)
	if fast != nil {
		slow = slow.Next
	}

	fast = rh.Next

	for fast != nil && slow != nil {
		//fmt.Printf("fast = %v, slow = %v\n", fast, slow)
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

func isPalindrome(head *ListNode) bool {
	return solution1(head)
}
// @lc code=end

