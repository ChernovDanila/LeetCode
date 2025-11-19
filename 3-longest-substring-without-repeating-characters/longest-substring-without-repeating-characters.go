func lengthOfLongestSubstring(s string) int {

    maxLen := 0
	start := 0
	lenMap := map[byte]int{}

	for end := 0; end < len(s); end++ {

		if idx, ok := lenMap[s[end]]; ok && idx >= start {
			start = idx + 1
		}

		lenMap[s[end]] = end
		tempLen := end - start + 1
		if tempLen > maxLen{
			maxLen = tempLen
		}

	}
	return maxLen

}