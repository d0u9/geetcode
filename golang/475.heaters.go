/*
 * @lc app=leetcode id=475 lang=golang
 *
 * [475] Heaters
 */

// @lc code=start
func Abs(x int) int {
	if x > 0 {
		return x
	}

	return (0 - x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findRadius(houses []int, heaters []int) int {
	l1, l2 := len(houses), len(heaters)
	if l1 == 0 || l2 == 0 {
		return 0
	}

	if l2 == 1 {
		return Max(Abs(houses[0] - heaters[0]), Abs(houses[l1-1] - heaters[0]))
	}

	sort.Ints(houses)
	sort.Ints(heaters)

	x, y := 0, 1
	l, r := heaters[x], heaters[y]
	radius := 0
	for i := 0; i < l1; {
		d1 := Abs(houses[i] - l)
		d2 := Abs(houses[i] - r)
		
		if d1 >= d2 {
			x, y = y, y + 1

			l = heaters[x]
			if y >= l2 {
				r = math.MaxInt32
			} else {
				r = heaters[y]
			}
			continue
		}

		radius = Max(radius, Min(d1, d2))
		i++
	}

	return radius
}
// @lc code=end

