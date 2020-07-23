/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func solution1(nums1, nums2 []int) []int {
	mp := make(map[int]bool)
	ret := []int{}
	l1, l2 := len(nums1), len(nums2)
	ls, ns, ll, nl := l1, &nums1, l2, &nums2
	if l2 < l1 {
		ls, ns, ll, nl = l2, &nums2, l1, &nums1
	}

	for i := 0; i < ls; i++ {
		k := (*ns)[i]
		if _, ok := mp[k]; !ok {
			mp[k] = true
		}
	}

	//fmt.Println(mp)

	for i := 0; i < ll; i++ {
		k := (*nl)[i]
		if _, ok := mp[k]; ok {
			mp[k] = false
		}
	}

	//fmt.Println(mp)

	for k, v := range mp {
		if v == false {
			ret = append(ret, k)
		}
	}
	return ret
}

func intersection(nums1 []int, nums2 []int) []int {
	return solution1(nums1, nums2)
}
// @lc code=end

