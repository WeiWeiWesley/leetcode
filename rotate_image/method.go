package rotateimage

func rotate(matrix [][]int) {
	n := len(matrix)

	// 先對矩陣進行 轉置 操作
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// 需同時對換對角元素，需寫在同一行
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 再對 row 進行 反轉 操作，將 row 的元素逆序排列
	for i := 0; i < n; i++ {
		// 到中間點即完成同 row 對調
		for j := 0; j < n/2; j++ {
			// 需同時對換同 row 元素，需寫在同一行
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
