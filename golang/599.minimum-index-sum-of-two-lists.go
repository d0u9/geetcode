/*
 * @lc app=leetcode id=599 lang=golang
 *
 * [599] Minimum Index Sum of Two Lists
 */

// @lc code=start
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func findRestaurant(list1 []string, list2 []string) []string {
	mp := make(map[string]int)
	for i, s := range(list1) {
		mp[s] = i
	}

	min := math.MaxInt32
	ret := []string{}
	for i, s := range(list2) {
		if id1, ok := mp[s]; ok {
			if  id1 + i < min {
				ret = []string{}
				min = id1 + i
			}
			if id1 + i == min {
				ret = append(ret, s)
			}
		}
	}


	return ret
}
// @lc code=end