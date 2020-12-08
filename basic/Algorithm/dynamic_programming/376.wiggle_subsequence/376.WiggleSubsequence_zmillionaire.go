// Source : https://leetcode.com/problems/wiggle-subsequence/
// Author : zmillionaire
// Date   : 2020-09-15
package main
/***************************************************************************************************** 
 *
 * A sequence of numbers is called a wiggle sequence if the differences between successive numbers 
 * strictly alternate between positive and negative. The first difference (if one exists) may be 
 * either positive or negative. A sequence with fewer than two elements is trivially a wiggle sequence.
 * 
 * For example, [1,7,4,9,2,5] is a wiggle sequence because the differences (6,-3,5,-7,3) are 
 * alternately positive and negative. In contrast, [1,4,7,2,5] and [1,7,4,5,5] are not wiggle 
 * sequences, the first because its first two differences are positive and the second because its last 
 * difference is zero.
 * 
 * Given a sequence of integers, return the length of the longest subsequence that is a wiggle 
 * sequence. A subsequence is obtained by deleting some number of elements (eventually, also zero) 
 * from the original sequence, leaving the remaining elements in their original order.
 * 
 * Example 1:
 * 
 * Input: [1,7,4,9,2,5]
 * Output: 6
 * Explanation: The entire sequence is a wiggle sequence.
 * 
 * Example 2:
 * 
 * Input: [1,17,5,10,13,15,10,5,16,8]
 * Output: 7
 * Explanation: There are several subsequences that achieve this length. One is [1,17,10,13,10,16,8].
 * 
 * Example 3:
 * 
 * Input: [1,2,3,4,5,6,7,8,9]
 * Output: 2
 * 
 * Follow up:
 * Can you do it in O(n) time?
 * 
 ******************************************************************************************************/
// 动态规划
// dp[i]：以nums[i]结尾时，摆动的最长长度
// 状态转移方程:
//      nums[j]<nums[i]=> up[i] = max(up[i],down[j]+1) j=[0,i-1]
//      nums[j]>nums[i]=>down[i] = max(down[i],up[j]+1) j=[0,i-1]
// 初始条件 basecase:
//      down[0] = 0,up[0] = 1
//

//复杂度：
//time:o(n^2) mem:o(n)
//ramk
//time:100 mem:13.04
func wiggleMaxLength(nums []int) int {
	nNums:=len(nums)
	if nNums<2{
		return nNums
	}

	up:=make([]int,len(nums))
	down:=make([]int,len(nums))
	max:= func(x,y int) int{
		if x>y{
			return x
		}
		return y
	}
//	初始条件
	up[0],down[0]=1,1
	for i:=1;i<len(nums);i++{
		for j:=0;j<i;j++{
			if nums[j]>nums[i]{
				down[i] = max(down[i],1+up[j])
			}else if nums[j]<nums[i]{
				up[i] = max(up[i],1+down[j])
			}
		}
	}
	return max(max(up[nNums-1],down[nNums-1]),1)
}