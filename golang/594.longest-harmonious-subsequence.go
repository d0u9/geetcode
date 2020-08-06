/*
 * @lc app=leetcode id=594 lang=golang
 *
 * [594] Longest Harmonious Subsequence
 */

// @lc code=start
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solution1(nums []int) int {
	ln := len(nums)
	max := 0
	for i := 0; i < ln; i++ {
		up_c, dn_c, self_c := 0, 0, 0
		for j := i; j < ln; j++ {
			if nums[j] - nums[i] == 1 {
				dn_c++
			}
			if nums[i] - nums[j] == 1 {
				up_c++
			}
			if nums[i] == nums[j] {
				self_c++
			}
		}
		if dn_c != 0 {
			dn_c = dn_c + self_c
		}
		if up_c != 0 {
			up_c = up_c + self_c
		}
		max = MaxInt(max, MaxInt(dn_c, up_c))
	}

	return max
}

func solution2(nums []int) int {
	//ln := len(nums)
	max := 0
	mp := make(map[int]int)
	for _, v := range(nums) {
		mp[v]++
	}

	for _, v := range(nums) {
		//fmt.Println("----", v)
		up, dn := v + 1, v - 1
		up_c, dn_c := 0, 0
		if c, ok := mp[up]; ok {
			up_c = c + mp[v]
			//fmt.Printf("up_c = %d + %d\n", v, mp[v])
		}
		if c, ok := mp[dn]; ok {
			dn_c = c + mp[v]
			//fmt.Printf("dn_c = %d + %d\n", v, mp[v])
		}
		mp[v]--
		//fmt.Println(mp)
		max = MaxInt(max, MaxInt(up_c, dn_c))
	}

	return max
}

func solution3(nums []int) int {
	max := 0
	return max
}

func findLHS(nums []int) int {
    return solution3(nums)
}
// @lc code=end

