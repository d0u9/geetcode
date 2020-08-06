/*
 * @lc app=leetcode id=500 lang=golang
 *
 * [500] Keyboard Row
 */

// @lc code=start

func lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func findWords(words []string) []string {
    dict := map[byte]int {
		'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
	}

	ret := []string{}
	for j := 0; j < len(words); j++ {
		w := words[j]
		r := dict[ lower(w[0]) ]
		valid := true
		for i := 1; i < len(w); i++ {
			if dict[ lower(w[i]) ] != r {
				//fmt.Printf("r = %d, cur = %d\n", r, dict[ lower(w[i]) ])
				valid = false
				break
			}
		}

		if valid {
			ret = append(ret, w)
		}
	}

	return ret
}
// @lc code=end

