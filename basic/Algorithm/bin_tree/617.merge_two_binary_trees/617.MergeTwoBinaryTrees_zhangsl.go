// Source : https://leetcode.com/problems/merge-two-binary-trees/
// Author : zhangsl
// Date   : 2020-11-19
package main

/*****************************************************************************************************
 *
 * Given two binary trees and imagine that when you put one of them to cover the other, some nodes of
 * the two trees are overlapped while the others are not.
 *
 * You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then
 * sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used
 * as the node of new tree.
 *
 * Example 1:
 *
 * Input:
 * 	Tree 1                     Tree 2
 *           1                         2
 *          / \                       / \
 *         3   2                     1   3
 *        /                           \   \
 *       5                             4   7
 * Output:
 * Merged tree:
 * 	     3
 * 	    / \
 * 	   4   5
 * 	  / \   \
 * 	 5   4   7
 *
 * Note: The merging process must start from the root nodes of both trees.
 ******************************************************************************************************/

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路 将B合并到A上
// 同时存在，只有A，只有B
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var dfs func(*TreeNode, *TreeNode) *TreeNode
	dfs = func(t1 *TreeNode, t2 *TreeNode) *TreeNode {
		if t1 != nil && t2 != nil {
			t1.Val += t2.Val
		} else if t1 == nil && t2 != nil {
			return t2
		} else if t1 != nil && t2 == nil {
			return t1
		} else {
			return nil
		}
		t1.Left = dfs(t1.Left, t2.Left)
		t1.Right = dfs(t1.Right, t2.Right)
		return t1
	}
	return dfs(t1, t2)
}
