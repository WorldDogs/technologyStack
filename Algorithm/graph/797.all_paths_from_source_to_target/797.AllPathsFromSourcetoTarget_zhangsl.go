// Source : https://leetcode.com/problems/all-paths-from-source-to-target/
// Author : zhangsl
// Date   : 2020-10-30
package main

/*****************************************************************************************************
 *
 * Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find all possible paths
 * from node 0 to node n - 1, and return them in any order.
 *
 * The graph is given as follows: graph[i] is a list of all nodes you can visit from node i (i.e.,
 * there is a directed edge from node i to node graph[i][j]).
 *
 * Example 1:
 *
 * Input: graph = [[1,2],[3],[3],[]]
 * Output: [[0,1,3],[0,2,3]]
 * Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
 *
 * Example 2:
 *
 * Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
 * Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
 *
 * Example 3:
 *
 * Input: graph = [[1],[]]
 * Output: [[0,1]]
 *
 * Example 4:
 *
 * Input: graph = [[1,2,3],[2],[3],[]]
 * Output: [[0,1,2,3],[0,2,3],[0,3]]
 *
 * Example 5:
 *
 * Input: graph = [[1,3],[2],[3],[]]
 * Output: [[0,1,2,3],[0,3]]
 *
 * Constraints:
 *
 * 	n == graph.length
 * 	2 <= n <= 15
 * 	0 <= graph[i][j] < n
 * 	graph[i][j] != i (i.e., there will be no self-loops).
 * 	The input graph is guaranteed to be a DAG.
 ******************************************************************************************************/
// 思路dfs 到达target 保存结果，
func allPathsSourceTarget(graph [][]int) [][]int {
	if graph == nil {
		return nil
	}
	var (
		ans    = [][]int{}
		mp     = map[int][]int{}
		dfs    = func([]int, int) {}
		maxTar = len(graph) - 1
	)
	//	构造图
	for i := range graph {
		for j := range graph[i] {
			mp[i] = append(mp[i], graph[i][j])
		}
	}
	//  遍历图 获得结果
	dfs = func(path []int, target int) {
		if target == maxTar {
			tmp:=make([]int,len(path))
			copy(tmp,path)
			ans = append(ans, append(tmp,target))
			return
		}
		for _,v := range mp[target] {
			dfs(append(path, target), v)
		}
	}
	dfs([]int{},0)
	return ans
}
