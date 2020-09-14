// Source : https://leetcode.com/problems/jian-sheng-zi-lcof/
// Author : zhangsl
// Date   : 2020-09-14
package main

import "math"

/*****************************************************************************************************
 *
 *
 ******************************************************************************************************/

// 思路 动态规划
// 动态规划三步骤：
// 1.确定dp[i] 含义： dp[i]表示将i拆分成 k个整数后形成的最大乘积。
// 2.状态转移方程:
//  枚举将i分解的每一种情况: j,i-j .而可以将j继续拆解,将j再次拆解后就变成了子问题.
// 3.确定初始条件: 绳子能够分割的最小整数: 1. dp[1] = 0


// 复杂度分析
//mem:需要记录n种状态,空间复杂度: O(n)
//time:O(n^2)
// OJ 评测数据
//time:100 mem:38.58
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	max:= func(x,y,z int) int{
		m:=math.MinInt32
		if x>m{
			m = x
		}
		if y>m{
			m=y
		}
		if z>m{
			m = z
		}
		return m
	}
	for i := 2; i <= n; i++ {
		for j:=1;j<i;j++{ // 枚举每一种分割的可能
			dp[i] = max(dp[i],j*dp[i-j],j*(i-j))
		}
	}
	return dp[n]
}
