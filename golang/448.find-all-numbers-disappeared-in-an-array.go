/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func solution1(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return []int{}
	}
	s, i := 0, nums[0] - 1
	cnt := 0
	for {
		if nums[i] == i + 1 {
			s++
			if s >= l {
				break
			}
			i = nums[s] - 1
			continue
		}

		cnt++

		next := nums[i] - 1
		nums[i] = i + 1
		i = next
	}
	
	ret := []int{}
	for i = 0; i < l; i++ {
		if nums[i] != i + 1 {
			ret = append(ret, i + 1)
		}
	}
	

	return ret
}


func solution2(nums []int) []int {
	ret := []int{}
	l := len(nums)
	if l == 0 {
		return ret
	}
	d := make([]int, l + 1)
	for i := 0; i < l; i++ {
		k := nums[i]
		d[k] = 1
	}

	for i := 1; i < l + 1; i++ {
		if d[i] == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}

func findDisappearedNumbers(nums []int) []int {
	return solution2(nums)
}

// @lc code=end

