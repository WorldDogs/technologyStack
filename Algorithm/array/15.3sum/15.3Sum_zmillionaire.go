// Source : https://leetcode.com/problems/3sum/
// Author : zmillionaire
// Date   : 2020-09-29
package main

import "sort"

/*****************************************************************************************************
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find
 * all unique triplets in the array which gives the sum of zero.
 *
 * Notice that the solution set must not contain duplicate triplets.
 *
 * Example 1:
 * Input: nums = [-1,0,1,2,-1,-4]
 * Output: [[-1,-1,2],[-1,0,1]]
 * Example 2:
 * Input: nums = []
 * Output: []
 * Example 3:
 * Input: nums = [0]
 * Output: []
 *
 * Constraints:
 *
 * 	0 <= nums.length <= 3000
 * 	-105 <= nums[i] <= 105
 ******************************************************************************************************/
// 思路：首先是解决2sum的问题
// 2sum可以使用双指针解决，但是要排除重复情况。
// nSum 通过递归不断缩小问题规模，最后到2Sum
// nSum
func nSumTarget(nums []int, start, n, target int) [][]int {
	sz := len(nums)
	if n < 2 || sz < n {
		return nil
	}
	ans := [][]int{}
	if n == 2 {
		l, r := start, sz-1
		for l < r {
			left, right := nums[l], nums[r]
			sum := left + right
			if sum == target {
				ans = append(ans, []int{left, right})
				for l < r && left == nums[l] {
					l++
				}
				for l < r && right == nums[r] {
					r--
				}
			} else if sum < target {
				for l < r && left == nums[l] {
					l++
				}
			} else if sum > target {
				for l < r && right == nums[r] {
					r--
				}
			}
		}
	} else {
		for i := start; i < sz; i++ {
			sub := nSumTarget(nums, i+1, n-1, target-nums[i])
			for _, v := range sub {
				ans = append(ans, append(v, nums[i]))
			}
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return ans
}
func threeSum(nums []int) [][]int {
	if len(nums) < 1 {
		return nil
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nSumTarget(nums, 0, 3, 0)
}
