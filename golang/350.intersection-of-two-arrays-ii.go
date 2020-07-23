/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start

func intersect_sorted(nums1, nums2 []int) []int {
	// if sorted
	sort.Ints(nums1)
	sort.Ints(nums2)

	l1, l2 := len(nums1), len(nums2)
	ls, ns, ll, nl := l1, &nums1, l2, &nums2
	if l2 < l1 {
		ls, ns, ll, nl = l2, &nums2, l1, &nums1
	}

	ret := []int{}
	for i, j := 0, 0; i < ll && j < ls; {
		vl, vs := (*nl)[i], (*ns)[j]
		if vl > vs {
			j++
		} else if vl < vs {
			i++
		} else {
			ret = append(ret, vl)
			i, j = i + 1, j + 1
		}
	}
	return ret
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solution1(nums1, nums2 []int) []int {
	mp := make(map[int]int)
	mp2 := make(map[int]int)
	l1, l2 := len(nums1), len(nums2)
	ls, ns, ll, nl := l1, &nums1, l2, &nums2
	if l2 < l1 {
		ls, ns, ll, nl = l2, &nums2, l1, &nums1
	}

	for i := 0; i < ls; i++ {
		k := (*ns)[i]
		if v, ok := mp[k]; ok {
			mp[k] = v + 1
		} else {
			mp[k] = 1
			mp2[k] = 0
		}
	}

	for i := 0; i < ll; i++ {
		k := (*nl)[i]
		if _, ok := mp[k]; ok {
			mp2[k] = mp2[k] + 1
		}
	}
	//fmt.Println(mp)
	//fmt.Println(mp2)

	ret := []int{}
	for k := range(mp) {
		rep := minInt(mp[k], mp2[k])
		if rep != 0 {
			for j := 0; j < rep; j++ {
				ret = append(ret, k)
			}
		}
	}

	return ret
}

func solution2(nums1, nums2 []int) []int {
	mp := make(map[int]int)
	l1, l2 := len(nums1), len(nums2)
	ls, ns, ll, nl := l1, &nums1, l2, &nums2
	if l2 < l1 {
		ls, ns, ll, nl = l2, &nums2, l1, &nums1
	}

	for i := 0; i < ls; i++ {
		k := (*ns)[i]
		if v, ok := mp[k]; ok {
			mp[k] = v + 1
		} else {
			mp[k] = 1
		}
	}

	fmt.Println(mp)
	ret := []int{}
	for i := 0; i < ll; i++ {
		k := (*nl)[i]
		if v, ok := mp[k]; ok && v > 0 {
			ret = append(ret, k)
			mp[k] = v - 1
		}
	}


	return ret
}

func intersect(nums1 []int, nums2 []int) []int {
    //return solution1(nums1, nums2)
    return solution2(nums1, nums2)
    //return intersect_sorted(nums1, nums2)
}
// @lc code=end

