/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

// @lc code=start
func solution1(nums1 []int, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)
	ret := []int{}
	for i := 0; i < l1; i++ {
		j, flag := 0, 0
		for j = 0; j < l2; j++ {
			if flag == 1 {
				if nums2[j] > nums1[i] {
					ret = append(ret, nums2[j])
					break
				}
			}
			if nums2[j] == nums1[i] {
				flag = 1
			}
		}
		if j == l2 {
			ret = append(ret, -1)
		}
	}
	return ret
}

type R struct {
	max int
	id int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solution2(nums1 []int, nums2 []int) []int {
	ret := []int{}
	l1, l2 := len(nums1), len(nums2)
	mp := make(map[int]R)

	for max, i := math.MinInt32, l2 - 1; i >= 0; i-- {
		v := nums2[i]
		mp[v] = R { max: max, id: i, }
		max = Max(max, v)
	}

	//fmt.Println(mp)

	for i := 0; i < l1; i++ {
		v := nums1[i]
		R := mp[v]
		if R.max <= v {
			ret = append(ret, -1)
			continue
		}

		for j := R.id; nums2[j] <= R.max; j++ {
			if nums2[j] > v {
				ret = append(ret, nums2[j])
				break
			}
		}
	}

	return ret
}

func solution3(nums1 []int, nums2 []int) []int {
	ret := []int{}
	l1, l2 := len(nums1), len(nums2)
	mp := make(map[int]int)
	stk := []int{}
	sl := 0

	for i := 0; i < l2; i++ {
		for sl > 0 && nums2[i] > stk[sl - 1] {
			top := stk[sl - 1]
			mp[top] = nums2[i]
			stk = stk[:sl - 1]
			sl--
		}

		stk = append(stk, nums2[i])
		sl++
	}

	for i := 0; i < l1; i++ {
		if v, ok := mp[nums1[i]]; ok {
			ret = append(ret, v)
		} else {
			ret = append(ret, -1)
		}
	}

	return ret
}

func solution4(nums1 []int, nums2 []int) []int {
	ret := []int{}
	_, l2 := len(nums1), len(nums2)
	mp := make([]int, l2, l2)
	stk := []int{}
	sl := 0

	for i := l2 - 1; i >= 0; i-- {
		for sl > 0 && nums2[i] > stk[sl - 1] {
			stk = stk[:sl - 1]
			sl--
		}

		if sl > 0 {
			mp[i] = stk[sl - 1]
		} else {
			mp[i] = -1
		}
		stk = append(stk, nums2[i])
		sl++
	}

	return ret
}


func nextGreaterElement(nums1 []int, nums2 []int) []int {
    return solution3(nums1, nums2)
}
// @lc code=end

