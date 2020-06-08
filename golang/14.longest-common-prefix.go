/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start

/*
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	chs := make([]chan rune, len(strs))
	for i := range chs {
	  chs[i] = make(chan rune)
	}
  
	for _i, _s := range strs {
  
	  go func(i int, s string) {
		for _, c := range s {
		  chs[i] <- c
		}
  
		close(chs[i])
	  }(_i, _s)
  
	}
  
	var ret []rune
  L:
	for {
	  var char rune
	  for i := range chs {
		c, ok := <-chs[i]
		if !ok {
		  break L
		}
  
		if i == 0 {
		  char = c
		} else if c != char {
		  break L
		}
	  }
  
	  ret = append(ret, char)
	}
  
	return string(ret)
}  

*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	l := math.MaxInt32
	w := len(strs)
	for _, str := range strs {
		if len(str) < l {
			l = len(str)
		}
	}

	ret := make([]byte, 0)
L:
	for i := 0; i < l; i++ {
		cur := strs[0][i]
		for j := 1; j < w; j++ {
			if strs[j][i] != cur {
				break L
			}
		}
		ret = append(ret, cur)
	}

	return string(ret)
}

// @lc code=end

