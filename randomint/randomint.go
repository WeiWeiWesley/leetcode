package randomint

func romanToInt(s string) int {

	romanMap := make(map[string]int, 17)
	romanMap["MM"] = 2000
	romanMap["M"] = 1000
	romanMap["CM"] = 900
	romanMap["D"] = 500
	romanMap["CD"] = 400
	romanMap["CC"] = 200
	romanMap["C"] = 100
	romanMap["XC"] = 90
	romanMap["L"] = 50
	romanMap["XL"] = 40
	romanMap["XX"] = 20
	romanMap["X"] = 10
	romanMap["IX"] = 9
	romanMap["V"] = 5
	romanMap["IV"] = 4
	romanMap["II"] = 2
	romanMap["I"] = 1

	res := 0

	length := len(s)

	for i := 1; i <= length; i++ {

		if i < length {
			if v, ok := romanMap[s[i-1:i+1]]; ok {
				res += v
				i++
				continue
			}
		}

		if v, ok := romanMap[s[i-1:i]]; ok {
			res += v
			continue
		}

	}

	return res
}
