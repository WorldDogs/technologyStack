// Source : https://leetcode.com/problems/product-of-array-except-self/
// Author : zmillionaire
// Date   : 2020-09-29
package main

/*****************************************************************************************************
 *
 * Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal
 * to the product of all the elements of nums except nums[i].
 *
 * Example:
 *
 * Input:  [1,2,3,4]
 * Output: [24,12,8,6]
 *
 * Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array
 * (including the whole array) fits in a 32 bit integer.
 *
 * Note: Please solve it without division and in O(n).
 *
 * Follow up:
 * Could you solve it with constant space complexity? (The output array does not count as extra space
 * for the purpose of space complexity analysis.)
 ******************************************************************************************************/
// 思路
// 由于不能用除法，所以就用乘法，求出前缀之积和后缀之积的数组，最后用乘法合并
// 可以将空间复杂度合并为O(1)
func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	ans:=make([]int,len(nums))
	prefix[0], suffix[len(nums)-1] = 1, 1
	for i:=1;i<len(nums);i++{
		prefix[i]=prefix[i-1]*nums[i-1]
	}
	for i:=len(nums)-2;i>=0;i--{
		suffix[i]=suffix[i+1]*nums[i+1]
	}
	for i:=0;i<len(nums);i++{
		ans[i]=prefix[i]*suffix[i]
	}
	return ans
}
