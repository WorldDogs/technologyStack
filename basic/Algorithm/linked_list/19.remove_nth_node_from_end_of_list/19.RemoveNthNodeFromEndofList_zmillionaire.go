// Source : https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Author : zmillionaire
// Date   : 2020-10-11
package main

/*****************************************************************************************************
 *
 * Given the head of a linked list, remove the nth node from the end of the list and return its head.
 *
 * Follow up: Could you do this in one pass?
 *
 * Example 1:
 *
 * Input: head = [1,2,3,4,5], n = 2
 * Output: [1,2,3,5]
 *
 * Example 2:
 *
 * Input: head = [1], n = 1
 * Output: []
 *
 * Example 3:
 *
 * Input: head = [1,2], n = 1
 * Output: [1]
 *
 * Constraints:
 *
 * 	The number of nodes in the list is sz.
 * 	1 <= sz <= 30
 * 	0 <= Node.val <= 100
 * 	1 <= n <= sz
 ******************************************************************************************************/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre:=&ListNode{Val: 0}
	pre.Next = head
	var fast,slow = pre,pre

	for i:=0;i<n+1;i++{
		if fast == nil{
			return nil
		}
		fast = fast.Next
	}
	for fast!=nil{
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return pre.Next
}
