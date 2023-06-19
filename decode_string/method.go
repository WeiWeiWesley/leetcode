package decodestring

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {

	//we need a stack(first in last out)
	stack := Stack{}

	for i := range s {

		// Find the first "]"
		if string(s[i]) != "]" {
			stack.Push(string(s[i])) //push into stack
			continue
		}

		str := ""
		index := len(stack.arr) - 1

		// 尋找 "[" 並記錄下 "[]" 中字串
		for index >= 0 {
			c := stack.Pop() //彈出
			index--          //記錄現在位置

			if c == "[" {
				break
			}

			str = c + str //記錄下 "[]" 中字串
		}

		strNum := ""
		for index >= 0 {
			if _, err := strconv.Atoi(stack.arr[index]); err != nil {
				// 不是數字了 跳出迴圈
				break
			}

			strNum = stack.Pop() + strNum
			index--
		}

		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}

		stack.Push(strings.Repeat(str, num))

		if str == "3 [ a cc" {
			fmt.Println(stack.arr, str)
		}
	}

	return strings.Join(stack.arr, "")
}

type Stack struct {
	arr []string
}

func (s *Stack) Push(target string) {
	s.arr = append(s.arr, target)
}

func (s *Stack) Pop() string {

	if len(s.arr) == 0 {
		return ""
	}

	l := len(s.arr) - 1

	pop := string(s.arr[l])

	s.arr = s.arr[:l]

	return pop
}
