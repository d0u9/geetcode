/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	
	minVal, profit := prices[0], 0

	for i := 0; i < l; i++ {
		p := prices[i] - minVal
		if p > profit {
			profit = p
		}
		if prices[i] < minVal {
			minVal = prices[i]
		}
	}

	return profit  
}
// @lc code=end

