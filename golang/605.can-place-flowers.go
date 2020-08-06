/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 */

// @lc code=start

func test(flowerbed []int, plot int) bool {
	if flowerbed[plot] == 1 {
		return false
	}

	ln := len(flowerbed)
	ret := true
	if plot + 1 <= ln - 1 && flowerbed[plot + 1] != 0 {
		ret = false
	}

	if plot - 1 >= 0 && flowerbed[plot - 1] != 0 {
		ret = false
	}

	return ret
}

func solution1(flowerbed []int, n int) bool {
	ln := len(flowerbed)
	if n > ln {
		return false
	}

	cnt := 0
	for i := 0; i < ln; i++ {
		tst := test(flowerbed, i)
		//fmt.Println(i, tst) 
		if tst {
			flowerbed[i] = 1
			cnt++
		}
	}

	return cnt >= n
}

func solution2(flowerbed []int, n int) bool {
	ln := len(flowerbed)
	if ln <= 0 || n > ln {
		return false
	}

	tillable := 1 // 0: not ok, 1: ok
	cnt := 0
	for i := 0; i < ln; i++ {
		cur := flowerbed[i]
		if cur == 1 {
			tillable = 0
		} else if cur == 0 {
			if tillable == 1 {
				if i + 1 == ln || (i + 1 < ln && flowerbed[i + 1] == 0) {
					flowerbed[i] = 1
					tillable = 0
					cnt++
				}
			} else {
				tillable = 1
			}
		}
	}
	return cnt >= n
}

func canPlaceFlowers(flowerbed []int, n int) bool {
    return solution2(flowerbed, n)
}
// @lc code=end

