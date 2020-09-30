// Source : https://leetcode.com/problems/maximum-subarray/
// Author : zhangsl
// Date   : 2020-09-30
package main

import "math"

/*****************************************************************************************************
 *
 * Given an integer array nums, find the contiguous subarray (containing at least one number) which
 * has the largest sum and return its sum.
 *
 * Follow up: If you have figured out the O(n) solution, try coding another solution using the divide
 * and conquer approach, which is more subtle.
 *
 * Example 1:
 *
 * Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 * Example 2:
 *
 * Input: nums = [1]
 * Output: 1
 *
 * Example 3:
 *
 * Input: nums = [0]
 * Output: 0
 *
 * Example 4:
 *
 * Input: nums = [-1]
 * Output: -1
 *
 * Example 5:
 *
 * Input: nums = [-2147483647]
 * Output: -2147483647
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 2 * 104
 * 	-231 <= nums[i] <= 231 - 1
 ******************************************************************************************************/

// 思路：
// 如果一个区间中的和是小于0的 那会给整个sum造成负增长，所以当sum<0时应该清零
// 整个数组维护一个最大值，ans恰好承担这个使命，它是区间和与每一个数字的综合最大值。

func maxSubArray(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return x
	}
	var (
		ans = math.MinInt32
		sum = 0
	)
	for i:=0;i<len(nums);i++{
		sum+=nums[i]
		ans = max(sum,ans)
		if sum<0{
			sum=0
		}
	}
	return ans
}
