// Source : https://leetcode.com/problems/partition-list/
// Author : zmillionaire
// Date   : 2020-10-11
package main

/*****************************************************************************************************
 *
 * Given a linked list and a value x, partition it such that all nodes less than x come before nodes
 * greater than or equal to x.
 *
 * You should preserve the original relative order of the nodes in each of the two partitions.
 *
 * Example:
 *
 * Input: head = 1->4->3->2->5->2, x = 3
 * Output: 1->2->2->4->3->5
 *
 ******************************************************************************************************/

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head==nil{
		return head
	}
	small,big:=&ListNode{Val: 0},&ListNode{Val: 0}
	pSmall,pBig:=small,big
	for head!=nil{
		if head.Val<x{
			pSmall.Next = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
			pSmall = pSmall.Next
		}else{
			pBig.Next =  &ListNode{
				Val:  head.Val,
				Next: nil,
			}
			pBig = pBig.Next
		}
		head=head.Next
	}
	if small.Next==nil{
		return big.Next
	}
	pSmall.Next = big.Next

	return small.Next
}
