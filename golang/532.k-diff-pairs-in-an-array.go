/*
 * @lc app=leetcode id=532 lang=golang
 *
 * [532] K-diff Pairs in an Array
 */

// @lc code=start
func solution1(nums []int, k int) int {
	sort.Ints(nums)
	ret := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			if nums[j] - nums[i] == k {
				//fmt.Println(nums[i], nums[j])
				ret++
			} else if nums[j] - nums[i] > k {
				break
			}
		}
	}
	return ret
}

func solution2(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	
	mp := make(map[int]int)
	for _, v := range(nums) {
		mp[v]++
	}

	//fmt.Println(mp)
	ret := 0
	for v, _ := range(mp) {
		t := v + k
		if c, ok := mp[t]; ok {
			if t == v && c <= 1 {
				continue
			}
			//fmt.Println(v, t)
			ret++
		}
	}
	return ret
}

func findPairs(nums []int, k int) int {
    return solution2(nums, k)
}
// @lc code=end

