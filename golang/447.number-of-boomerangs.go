/*
 * @lc app=leetcode id=447 lang=golang
 *
 * [447] Number of Boomerangs
 */

// @lc code=start
func solution1(points [][]int) int {
	l := len(points)
	cnt := 0
	for i := 0; i < l; i++ {
		mp := make(map[int]int)
		x, y := points[i][0], points[i][1]
		for j := 0; j < l; j++ {
			p, q := points[j][0], points[j][1]
			d := (x - p) * (x - p) + (y - q) * (y - q)
			if v, ok := mp[d]; ok {
				mp[d] = v + 1
			} else {
				mp[d] = 1
			}
		}

		//fmt.Println(mp)
		for _, v := range(mp) {
			if v < 2 {
				continue
			}
			cnt = cnt + (v - 1) * v
		}
	}
	return cnt
}

func numberOfBoomerangs(points [][]int) int {
	return solution1(points)
}
// @lc code=end

