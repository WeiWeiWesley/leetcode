package lettercombinationofphonenumber

//號碼輸入有順序性
func letterCombinations(digits string) []string {

	table := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	//持續記錄當下結果
	res := []string{}

	for i := range digits {

		num := string(digits[i])

		if pool, ok := table[num]; ok {

			if len(res) == 0 {
				res = pool
				continue
			}

			tmp := []string{}

			//每一次都直接使用前次結果組合
			//ex. ad + ("g", "h", "i")
			for _, v := range res {
				for _, v2 := range pool {
					tmp = append(tmp, v+v2)
				}
			}

			res = tmp
		}

	}

	return res
}
