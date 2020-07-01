/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func solution1(nums []int) int {
	m := make(map[int]int)
	l := len(nums)
	for _, n := range(nums) {
		if v, ok := m[n]; ok {
			m[n] = v + 1
		} else {
			m[n] = 1
		}
		if m[n] > int(l / 2) {
			return n
		}
	}
	return 0
}

func solution2(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	c, cur := 1, nums[0]
	for i := 1; i < l; i++ {
		if c > l / 2 {
			return cur
		}
		if cur != nums[i] {
			c, cur = 1, nums[i]
		} else {
			c++
		}
	}
	if c > l / 2 {
		return cur
	}
	return 0
}

// Solution from internet
func solution3(nums []int) int {
	cur, cnt := 0, 0
	for _,v := range nums {
		if cnt == 0 {
			cur = v
			cnt++
		} else {
			if v == cur {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return cur
}

func solution4(nums []int) int {
	l := len(nums)
	sort.Ints(nums)
	return nums[l / 2]
}

func majorityElement(nums []int) int {
	return solution2(nums)
}
// @lc code=end

