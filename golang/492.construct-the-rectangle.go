/*
 * @lc app=leetcode id=492 lang=golang
 *
 * [492] Construct the Rectangle
 */

// @lc code=start
func constructRectangle(area int) []int {
	if area == 0 {
		return []int{}
	}
	
	w := int(math.Sqrt(float64(area)))
	l := int(area / w)
	for area % w != 0 {
		w--
		l = int(area / w)
	}
	return []int{ l, w, }
}
// @lc code=end

