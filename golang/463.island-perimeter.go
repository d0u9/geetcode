/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 */

// @lc code=start

func block_perimeter(i, j int, x, y int, grid [][]int) int {
	if grid[i][j] == 0 {
		return 0
	}
	ux, uy := i - 1, j
	lx, ly := i, j - 1
	dx, dy := i + 1, j
	rx, ry := i, j + 1
	ret := 4

	//fmt.Printf("i=%d, j=%d, x=%d, y=%d, ret=%d\n", i, j, x, y, -1)
	if ux >= 0 && grid[ux][uy] == 1 {
		ret--
		//fmt.Println(1)
	}

	if ly >= 0 && grid[lx][ly] == 1 {
		ret --
		//fmt.Println(2)
	}

	if dx < x && grid[dx][dy] == 1 {
		ret--
		//fmt.Println(3)
	}

	if ry < y && grid[rx][ry] == 1 {
		ret --
		//fmt.Println(4)
	}

	//fmt.Printf("i=%d, j=%d, x=%d, y=%d, ret=%d\n", i, j, x, y, ret)
	return ret
}

func solution1(grid [][]int) int {
	x := len(grid)
	if x <= 0 {
		return 0
	}
	y := len(grid[0])
	if y <= 0 {
		return 0
	}

	//fmt.Println(x, y)
	ret := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			//fmt.Println(i, j)
			ret = ret + block_perimeter(i, j, x, y, grid)
		}
	}
	return ret
}

// iter
func recr_iter(i, j, x, y int, grid [][]int) int {
	if i < 0 || j < 0 || i >= x || j >= y {
		return 1
	}

	if grid[i][j] == 0 {
		return 1
	}

	if grid[i][j] == -1 {
		return 0
	}

	//fmt.Printf("recr_iter(%d, %d)\n", i, j)

	grid[i][j] = -1

	cur := 0
	// up
	if v := recr_iter(i - 1, j, x, y, grid); v != 0 {
		cur = cur + v
	} else {
		//cur++
	}
	// down
	if v := recr_iter(i + 1, j, x, y, grid); v != 0 {
		cur = cur + v
	} else {
		//cur++
	}
	// left
	if v := recr_iter(i, j - 1, x, y, grid); v != 0 {
		cur = cur + v
	} else {
		//cur++
	}
	// right
	if v := recr_iter(i, j + 1, x, y, grid); v != 0 {
		cur = cur + v
	} else {
		//cur++
	}
	//fmt.Printf("recr_iter(%d, %d), ret=%d\n", i, j, cur)
	
	return cur
}

// back trace
func solution2(grid [][]int) int {
	x := len(grid)
	if x <= 0 {
		return 0
	}
	y := len(grid[0])
	if y <= 0 {
		return 0
	}
	i, j := 0, 0
L:
	for i = 0; i < x; i++ {
		for j = 0; j < y; j++ {
			if grid[i][j] == 1 {
				break L
			}
		}
	}
	return recr_iter(i, j, x, y, grid)
}
	
func islandPerimeter(grid [][]int) int {
	return solution2(grid)
}
// @lc code=end

