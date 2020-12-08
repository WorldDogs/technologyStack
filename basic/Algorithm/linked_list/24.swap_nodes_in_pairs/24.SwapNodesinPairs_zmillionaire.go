// Source : https://leetcode.com/problems/swap-nodes-in-pairs/
// Author : zmillionaire
// Date   : 2020-10-11
package main

/*****************************************************************************************************
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 *
 * You may not modify the values in the list's nodes. Only nodes itself may be changed.
 *
 * Example 1:
 *
 * Input: head = [1,2,3,4]
 * Output: [2,1,4,3]
 *
 * Example 2:
 *
 * Input: head = []
 * Output: []
 *
 * Example 3:
 *
 * Input: head = [1]
 * Output: [1]
 *
 * Constraints:
 *
 * 	The number of nodes in the list is in the range [0, 100].
 * 	0 <= Node.val <= 100
 ******************************************************************************************************/

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
// 思路，增加新的头节点进行辅助前两个节点的交换
func swapPairs(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	pre :=&ListNode{Val: 0,Next: head}
	now:=pre
	for now.Next!=nil && now.Next.Next!=nil{
		a,b:=now.Next,now.Next.Next
		now.Next,a.Next = b,b.Next
		b.Next = a
		now = now.Next.Next
	}
	return pre.Next
}
