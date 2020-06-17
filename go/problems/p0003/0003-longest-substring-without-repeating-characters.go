package p0003

func lengthOfLongestSubstring(s string) int {
	stringMap := make(map[string]int)
	start := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if _, ok := stringMap[c]; ok {
			if stringMap[c] > start {
				start = stringMap[c]
			}
		}
		length := i - start + 1
		if ans < length {
			ans = length
		}
		stringMap[c] = i + 1
	}
	return ans
}
