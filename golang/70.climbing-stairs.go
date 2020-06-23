/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start

// 思路：递归，关键在于上一个解的结果最后有几个1，几个2。
// 上一级最后一步若为1，则可以拆分为1 + 1 或者 2
// 上一级最后一步若为2，则不可拆分只能为1
func climbStairs(n int) int {
    if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	num1, num2 := 1, 1
	for i := 3; i <= n; i++ {
		t := num2
		num2 = num1
		num1 = t + num1
	}

	return num1 + num2
}
// @lc code=end

