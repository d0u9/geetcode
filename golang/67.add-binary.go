/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	c, max, min, i, j, k := byte(0), 0, 0, 0, 0, 0
	var maxs, mins *string
	if len(a) > len(b) {
		max, min = len(a), len(b)
		maxs, mins = &a, &b
	} else {
		max, min = len(b), len(a)
		maxs, mins = &b, &a
	}

	ret := make([]byte, max + 1)
	i, j, k = max - 1, min - 1, max
    for ; j >= 0; i, j, k = i-1, j-1, k-1 {
        sum := ((*maxs)[i] - '0') + ((*mins)[j] - '0') + c
        ret[k] = (sum % 2) + '0'
        c = sum / 2
    }
    for ; i >= 0; i, k = i-1, k-1 {
        sum := ((*maxs)[i] - '0') + c
        ret[k] = (sum % 2) + '0'
        c = sum / 2
    }

    if c > 0 {
        ret[k] = '1'
    } else {
        ret = ret[1:]
    }

	return string(ret)
}
// @lc code=end

