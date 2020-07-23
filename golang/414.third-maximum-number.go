/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

// @lc code=start
func insert(arr []int64, val int64, cap int) int {
	for i := 0; i < 3; i ++ {
		if val > arr[i] {
			for j := 2; j > i; j-- {
				arr[j] = arr[j - 1]
			}
			arr[i] = val
			return cap + 1
		} else if val == arr[i] {
			break
		}
	}
	return cap
}

func solution1(nums []int) int {
	arr := []int64{ math.MinInt64, math.MinInt64, math.MinInt64, } // [max, mid, min]
	cap := 0
	for _, v := range(nums) {
		cap = insert(arr, int64(v), cap)
		fmt.Println(arr, cap)
	}
	if cap < 3 {
		return int(arr[0])
	}
	return int(arr[2])
}

func solution2(nums []int) int {
	max, mid, min := math.MinInt32, math.MinInt32, math.MinInt32
	maxv, midv, minv := 0, 0, 0

	for _, v := range(nums) {
		if v >= max {
			if v > max {
			max, mid, min = v, max, mid
			maxv, midv, minv = 1, maxv, midv
			} else {
				maxv = 1
			}
		} else if v >= mid {
			if v > mid {
				mid, min = v, mid
				midv, minv = 1, midv
			} else {
				midv = 1
			}
		} else if v >= min {
			if v > min {
				min = v
				minv = 1
			} else {
				minv = 1
			}
		}
		fmt.Println(max, mid, min)
		fmt.Println(maxv, midv, minv)
	}

	if minv == 0 {
		return max
	}
	
	return min
}

func thirdMax(nums []int) int {
	return solution2(nums)
}
// @lc code=end

