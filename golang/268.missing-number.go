/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func solution1(nums []int) int {
	var z, t int = 0, 0
	l := len(nums)
	t = l * (l + 1) / 2
	for i := 0; i < l; i ++ {
		z = z + (nums[i])
	}
	return (t - z)
}

func solution2(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums) - 1
	for l < r {
		c := (l + r) / 2
		if nums[c] != c {
			r = c - 1
		} else {
			l = c + 1
		}
	}
	fmt.Println(nums[l], l)
	if nums[l] != l {
		return l
	}
	return l + 1
}

/*
XOR has certain properties
Assume a1 ^ a2 ^ a3 ^ ?^ an = a and a1 ^ a2 ^ a3 ^ ?^ an-1 = b
Then a ^ b = an
*/
func solution3(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		a = a ^ nums[i]
		b = b ^ (i + 1)
	}
	return a ^ b
}

func missingNumber(nums []int) int {
	return solution2(nums)
}
// @lc code=end

