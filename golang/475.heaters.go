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

	//sort.Ints(houses)
	//sort.Ints(heaters)
	sort.SliceStable(houses,  func(i, j int) bool {return houses[i] < houses[j]})
	sort.SliceStable(heaters, func(i, j int) bool {return heaters[i] < heaters[j]})

	if l2 > 10000 {
		fmt.Println(l1, l2)
		fmt.Println(heaters[1522], heaters[1521])
	}

	x, y := 0, 1
	l, r := heaters[x], heaters[y]
	radius := 0
	for i := 0; i < l1; {
		d1 := Abs(houses[i] - l)
		d2 := Abs(houses[i] - r)
		//fmt.Printf("l = %d, r = %d\n", l, r)
		//fmt.Printf("d1 = %d, d2 = %d\n", d1, d2)
		if d1 > d2 {
			if i > 1500 {
				fmt.Printf("-- x = %d, y = %d\n", x, y)
			}
			x = y
			y = y + 1
			//x, y = y, y + 1

			if i > 1500 {
				fmt.Printf("-+ x = %d, y = %d, l2 = %d\n", x, y, l2)
			}
			l = heaters[x]
			if y >= l2 {
				r = math.MaxInt32
			} else {
				r = heaters[y]
				if i > 1500 {
					fmt.Println(",.,.,.", heaters[1522], heaters[1521])
				}
			}
			if i > 1500 {
				fmt.Printf("-- %d, %d\n", l, r)
			}
			continue
		}
		if Min(d1, d2) != 0 {
			//fmt.Printf("curr = %d\n", houses[i])
			//fmt.Printf("l = %d, r = %d\n", l, r)
			//fmt.Printf("d1 = %d, d2 = %d\n", d1, d2)
		}
		radius = Max(radius, Min(d1, d2))
		i++
	}

	return radius
}
// @lc code=end

