/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

// @lc code=start

func toArray(num int) []int {
	a := make([]int, 0)
	for num != 0 {
		r := num % 10
		num = num / 10
		a = append(a, r)
	}
	return a
}

func solution1(num int) int {
	for num > 9 {
		a := toArray(num)
		num = 0
		for _, v := range a {
			num = num + v
		}
	}
	return num
}

func addDigits(num int) int {
    return solution1(num)
}
// @lc code=end

