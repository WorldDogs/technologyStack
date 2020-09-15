// Source : https://leetcode.com/problems/word-break/
// Author : zhangsl
// Date   : 2020-09-15
package main
/***************************************************************************************************** 
 *
 * Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, 
 * determine if s can be segmented into a space-separated sequence of one or more dictionary words.
 * 
 * Note:
 * 
 * 	The same word in the dictionary may be reused multiple times in the segmentation.
 * 	You may assume the dictionary does not contain duplicate words.
 * 
 * Example 1:
 * 
 * Input: s = "leetcode", wordDict = ["leet", "code"]
 * Output: true
 * Explanation: Return true because "leetcode" can be segmented as "leet code".
 * 
 * Example 2:
 * 
 * Input: s = "applepenapple", wordDict = ["apple", "pen"]
 * Output: true
 * Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
 *              Note that you are allowed to reuse a dictionary word.
 * 
 * Example 3:
 * 
 * Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * Output: false
 * 
 ******************************************************************************************************/

// 动态规划
// 思路
// 对于0-i这个subStr，j∈[0,i),将subStr分成了两部分，如果前部分符号要求只需要判断后部分。
// 动态规划三步骤：
// 1.确定 dp[i]含义: subStr = str[0,i), 从0-i的串是否满足题目要求
// 2.状态转移方程: dp[i] = dp[j] && check(str[j+1:i+1])
// 3. basecase 初始条件 dp[0] = true
func wordBreak(s string, wordDict []string) bool {
	mark:=make(map[string]bool)
	for _,v:=range wordDict{
		mark[v] = true
	}
	dp:=make([]bool,len(s))
	dp[0] = true
	for i:=1;i<len(s);i++{
		for j:=
	}

}