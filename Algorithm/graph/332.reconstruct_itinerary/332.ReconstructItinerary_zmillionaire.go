// Source : https://leetcode.com/problems/reconstruct-itinerary/
// Author : zmillionaire
// Date   : 2020-10-28
package main

import "sort"

/*****************************************************************************************************
 *
 * Given a list of airline tickets represented by pairs of departure and arrival airports [from, to],
 * reconstruct the itinerary in order. All of the tickets belong to a man who departs from JFK. Thus,
 * the itinerary must begin with JFK.
 *
 * Note:
 *
 * 	If there are multiple valid itineraries, you should return the itinerary that has the
 * smallest lexical order when read as a single string. For example, the itinerary ["JFK", "LGA"] has
 * a smaller lexical order than ["JFK", "LGB"].
 * 	All airports are represented by three capital letters (IATA code).
 * 	You may assume all tickets form at least one valid itinerary.
 * 	One must use all the tickets once and only once.
 *
 * Example 1:
 *
 * Input: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
 * Output: ["JFK", "MUC", "LHR", "SFO", "SJC"]
 *
 * Example 2:
 *
 * Input: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
 * Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
 * Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"].
 *              But it is larger in lexical order.
 ******************************************************************************************************/

// 知识前置：欧拉通路(路径通过图中所有的边一次且经过所有节点的<通路>称为<欧拉通路>)
// 欧拉回路：路径通过所有的边一次且经过所有节点的<回路>称为<欧拉回路>
// 欧拉图：具有欧拉回路的无向图称为欧拉图
// 半欧拉图：具有欧拉通路但不具有欧拉回路的无向图称为半欧拉图。
// 思路 本题要求 欧拉通路的字符序最小路径
// 使用欧拉通路的Hierholzer算法
func findItinerary(tickets [][]string) []string {
	var (
		mp  = make(map[string][]string) // 记录边
		dfs func(tar string)            // 递归函数
		ans []string
	)
	for _, v := range tickets { // 记录所有边
		mp[v[0]] = append(mp[v[0]], v[1])

	}
	//	排序
	for k := range mp {
		sort.Strings(mp[k])
	}
	dfs = func(tar string) {
		for {
			if len(mp[tar]) < 1 {
				break
			}
			newTar := mp[tar][0]
			mp[tar] = mp[tar][1:]
			dfs(newTar)
		}
		ans = append(ans, tar)

	}
	dfs("JFK")
	// 逆序
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return ans
}
