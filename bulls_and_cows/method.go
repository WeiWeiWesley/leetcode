package bullsandcows

import (
	"fmt"
)

func getHint(secret string, guess string) string {

	sMap := map[string]int{}
	gMap := map[string]int{}

	A := 0
	B := 0

	for i := range secret {
		s := string(secret[i])

		if _, ok := sMap[s]; ok {
			sMap[s] = sMap[s] + 1
		} else {
			sMap[s] = 1
		}

		g := string(guess[i])

		if _, ok := gMap[g]; ok {
			gMap[g] = gMap[g] + 1
		} else {
			gMap[g] = 1
		}

		if s == g {
			A++
		}
	}

	for key := range gMap {
		if num, ok := sMap[key]; ok {
			if num > gMap[key] {
				B += gMap[key]
			} else {
				B += num
			}
		}
	}

	B -= A

	return fmt.Sprintf("%dA%dB", A, B)
}
