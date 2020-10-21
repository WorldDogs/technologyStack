// Source : https://leetcode.com/problems/task-scheduler/
// Author : zhangsl
// Date   : 2020-10-21
package main

import "sort"

/*****************************************************************************************************
 *
 * Given a characters array tasks, representing the tasks a CPU needs to do, where each letter
 * represents a different task. Tasks could be done in any order. Each task is done in one unit of
 * time. For each unit of time, the CPU could complete either one task or just be idle.
 *
 * However, there is a non-negative integer n that represents the cooldown period between two same
 * tasks (the same letter in the array), that is that there must be at least n units of time between
 * any two same tasks.
 *
 * Return the least number of units of times that the CPU will take to finish all the given tasks.
 *
 * Example 1:
 *
 * Input: tasks = ["A","A","A","B","B","B"], n = 2
 * Output: 8
 * Explanation:
 * A -> B -> idle -> A -> B -> idle -> A -> B
 * There is at least 2 units of time between any two same tasks.
 *
 * Example 2:
 *
 * Input: tasks = ["A","A","A","B","B","B"], n = 0
 * Output: 6
 * Explanation: On this case any permutation of size 6 would work since n = 0.
 * ["A","A","A","B","B","B"]
 * ["A","B","A","B","A","B"]
 * ["B","B","B","A","A","A"]
 * ...
 * And so on.
 *
 * Example 3:
 *
 * Input: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
 * Output: 16
 * Explanation:
 * One possible solution is
 * A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> idle -> idle -> A -> idle -> idle -> A
 *
 * Constraints:
 *
 * 	1 <= task.length <= 104
 * 	tasks[i] is upper-case English letter.
 * 	The integer n is in the range [0, 100].
 ******************************************************************************************************/
// 思路：通过观察，先执行数量多的指令，会使时间更少
// 假设间隔=2； abcabca
//
// public class Solution {
//    public int leastInterval(char[] tasks, int n) {
//        int[] map = new int[26];
//        for (char c: tasks)
//            map[c - 'A']++;
//        Arrays.sort(map);
//        int time = 0;
//        while (map[25] > 0) {
//            int i = 0;
//            while (i <= n) {
//                if (map[25] == 0)
//                    break;
//                if (i < 26 && map[25 - i] > 0)
//                    map[25 - i]--;
//                time++;
//                i++;
//            }
//            Arrays.sort(map);
//        }
//        return time;
//    }
//}
//
func leastInterval(tasks []byte, n int) int {
	mp := make([]int, 26)
	for _, v := range tasks {
		mp[v-'A']++
	}
	sort.Slice(mp, func(i, j int) bool {
		return mp[i] < mp[j]
	})
	time := 0
	for mp[25] > 0 {
		i := 0
		for i <= n {
			if mp[25] == 0 {
				break
			}
			if i < 26 && mp[25-i] > 0 {
				mp[25-i]--
			}
			time++
			i++
		}
		sort.Slice(mp, func(i, j int) bool {
			return mp[i] < mp[j]
		})
	}
	return time
}
