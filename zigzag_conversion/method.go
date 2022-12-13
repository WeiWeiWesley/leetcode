package zigzagconversion

import (
	"math"
)

/*
PAYPALISHIRING
01234567890123
PAHNAPLSIIGYIR
*/
func convert(s string, numRows int) string {

	if len(s) < 3 || numRows == 1 {
		return s
	}

	n := len(s)

	block := int(math.Ceil((float64(n) / (2*float64(numRows) - 2))))

	col := block * (numRows - 1)

	matrix := make([][]string, numRows)

	for i := range matrix {
		matrix[i] = make([]string, col)
	}

	var strIndex, rowIndex, colIndex int

	for strIndex < n {

		// move down
		for rowIndex < numRows && strIndex < n {
			matrix[rowIndex][colIndex] = string(s[strIndex])
			rowIndex++
			strIndex++
		}

		rowIndex -= 2
		colIndex++

		// move up
		for rowIndex > 0 && colIndex < col && strIndex < n {
			matrix[rowIndex][colIndex] = string(s[strIndex])
			rowIndex--
			colIndex++
			strIndex++
		}
	}

	res := ""

	for _, row := range matrix {
		for _, char := range row {
			if char != "" {
				res += char
			}
		}
	}

	return res
}
