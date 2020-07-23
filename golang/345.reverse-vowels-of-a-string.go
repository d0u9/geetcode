/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start
func isVowel(c byte) bool {
    if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
        return true
    }
    if c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
        return true
    }
    return false
}

func reverseVowels(s string) string {
    l := len(s)
    if l < 2 {
        return s
    }

    ret := []byte(s)

    for i, j := 0, l - 1; i <= j; {
        if !isVowel(ret[i]) {
            i++
            continue
        }
        if !isVowel(ret[j]) {
            j--
            continue
        }
        ret[i], ret[j] = ret[j], ret[i]
        i, j = i + 1, j - 1
    }
    return string(ret)
}
// @lc code=end

