/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=start
func maxProfit(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		t := prices[i] - prices[i - 1]
		if t <= 0 {
		} else {
			sum = sum + t
		}
	}
	return sum
}
// @lc code=end

