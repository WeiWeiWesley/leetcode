package numberofisland

//邊界探索 DFS
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }

    rows := len(grid)
    cols := len(grid[0])
    count := 0

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == '1' {
                count++
                dfs(grid, i, j, rows, cols) //將相鄰的部分移除，避免重複計算
            }
        }
    }

    return count
}

func dfs(grid [][]byte, row, col, rows, cols int) {

	//超出邊界終止
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return
	}

	//碰到海水 '0' 終止
    if grid[row][col] == '0' {
        return
    }

    grid[row][col] = '0' // 將島嶼標記為已訪問

    // 遞歸探索相鄰的島嶼
    dfs(grid, row+1, col, rows, cols) //上
    dfs(grid, row-1, col, rows, cols) //下
    dfs(grid, row, col-1, rows, cols) //左
    dfs(grid, row, col+1, rows, cols) //右
}
