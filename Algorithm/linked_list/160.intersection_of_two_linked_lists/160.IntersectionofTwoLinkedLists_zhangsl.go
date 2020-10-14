// Source : https://leetcode.com/problems/intersection-of-two-linked-lists/
// Author : zhangsl
// Date   : 2020-10-14
package main

/*****************************************************************************************************
 *
 * Write a program to find the node at which the intersection of two singly linked lists begins.
 *
 * For example, the following two linked lists:
 *
 * begin to intersect at node c1.
 *
 * Example 1:
 *
 * Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
 * Output: Reference of the node with value = 8
 * Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists
 * intersect). From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as
 * [5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the
 * intersected node in B.
 *
 * Example 2:
 *
 * Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
 * Output: Reference of the node with value = 2
 * Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists
 * intersect). From the head of A, it reads as [1,9,1,2,4]. From the head of B, it reads as [3,2,4].
 * There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in
 * B.
 *
 * Example 3:
 *
 * Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
 * Output: null
 * Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5].
 * Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be
 * arbitrary values.
 * Explanation: The two lists do not intersect, so return null.
 *
 * Notes:
 *
 * 	If the two linked lists have no intersection at all, return null.
 * 	The linked lists must retain their original structure after the function returns.
 * 	You may assume there are no cycles anywhere in the entire linked structure.
 * 	Each value on each linked list is in the range [1, 10^9].
 * 	Your code should preferably run in O(n) time and use only O(1) memory.
 ******************************************************************************************************/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：双指针法
// 假设交点为k，则再A链上k之前长度为x，在B链上K之前的长度为y,K点后共同的长度为Z
// 两个点m,n在A，B上以相同的速度运动，当走到底后，切换到另外一条链上走,则当相遇时候
// f(m) = x+Z+y
// f(n) = y+Z+x
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA==nil || headB==nil{
		return nil
	}
	pa, pb := headA, headB
	for pa != pb  {
		if pa != nil{
			pa = pa.Next
		}else{
			pa = headB
		}
		if pb!=nil{
			pb = pb.Next
		}else{
			pb = headA
		}


	}
	return pa
}
