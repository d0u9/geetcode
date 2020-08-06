/*
 * @lc app=leetcode id=628 lang=golang
 *
 * [628] Maximum Product of Three Numbers
 */

// @lc code=start
func solution1(nums []int) int {
	ln := len(nums)
	if ln < 3 {
		return 0
	}

	sort.Ints(nums)
	mp := make(map[int]int)
	mp[0] = nums[0]
	mp[1] = nums[1]
	mp[2] = nums[2]
	mp[ln-1] = nums[ln-1]
	mp[ln-2] = nums[ln-2]
	mp[ln-3] = nums[ln-3]
	d := []int{}
	for _, v := range(mp) {
		d = append(d, v)
	}
	lx := len(d)
	//fmt.Println(d)
	max := math.MinInt32
	for i := 0; i < lx; i++ {
		for j := i+1; j < lx; j++ {
			for k := j+1; k < lx; k++ {
				prod := d[i] * d[j] * d[k]
				if prod > max {
					max = prod
				}
			}
		}
	}
	return max
}

func solution2(nums []int) int {
	ln := len(nums)
	if ln < 3 {
		return 0
	}

	sort.Ints(nums)
	
	a := nums[ln-1] * nums[ln-2] * nums[ln-3]
	b := nums[0] * nums[1] * nums[ln-1]
	if a > b {
		return a
	}
	return b
}

func solution3(nums []int) int {
	ln := len(nums)
	if ln < 3 {
		return 0
	}

	// max1 < max2 < max3
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	// min1 < min2
	min1, min2 := math.MaxInt32, math.MaxInt32
	for _, v := range(nums) {
		if v > max3 {
			max1, max2, max3 = max2, max3, v
		} else if v > max2 {
			max1, max2 = max2, v
		} else if v > max1 {
			max1 = v
		}

		if v < min1 {
			min1, min2 = v, min1
		} else if v < min2 {
			min2 = v
		}
	}

	a := max1 * max2 * max3
	b := min1 * min2 * max3
	if a > b {
		return a
	}
	return b
}

func maximumProduct(nums []int) int {
    return solution3(nums)
}
// @lc code=end

