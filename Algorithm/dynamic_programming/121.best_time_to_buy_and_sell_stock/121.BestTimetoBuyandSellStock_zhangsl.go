// Source : https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// Author : zhangsl
// Date   : 2020-09-08
package main

import "math"

/*****************************************************************************************************
 *
 * Say you have an array for which the ith element is the price of a given stock on day i.
 * 
 * If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of 
 * the stock), design an algorithm to find the maximum profit.
 * 
 * Note that you cannot sell a stock before you buy one.
 * 
 * Example 1:
 * 
 * Input: [7,1,5,3,6,4]
 * Output: 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
 *              Not 7-1 = 6, as selling price needs to be larger than buying price.
 * 
 * Example 2:
 * 
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 * 
 ******************************************************************************************************/
// 动态规划三步骤
// 此理解有误。

// 确定初始条件
// dp[0] = 0

// 复杂度:
// time: o(n) ;只需要遍历数组一遍
// mem: o(n)  ; 需要开辟和数组相同大小的空间 记录状态,由于只需要dp[i-1] 所以,可以进行状态压缩

func maxProfit(prices []int) int {
	if len(prices)==0{
		return 0
	}

	dp:=make([]int,len(prices))
	dp[0] = 0
	ans := 0
	for i:=1;i<len(prices);i++{
		dp[i] = dp[i-1]+prices[i]-prices[i-1]
		if dp[i]<0{
			dp[i] =0
		}
		if dp[i]>ans{
			ans = dp[i]
		}
	}
	return ans
}


// 换一种思路来思考
// 最终要求最大的利润，需要找到最低点买进，然后在最高点卖出。可以在遍历的时候记录历史最低值，然后和当前值进行比较，更细结果。

//time: 97 mem:100
func maxProfit2(prices []int) int {
	if len(prices)<1{
		return 0
	}
	minVal:=math.MaxInt32
	ans:=0 //利润最小为0，可以当天买进再卖出
	for _,v:=range prices{
		if minVal>v{
			minVal = v
		}else if v - minVal>ans{
			ans = v-minVal
		}
	}
	return ans
}