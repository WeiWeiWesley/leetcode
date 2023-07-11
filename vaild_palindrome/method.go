package vaildpalindrome

import (
	"strings"
)

func isPalindrome(s string) bool {

	str := formateString(s)

	start := 0
	end := len(str) - 1

	for start < end {
		if str[start] != str[end] {
			return false
		}

		start++
		end--
	}

	return true
}

func formateString(s string) string {

	res := ""

	for _, v := range s {
		if isStringNum(v) {
			res += string(v)
		}
	}

	return strings.ToLower(res)
}

func isStringNum(s rune) bool {

	switch {
	case s >= '0' && s <= '9':
	case s >= 'a' && s <= 'z':
	case s >= 'A' && s <= 'Z':
	default:
		return false
	}

	return true
}
