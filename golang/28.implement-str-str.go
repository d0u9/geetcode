/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
    l := len(haystack)
    ln := len(needle)

    if ln > l {
        return -1
    }
    if ln == 0 {
        return 0
    }

    for i := 0; i < l - ln + 1; i++ {
        if haystack[i] == needle[0] {
            f := true
            for j := 1; j < ln; j++ {
                if haystack[i + j] != needle[j] {
                    f = false
                    break
                }
            }
            if f == true {
                return i
            }
        }
    }

    return -1
}
// @lc code=end

