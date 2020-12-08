// Source : https://leetcode.com/problems/median-of-two-sorted-arrays/
// Author : zmillionaire
// Date   : 2020-09-29
package main

/*****************************************************************************************************
 *
 * Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two
 * sorted arrays.
 *
 * Follow up: The overall run time complexity should be O(log (m+n)).
 *
 * Example 1:
 *
 * Input: nums1 = [1,3], nums2 = [2]
 * Output: 2.00000
 * Explanation: merged array = [1,2,3] and median is 2.
 *
 * Example 2:
 *
 * Input: nums1 = [1,2], nums2 = [3,4]
 * Output: 2.50000
 * Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 *
 * Example 3:
 *
 * Input: nums1 = [0,0], nums2 = [0,0]
 * Output: 0.00000
 *
 * Example 4:
 *
 * Input: nums1 = [], nums2 = [1]
 * Output: 1.00000
 *
 * Example 5:
 *
 * Input: nums1 = [2], nums2 = []
 * Output: 2.00000
 *
 * Constraints:
 *
 * 	nums1.length == m
 * 	nums2.length == n
 * 	0 <= m <= 1000
 * 	0 <= n <= 1000
 * 	1 <= m + n <= 2000
 * 	-106 <= nums1[i], nums2[i] <= 106
 ******************************************************************************************************/

// 思路：转化成求两个有序数组的第k小
// todo 此题需要recode
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	left := (n + m + 1) / 2
	right := (n + m + 2) / 2
	//将偶数和奇数的情况合并，如果是奇数，会求两次同样的 k 。
	return float64(getKth(nums1, 0, n-1, nums2, 0, m-1, left) + getKth(nums1, 0, n-1, nums2, 0, m-1, right)) * 0.5
}

func getKth(nums1 []int, start1, end1 int, nums2 []int, start2, end2, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	//让 len1 的长度小于 len2，这样就能保证如果有数组空了，一定是 len1
	if len1 > len2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start2+k-1]
	}

	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	i := start1 + min(len1, k/2) - 1
	j := start2 + min(len2, k/2) - 1

	if nums1[i] > nums2[j] {
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	} else {
		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	}
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
