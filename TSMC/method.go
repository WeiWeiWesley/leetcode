package tsmc

/*
 * hackerrank Bucket Full
 * https://stackoverflow.com/questions/64111062/want-to-get-solution-for-the-following-problem-of-picture-fill-algo
 */

func strokesRequired(picture []string) int32 {
	count := int32(0)                       // 計數器，用於計算所需的筆劃數
	visited := make([][]bool, len(picture)) // 記錄每個點是否已經訪問過

	// 初始化 visited 矩陣
	for i := range visited {
		visited[i] = make([]bool, len(picture[i]))
	}

	// 遍歷圖像的每個點
	for i := range picture {
		for j := range picture[i] {
			if !visited[i][j] { // 如果點未被訪問過
				color := picture[i][j]              // 當前點的顏色
				fill(picture, visited, i, j, color) // 開始填充相同顏色的區域
				count++                             // 每次填充完成，計數器加1
			}
		}
	}

	return count
}

// fill 函數用於填充相同顏色的區域
func fill(picture []string, visited [][]bool, row, col int, color byte) {
	if row < 0 || row >= len(picture) || col < 0 || col >= len(picture[row]) {
		return // 範圍超出邊界，返回
	}

	if visited[row][col] || picture[row][col] != color {
		return // 當前點已訪問過或顏色不匹配，返回
	}

	visited[row][col] = true                                // 標記當前點為已訪問
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 定義四個方向：上、下、左、右

	// 遍歷四個方向的鄰居點
	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]
		fill(picture, visited, newRow, newCol, color) // 遞歸填充鄰居點
	}
}
