// Source : https://leetcode.com/problems/two-sum/
// Author : zhangsl
// Date   : 2020-09-30
package main
/***************************************************************************************************** 
 *
 * Given an array of integers nums and an integer target, return indices of the two numbers such that 
 * they add up to target.
 * 
 * You may assume that each input would have exactly one solution, and you may not use the same 
 * element twice.
 * 
 * You can return the answer in any order.
 * 
 * Example 1:
 * 
 * Input: nums = [2,7,11,15], target = 9
 * Output: [0,1]
 * Output: Because nums[0] + nums[1] == 9, we return [0, 1].
 * 
 * Example 2:
 * 
 * Input: nums = [3,2,4], target = 6
 * Output: [1,2]
 * 
 * Example 3:
 * 
 * Input: nums = [3,3], target = 6
 * Output: [0,1]
 * 
 * Constraints:
 * 
 * 	2 <= nums.length <= 105
 * 	-109 <= nums[i] <= 109
 * 	-109 <= target <= 109
 * 	Only one valid answer exists.
 ******************************************************************************************************/
// 思路
func twoSum(nums []int, target int) []int {
	if len(nums)==0{
		return nums
	}
	ans:=[]int{}
	m:=make(map[int]int)
	for i,v:=range nums{
		m[v] = i
	}
	for i,v:=range nums{
		if index,ok:=m[target-v];ok && i!=index{
			ans  = append(ans, index)
			ans = append(ans, i)
			return ans
		}
	}
	return ans
}
// 思路 双指针
// 当为有序数组时，可以使用 双指针
//
// 对于有序数组：twoSum是 nSum的baseCase
// 题目只要求 求出一个答案，此答案是通用解法（如果包含多个解）
func twoSum2(nums []int, target int) [][]int {
	//l,r 分别为指向数组头和尾的指针，left,right为指向的值
	ans:=[][]int{}
	l,r:=0,len(nums)-1
	for l<r{
		left,right:=nums[l],nums[r]
		sum:=left+right
		if sum==target{
			ans = append(ans, []int{l,r}) // 多组答案 修改此处
		}else if sum>target{
			r--
		}else if sum<target{
			l++
		}
	//	过滤相同元素
		for l<r && left==nums[l]{
			l++
		}
		for l<r && right==nums[r]{
			r--
		}
	}
	return ans
}

