/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func reverse(nums []int) {
	l := len(nums)
	for i, j := 0, l - 1; i < l / 2; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func solution1(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func solution2(nums []int, k int) {
	l := len(nums)
	ret := make([]int, 0)
	ret = append(ret, nums[l-k:]...)
	fmt.Println(ret)
	ret = append(ret, nums[:l-k]...)
	fmt.Println(ret)
	copy(nums, ret)
}

func solution3(nums []int, k int) {
	l, tmp := len(nums), nums[0]
	start, c := 0, 0
	for i := 0; i < l; i++ {
		n := (c + k) % l
		tmp, nums[n] = nums[n], tmp
		c = n
		if n == start {
			start++
			c = start
			if c < l {
				tmp = nums[c]
			}
		}
	}
}

func rotate(nums []int, k int)  {
	k = k % len(nums)
    solution3(nums, k)
}
// @lc code=end

