package reverseinterger

import (
	"math"
	"sort"
)

func reverse(x int) int {

	//-231 <= x <= 231 - 1
	if x <= int(math.Pow(-2, 31)) || x >= int(math.Pow(2, 31)-1) {
		return 0
	}

	carry := 1
	multiple := 0
	cArr := []int{}
	tag := false

	if x < 0 {
		tag = true
		x *= -1
	}

	//找出最大位數
	for {
		carry = int(x / int(math.Pow10(multiple)))

		//已達最大位數
		if carry == 0 {
			break
		}

		// fmt.Println("multiple", multiple, "carry", carry)

		cArr = append(cArr, carry)

		multiple++
	}

	//reverse
	sort.Ints(cArr)

	res := 0

	for i, v := range cArr {
		//留尾數
		remainder := v % 10
		//乘上位數
		res += remainder * int(math.Pow10(i))
	}

	if tag {
		res *= -1
	}

	//-231 <= x <= 231 - 1
	if res <= int(math.Pow(-2, 31)) || res >= int(math.Pow(2, 31)-1) {
		return 0
	}

	return res
}
