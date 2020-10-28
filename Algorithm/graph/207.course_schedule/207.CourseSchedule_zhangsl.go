// Source : https://leetcode.com/problems/course-schedule/
// Author : zhangsl
// Date   : 2020-10-28
package main

/*****************************************************************************************************
 *
 * There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.
 *
 * Some courses may have prerequisites, for example to take course 0 you have to first take course 1,
 * which is expressed as a pair: [0,1]
 *
 * Given the total number of courses and a list of prerequisite pairs, is it possible for you to
 * finish all courses?
 *
 * Example 1:
 *
 * Input: numCourses = 2, prerequisites = [[1,0]]
 * Output: true
 * Explanation: There are a total of 2 courses to take.
 *              To take course 1 you should have finished course 0. So it is possible.
 *
 * Example 2:
 *
 * Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
 * Output: false
 * Explanation: There are a total of 2 courses to take.
 *              To take course 1 you should have finished course 0, and to take course 0 you should
 *              also have finished course 1. So it is impossible.
 *
 * Constraints:
 *
 * 	The input prerequisites is a graph represented by a list of edges, not adjacency matrices.
 * Read more about how a graph is represented.
 * 	You may assume that there are no duplicate edges in the input prerequisites.
 * 	1 <= numCourses <= 10^5
 ******************************************************************************************************/
// 思路 判断有没有环，有环则无法完成所有课程得学习
func canFinish(numCourses int, prerequisites [][]int) bool {
	//	构造邻接表 建立关联关系
	//	hash mark 标记已修课程
	var (
		mp     = map[int][]int{} // 邻接表 记录图
		status = map[int]int{}   // 课程的状态
		dfs    func(int) bool
	)
	//	创建邻接表
	for _, v := range prerequisites {
		mp[v[0]] = append(mp[v[0]], v[1])
	}
	// 课程搜索
	dfs = func(now int) bool {
		if status[now] == 1 {
			return false
		}
		status[now] = 1
		mark := true
		for _, v := range mp[now] {
			mark = mark && dfs(v)
			if !mark {
				return false
			}
		}
		status[now] = 2
		return true
	}
	for i := 0; i < numCourses; i++ {
		if status[i] != 2 {
			if !dfs(i) {
				return false
			}
		}
	}
	return true
}
