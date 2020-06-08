/*
* @lc app=leetcode id=27 lang=golang
*
* [27] Remove Element
*/

// @lc code=start
func removeElement2(nums []int, val int) int {
    sort.Ints(nums)
    ol := len(nums)

    start, end := 0, 0
    var i, n int
    for i, n = range nums {
        if n == val {
            break
        }
    }

    if i == ol {
        return 0
    }

    start = i

    for i = start; i < ol; i++ {
        if nums[i] != val {
            break
        }
    }
    end = i

    for i, j := start, end; i < ol && j < ol; {
        nums[i] = nums[j]
        i++
        j++
    }

    return ol - (end - start)
}

func removeElement(nums []int, val int) int {
    ol := len(nums)

    if ol == 0 {
        return 0
    }

    var i, j int
    for i = 0; i < ol; {
        for i < ol {
            if nums[i] == val {
                break
            }
            i++
        }

        for j = i + 1; j < ol; {
            if nums[j] != val {
                break
            }
            j++
        }

        if i < ol && j < ol {
            nums[i] = nums[j]
            nums[j] = val
            i++
            j++
        }

        if j >= ol {
            break
        }
    }

    return i
}

// @lc code=end

