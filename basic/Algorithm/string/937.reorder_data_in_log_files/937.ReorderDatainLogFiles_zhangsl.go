// Source : https://leetcode.com/problems/reorder-data-in-log-files/
// Author : zhangsl
// Date   : 2020-11-10
package main

import (
	"sort"
	"strings"
	"unicode"
)

/*****************************************************************************************************
 *
 * You have an array of logs.  Each log is a space delimited string of words.
 *
 * For each log, the first word in each log is an alphanumeric identifier.  Then, either:
 *
 * 	Each word after the identifier will consist only of lowercase letters, or;
 * 	Each word after the identifier will consist only of digits.
 *
 * We will call these two varieties of logs letter-logs and digit-logs.  It is guaranteed that each
 * log has at least one word after its identifier.
 *
 * Reorder the logs so that all of the letter-logs come before any digit-log.  The letter-logs are
 * ordered lexicographically ignoring identifier, with the identifier used in case of ties.  The
 * digit-logs should be put in their original order.
 *
 * Return the final order of the logs.
 *
 * Example 1:
 * Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
 * Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
 *
 * Constraints:
 *
 * 	0 <= logs.length <= 100
 * 	3 <= logs[i].length <= 100
 * 	logs[i] is guaranteed to have an identifier, and a word after the identifier.
 ******************************************************************************************************/
// 思路：自定义排序的问题
// 一共三种排序规则：
// 字母日志与字母日志
// 字母日志与数字日志
// 数字与数字
func reorderLogFiles(logs []string) []string {

	sort.SliceStable(logs, func(i, j int) bool {
		split1 := strings.SplitN(logs[i], " ", 2)
		split2 := strings.SplitN(logs[j], " ", 2)

		digital1 := unicode.IsDigit(rune(split1[1][0]))
		digital2 := unicode.IsDigit(rune(split2[1][0]))
		if !digital1 && !digital2 { // 字母与字母
			if split1[1] == split2[1] {
				return split1[0] < split2[0]
			}
			return split1[1] < split2[1]
		}
		// 数字与数字
		if digital1 && digital2 {
			return false
		}
		// 字母与数字
		if digital1 {
			return false
		}
		return true
	})
	return logs
}
func main() {
	reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"})
}
