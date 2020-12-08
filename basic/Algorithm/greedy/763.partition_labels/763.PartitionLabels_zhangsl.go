// Source : https://leetcode.com/problems/partition-labels/
// Author : zhangsl
// Date   : 2020-11-09
package main
/***************************************************************************************************** 
 *
 * A string S of lowercase English letters is given. We want to partition this string into as many 
 * parts as possible so that each letter appears in at most one part, and return a list of integers 
 * representing the size of these parts.
 * 
 * Example 1:
 * 
 * Input: S = "ababcbacadefegdehijhklij"
 * Output: [9,7,8]
 * Explanation:
 * The partition is "ababcbaca", "defegde", "hijhklij".
 * This is a partition so that each letter appears in at most one part.
 * A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
 * 
 * Note:
 * 
 * 	S will have length in range [1, 500].
 * 	S will consist of lowercase English letters ('a' to 'z') only.
 * 
 ******************************************************************************************************/

func partitionLabels(S string) []int {
	maxPos:=map[byte]int{} //记录每个字母在字符串中出现的最大位置
	for i:=0;i<len(S);i++{
		maxPos[S[i]] = i
	}

	ans :=[]int{} // 记录结果
	start:=0      //用于记录当前扫描区间的起点
	scannedMaxPos:=0
	for i:=0;i<len(S);i++{
		curMaxPos:=maxPos[S[i]] // 当前扫描字符的最大出现位置
		if curMaxPos > scannedMaxPos{
			scannedMaxPos = curMaxPos
		}
		if i== scannedMaxPos{
			ans = append(ans, i-start+1)
			start = i+1
		}
	}
	return ans

}

