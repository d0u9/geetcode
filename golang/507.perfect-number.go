/*
 * @lc app=leetcode id=507 lang=golang
 *
 * [507] Perfect Number
 */

// @lc code=start
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	iter := int(math.Sqrt(float64(num)))

	for i := 2; i <= iter; i++ {
		if num % i == 0 {
			sum = sum + i + (num / i)
		}
	}

	return sum == num
}
// @lc code=end

