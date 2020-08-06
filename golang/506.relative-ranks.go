/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 */

// @lc code=start
func binsrh(t int, list []int) int {
	l, r := 0, len(list) - 1
	for l <= r {
		mid := (l + r) / 2
		if list[mid] < t {
			r = mid - 1
		} else if list[mid] > t {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func solution1(nums []int) []string {
	l := len(nums)
	tmp := make([]int, l, l)
	ret := []string{}
	copy(tmp, nums)

	sort.Slice(tmp, func(i, j int)bool {
		return tmp[i] > tmp[j]
	})

	for i := 0; i < l; i++ {
		t := binsrh(nums[i], tmp)
		if t == 0 {
			ret = append(ret, "Gold Medal")
		} else if t == 1 {
			ret = append(ret, "Silver Medal")
		} else if t == 2 {
			ret = append(ret, "Bronze Medal")
		} else {
			ret = append(ret, strconv.Itoa(t + 1))
		}
	}

	return ret
}

func solution2(nums []int) []string {
	ret := []string{}
	l := len(nums)
	mp := make([]int, l, l)

	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				mp[j]++
			} else {
				mp[i]++
			}
		}
	}

	//fmt.Println(mp)
	for i := 0; i < l; i++ {
		t := mp[i]
		if t == 0 {
			ret = append(ret, "Gold Medal")
		} else if t == 1 {
			ret = append(ret, "Silver Medal")
		} else if t == 2 {
			ret = append(ret, "Bronze Medal")
		} else {
			ret = append(ret, strconv.Itoa(t + 1))
		}
	}

	return ret
}

func solution3(nums []int) []string {
	l := len(nums)
	tmp := make([]int, l, l)
	ret := []string{}
	copy(tmp, nums)
	mp := make(map[int]int)

	sort.Slice(tmp, func(i, j int)bool {
		return tmp[i] > tmp[j]
	})

	for i, v := range(tmp) {
		mp[v] = i
	}

	for i := 0; i < l; i++ {
		v := mp[nums[i]]
		if v == 0 {
			ret = append(ret, "Gold Medal")
		} else if v == 1 {
			ret = append(ret, "Silver Medal")
		} else if v == 2 {
			ret = append(ret, "Bronze Medal")
		} else {
			ret = append(ret, strconv.Itoa(v + 1))
		}
	}

	return ret
}

func findRelativeRanks(nums []int) []string {
    return solution1(nums)
}
// @lc code=end

