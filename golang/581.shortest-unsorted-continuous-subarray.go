/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 */

// @lc code=start
func solution1(nums []int) int {
	l := len(nums)
	sted := make([]int, l, l)
	copy(sted, nums)
	sort.Ints(sted)

	fst, lst := -1, -1
	for i := 0; i < l; i++ {
		if nums[i] != sted[i] {
			if fst == -1 {
				fst = i
			}
			lst = i
		}
	}

	if fst == -1 {
		return 0
	}
	return lst - fst + 1
}

func solution2(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return 0
	}

	l, r := -1, -1
	// check slope from left to right
	for i := 1; i < ln; i++ {
		if nums[i] < nums[i - 1] {
			l = i - 1
			break
		}
	}

	// check slope from right to left
	for i := ln - 1; i > 0; i-- {
		if nums[i] < nums[i - 1] {
			r = i
			break
		}
	}

	if l == -1 && r == -1 {
		return 0
	}

	lt, rt := 0, ln - 1
	for i := r - 1; i >= 0; i-- {
		if nums[i] <= nums[r] {
			lt = i + 1
			break
		}
	}

	for i := l + 1; i < ln; i++ {
		if nums[i] >= nums[l] {
			rt = i - 1
			break
		}
	}

	fmt.Println(l, r)
	fmt.Println(lt, rt)

	if lt < l {
		l = lt
	}

	if rt > r {
		r = rt
	}


	fmt.Printf("l %d, r %d, lt %d, rt %d\n", l, r, lt, rt)
	return r - l + 1
}

func findUnsortedSubarray(nums []int) int {
    return solution2(nums) 
}
// @lc code=end

