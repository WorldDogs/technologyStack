// Source : https://leetcode.com/problems/minimum-height-trees/
// Author : zhangsl
// Date   : 2020-11-11
package main

/*****************************************************************************************************
 *
 * A tree is an undirected graph in which any two vertices are connected by exactly one path. In other
 * words, any connected graph without simple cycles is a tree.
 *
 * Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where edges[i] = [ai,
 * bi] indicates that there is an undirected edge between the two nodes ai and bi in the tree, you can
 * choose any node of the tree as the root. When you select a node x as the root, the result tree has
 * height h. Among all possible rooted trees, those with minimum height (i.e. min(h))  are called
 * minimum height trees (MHTs).
 *
 * Return a list of all MHTs' root labels. You can return the answer in any order.
 *
 * The height of a rooted tree is the number of edges on the longest downward path between the root
 * and a leaf.
 *
 * Example 1:
 *
 * Input: n = 4, edges = [[1,0],[1,2],[1,3]]
 * Output: [1]
 * Explanation: As shown, the height of the tree is 1 when the root is the node with label 1 which is
 * the only MHT.
 *
 * Example 2:
 *
 * Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
 * Output: [3,4]
 *
 * Example 3:
 *
 * Input: n = 1, edges = []
 * Output: [0]
 *
 * Example 4:
 *
 * Input: n = 2, edges = [[0,1]]
 * Output: [0,1]
 *
 * Constraints:
 *
 * 	1 <= n <= 2 * 104
 * 	edges.length == n - 1
 * 	0 <= ai, bi < n
 * 	ai != bi
 * 	All the pairs (ai, bi) are distinct.
 * 	The given input is guaranteed to be a tree and there will be no repeated edges.
 ******************************************************************************************************/
// 中心点即为最小高度，因为把高度平分掉了
//

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 { // 如果树只有一个节点，直接返回
		return []int{0}
	}

	//	 创建有向图
	inCome := make([]int, n)  //记录入度
	mp := make(map[int][]int) //记录图
	for _, edge := range edges {
		inCome[edge[0]]++
		inCome[edge[1]]++
		mp[edge[0]] = append(mp[edge[0]], edge[1])
		mp[edge[1]] = append(mp[edge[1]], edge[0])
	}

	// 将叶子节点 入队
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if inCome[i] == 1 {
			queue = append(queue, i)
		}
	}

	//	就像剥洋葱一样 一层一层的将节点剥开，最后一层就是答案

	ans := []int{}
	for len(queue) > 0 {
		ans = []int{}
		size := len(queue)          // size表示这层洋葱有多少节点
		for i := 0; i < size; i++ { //剥去最外层
			top := queue[0]
			queue = queue[1:]
			ans = append(ans, top)
			for _, v := range mp[top] { //把最外层相连的放进去，下次这就是最外层啦
				inCome[v]--
				if inCome[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
	}
	return ans
}
