/*
 * @lc app=leetcode id=504 lang=golang
 *
 * [504] Base 7
 */

// @lc code=start
func convertToBase7(num int) string {
	ret := []byte{}

	if num == 0 {
		return "0"
	}

	negtaive := false
	if num < 0 {
		negtaive = true
		num = 0 - num
	}
	
	for num != 0 {
		ret = append(ret, byte((num % 7) + '0'))
		num = int(num / 7)
	}

	if negtaive {
		ret = append(ret, '-')
	}

	l := len(ret)
	for i := 0; i < l / 2; i++ {
		ret[i], ret[l - i - 1] = ret[l - i - 1], ret[i]
	}


	return string(ret)
}
// @lc code=end

