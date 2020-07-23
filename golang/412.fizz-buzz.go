/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func IntToStr(n int) string {
	ret := []byte{}
	for n > 0 {
		r, q := n % 10, n / 10
		ret = append(ret, byte(r + '0'))
		n = q
	}

	for i, l := 0, len(ret); i < l / 2; i++ {
		ret[i], ret[l - i - 1] = ret[l - i - 1], ret[i]
	}

	return string(ret)
}

func fizzBuzz(n int) []string {
	ret := []string{}
	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			ret = append(ret, "FizzBuzz")
		} else if  i % 3 == 0 {
			ret = append(ret, "Fizz")
		} else if i % 5 == 0 {
			ret = append(ret, "Buzz")
		} else {
			//ret = append(ret, IntToStr(i))
			ret = append(ret, strconv.Itoa(i))
		}
	}
	return ret
}
// @lc code=end

