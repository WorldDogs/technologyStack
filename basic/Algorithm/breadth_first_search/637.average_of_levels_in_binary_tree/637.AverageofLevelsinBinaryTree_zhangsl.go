// Source : https://leetcode.com/problems/average-of-levels-in-binary-tree/
// Author : zhangsl
// Date   : 2020-11-19
package main

/*****************************************************************************************************
 *
 * Given a non-empty binary tree, return the average value of the nodes on each level in the form of
 * an array.
 *
 * Example 1:
 *
 * Input:
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * Output: [3, 14.5, 11]
 * Explanation:
 * The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return
 * [3, 14.5, 11].
 *
 * Note:
 *
 * The range of node's value is in the range of 32-bit signed integer.
 *
 ******************************************************************************************************/
// 思路广度优先搜索 层序遍历

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	next := []*TreeNode{root}
	ans := []float64{}
	for len(next) > 0 {
		sum := 0
		cur := next
		next = nil
		for _, node := range cur {
			sum += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		ans = append(ans, float64(sum)/float64(len(cur)))

	}
	return ans
}
