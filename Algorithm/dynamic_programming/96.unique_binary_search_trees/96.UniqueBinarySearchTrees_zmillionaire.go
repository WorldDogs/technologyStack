// Source : https://leetcode.com/problems/unique-binary-search-trees/
// Author : zmillionaire
// Date   : 2020-09-03
package main
/***************************************************************************************************** 
 *
 * Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
 * 
 * Example:
 * 
 * Input: 3
 * Output: 5
 * Explanation:
 * Given n = 3, there are a total of 5 unique BST's:
 * 
 *    1         3     3      2      1
 *     \       /     /      / \      \
 *      3     2     1      1   3      2
 *     /     /       \                 \
 *    2     1         2                 3
 * 
 * Constraints:
 * 
 * 	1 <= n <= 19
 ******************************************************************************************************/

// 思路：
// 求n个节点的二叉搜索树的个数，令f(i)表示节点数为i时，有f(i)种二叉搜索树，令t(i)表示，当根节点为i时候，有t(i)种二叉搜索树
// f(n) = t(1)+t(2)...+t(n)
// t(i) = f(i-1)*f(n-i)
// f(n) = f(0)*f(n-1)+f(1)*f(n-2)...+f(n-1)*f(0) =>科特兰数

// 动态规划三要素：，
// dp[i]含义：        dp[i] 代表当有i个节点时，有dp[i]种二叉搜索树
// 状态转移方程：      dp[i] = dp[0]*dp[i-1]...+dp[i-1]*dp[0]   当前状态是由前n-1个状态确定
// 初始条件：         当dp[0]=0,dp[1]=1,dp[2]=2
// 复杂度

// 代码排行
// todo:@zhangsl code
func main() {

}