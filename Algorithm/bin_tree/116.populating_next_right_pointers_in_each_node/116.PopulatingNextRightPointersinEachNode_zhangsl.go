// Source : https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
// Author : zhangsl
// Date   : 2020-08-26
package main

/*****************************************************************************************************
 *
 * You are given a perfect binary tree where all leaves are on the same level, and every parent has
 * two children. The binary tree has the following definition:
 *
 * struct Node {
 *   int val;
 *   Node *left;
 *   Node *right;
 *   Node *next;
 * }
 *
 * Populate each next pointer to point to its next right node. If there is no next right node, the
 * next pointer should be set to NULL.
 *
 * Initially, all next pointers are set to NULL.
 *
 * Follow up:
 *
 * 	You may only use constant extra space.
 * 	Recursive approach is fine, you may assume implicit stack space does not count as extra
 * space for this problem.
 *
 * Example 1:
 *
 * Input: root = [1,2,3,4,5,6,7]
 * Output: [1,#,2,3,#,4,5,6,7,#]
 * Explanation: Given the above perfect binary tree (Figure A), your function should populate each
 * next pointer to point to its next right node, just like in Figure B. The serialized output is in
 * level order as connected by the next pointers, with '#' signifying the end of each level.
 *
 * Constraints:
 *
 * 	The number of nodes in the given tree is less than 4096.
 * 	-1000 <= node.val <= 1000
 ******************************************************************************************************/

//Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 思路1 层序遍历，将每一层连接起来
// time:63 mem:96
func connect(root *Node) *Node {
	//边界检查
	if root==nil{
		return nil
	}
	var (
		que   = []*Node{}
		queue = []*Node{}
	)
//	初始化队列
//
	que = append(que,root)
	for len(que)>0{
	//	处理当前层的每一个元素，并将下一层加入队列
		queue = []*Node{}
		for i:=0;i<len(que);i++{
			if i<len(que)-1{
				que[i].Next = que[i+1]
			}
			que[i].Next = que[i+1]
			if que[i].Left!=nil{
				queue = append(queue, que[i].Left)
			}
			if que[i].Right!=nil{
				queue = append(queue,que[i].Right)
			}
		}
		que = queue
	}
	return root
}

// 思路2 不使用额外空间
func connect2(root *Node) *Node {
	
}
