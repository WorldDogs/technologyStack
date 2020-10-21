// Source : https://leetcode.com/problems/queue-reconstruction-by-height/
// Author : zhangsl
// Date   : 2020-10-21
package main

import "sort"

/*****************************************************************************************************
 *
 * Suppose you have a random list of people standing in a queue. Each person is described by a pair of
 * integers (h, k), where h is the height of the person and k is the number of people in front of this
 * person who have a height greater than or equal to h. Write an algorithm to reconstruct the queue.
 *
 * Note:
 * The number of people is less than 1,100.
 *
 * Example
 *
 * Input:
 * [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
 *
 * Output:
 * [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
 *
 ******************************************************************************************************/
// 思路：
// 如果 h1>h2且h2在h1前面， h2对于h1是不可见的
// 按照
func reconstructQueue(input [][]int) [][]int {
	sort.Slice(input, func(i, j int) bool {
		return input[i][0] > input[j][0] || (input[i][0] == input[j][0] && input[i][1] < input[j][1])
	})
	for from, p := range input {
		to := p[1]
		copy(input[to+1:from+1], input[to:from])
		input[to] = p
	}
	return input
}
func reconstuctQueue_recode(input [][]int) [][]int {
	sort.Slice(input, func(i, j int) bool {
		if input[i][0] == input[j][0] {
			return input[i][1] < input[j][1]
		}
		return input[i][0] > input[j][0]
	})
	for from, p := range input {
		to := p[1]
		copy(input[to+1:from+1], input[to:from])
		input[to] = p
	}
	return input
}
