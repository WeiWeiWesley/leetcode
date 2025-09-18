package findindexoffirstoccurrence

// strStr 在 haystack 字串中尋找 needle 字串第一次出現的位置
// 使用暴力匹配法：從每個可能的位置開始檢查是否匹配
func strStr(haystack string, needle string) int {
	// 特殊情況：如果 needle 是空字串，根據題目要求返回 0
	if len(needle) == 0 {
		return 0
	}

	// 邊界檢查：如果 needle 比 haystack 長，不可能匹配
	if len(needle) > len(haystack) {
		return -1
	}

	// 主要邏輯：從 haystack 的每個位置開始嘗試匹配
	// 只需要檢查到 len(haystack)-len(needle) 的位置
	// 因為再往後就沒有足夠的字元來匹配 needle 了
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// 使用字串切片直接比較：從位置 i 開始取 len(needle) 個字元
		// 如果這個子字串等於 needle，就找到了第一次出現的位置
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	// 如果所有位置都檢查過了還是沒找到，返回 -1
	return -1
}
