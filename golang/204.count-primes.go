/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
func isPrimes(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n % 2 == 0 || n % 3 == 0 {
		return false
	}

	for i := 5; i * i <= n; i = i + 6 {
		if n % i == 0 || n % (i + 2) == 0 {
			return false
		}
	}

	return true
}

func solution1(n int) int {
	c := 0
	for i := 2; i < n; i++ {
		if isPrimes(i) {
			c++
		}
	}
    return c
}

func solution2(n int) int {
	// Timelimit exceed
	switch(n) {
	case 0: return 0
	case 1: return 0
	case 2: return 0
	case 3: return 1
	case 4: return 2
	}
	m := map[int]bool{ 2: true, 3: true, }
	c := 2
	for i := 4; i < n; i++ {
		z := false
		for v := range m {
			if i % v == 0 {
				z = true
				break
			}
		}

		if !z {
			m[i] = true
			c++
		}
	}
	return c
}

func solution3(n int) int {
	// From leetcode hints
	mp := make([]bool, n, n)
	for i := 0; i < n; i++ {
		mp[i] = true
	}
	for i := 2; i * i < n; i++ {
		if !isPrimes(i) {
			continue
		}
		for j := i * i; j < n; j += i {
			mp[j] = false
		}
	}

	c := 0
	for i := 2; i < n; i++ {
		if mp[i] {
			c++
		}
	}

	return c
}

func countPrimes(n int) int {
	//return solution1(n)
	//return solution2(n)
	return solution3(n)
}
// @lc code=end

