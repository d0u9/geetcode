/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func solution1(pattern string, str string) bool {
    mp := make(map[byte]string)
    mp2 := make(map[string]byte)
    pat_idx, wc := 0, 0
    for c, e := 0, 0; e <= len(str); e++ {
        if e == len(str) || str[e] == ' ' {
            if c != 0 {
                cstr := str[e-c:e]
                fmt.Printf("e = %d, c = %d\n", e, c)
                fmt.Println(cstr)
                fmt.Println(pat_idx)

                if pat_idx >= len(pattern) {
                    return false
                }
                p, ok := mp[pattern[pat_idx]];
                b, ok2 := mp2[cstr]
                if (ok && p != cstr) || (ok2 && b != pattern[pat_idx]) {
                    //fmt.Println(p, cstr)
                    return false
                }
                if !ok {
                    mp[pattern[pat_idx]] = str[e-c:e]
                    mp2[cstr] = pattern[pat_idx]
                }
                wc++
            }
            pat_idx++
            c = 0
        } else {
            c++
        }
    }
    if wc != len(pattern) {
        return false
    }
    return true
}

func solution2(pattern string, str string) bool {
    sub := strings.Split(str, " ")
    l, l1 := len(pattern), len(sub)
    if l != l1 {
        return false
    }
    mpp := make(map[byte]string)
    mps := make(map[string]byte)
    for i := 0; i < l; i++ {
        pk, sk := pattern[i], sub[i]
        pv, pok := mpp[pk]
        sv, sok := mps[sk]
        if (pok != sok) {
            return false
        } else if (pok && sok && (pk != sv || pv != sk)) {
            return false
        } else {
            mpp[pk] = sk
            mps[sk] = pk
        }
    }
    return true
}

func wordPattern(pattern string, str string) bool {
    return solution2(pattern, str)
}
// @lc code=end

