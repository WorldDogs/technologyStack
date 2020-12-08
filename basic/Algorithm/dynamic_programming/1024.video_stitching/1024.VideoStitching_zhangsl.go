// Source : https://leetcode.com/problems/video-stitching/
// Author : zhangsl
// Date   : 2020-09-16
package main

import "sort"

/*****************************************************************************************************
 *
 * You are given a series of video clips from a sporting event that lasted T seconds.  These video
 * clips can be overlapping with each other and have varied lengths.
 *
 * Each video clip clips[i] is an interval: it starts at time clips[i][0] and ends at time
 * clips[i][1].  We can cut these clips into segments freely: for example, a clip [0, 7] can be cut
 * into segments [0, 1] + [1, 3] + [3, 7].
 *
 * Return the minimum number of clips needed so that we can cut the clips into segments that cover the
 * entire sporting event ([0, T]).  If the task is impossible, return -1.
 *
 * Example 1:
 *
 * Input: clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
 * Output: 3
 * Explanation:
 * We take the clips [0,2], [8,10], [1,9]; a total of 3 clips.
 * Then, we can reconstruct the sporting event as follows:
 * We cut [1,9] into segments [1,2] + [2,8] + [8,9].
 * Now we have segments [0,2] + [2,8] + [8,10] which cover the sporting event [0, 10].
 *
 * Example 2:
 *
 * Input: clips = [[0,1],[1,2]], T = 5
 * Output: -1
 * Explanation:
 * We can't cover [0,5] with only [0,1] and [1,2].
 *
 * Example 3:
 *
 * Input: clips =
 * [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]],
 * T = 9
 * Output: 3
 * Explanation:
 * We can take clips [0,4], [4,7], and [6,9].
 *
 * Example 4:
 *
 * Input: clips = [[0,4],[2,8]], T = 5
 * Output: 2
 * Explanation:
 * Notice you can have extra video after the event ends.
 *
 * Constraints:
 *
 * 	1 <= clips.length <= 100
 * 	0 <= clips[i][0] <= clips[i][1] <= 100
 * 	0 <= T <= 100
 ******************************************************************************************************/
// 思路1：动态规划
// dp[i]：覆盖i时，需要的最小的片段数
// 状态转移方程：dp[i] = min(dp[i],dp[j]+1); 这里省略了排序，需要全部遍历一遍
// 初始条件(baseCase): dp[0] = 0;长度为0，不要要剪辑


// 复杂度分析
//time:O(T*len(clips)) mem:O(T)
//rank
//time:100 mem:19.05


func videoStitching(clips [][]int, T int) int {
	const MAX = 110
	dp := make([]int, T+1)
	//	初始化dp
	for i := 0; i < len(dp); i++ {
		dp[i] = MAX
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	//	basecase
	dp[0] = 0
	for i := 1; i <= T; i++ {
		for j := 0; j < len(clips); j++ {
			if clips[j][0] <= i && clips[j][1] >= i {
				dp[i] = min(dp[i], dp[clips[j][0]]+1)
			}
		}
	}
	if dp[T] == MAX {
		return -1
	}
	return dp[T]
}

// 思路2：贪心
// 将重叠部分切开，
// todo Err
func videoStitchingV2(clips [][]int, T int) int {
//	sort
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0]{
			return clips[i][1]>clips[j][1]
		}
		return clips[i][0]<clips[j][0]
	})
//	贪心遍历
	ans:=0
	now:=clips[0][1]
	for i:=1;i<len(clips);i++{
		if clips[i][0]<now{
			continue
		}
	//	 c存在间隙
		if clips[i-1][1]<clips[i][0]{
			return -1
		}else if clips[i-1][1]==clips[i][0]{
			now = clips[i][1]
			continue
		}else{
			ans++
			now = clips[i][1]
		}

	}
	return ans
}
