/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

package main

import "fmt"

func main() {
    n1 := []int { 1, 2, 3, 0, 0, 0, }
    n2 := []int { 2, 5, 6, }
    n1 = []int {0 ,}
    n2 = []int {1, }
    merge(n1, 0, n2, 1)
    fmt.Println(n1)
}

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	len1, len2 := len(nums1), len(nums2)
	i, j, q := m - 1, n - 1, m + n - 1
	if len1 == 0 || len2 == 0 {
		return
	}

	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[q] = nums2[j]
			j--
		} else {
			nums1[q] = nums1[i]
			i--
		}
		q--
	}

	if j < 0 {
		return
	}

    for i := 0; i < j + 1; i++ {
        nums1[i] = nums2[i]
    }
}
// @lc code=end

