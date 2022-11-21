package main

// leetcode 3. 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	var (
		maxLen int
		start  int
	)
	seen := make(map[byte]int)
	for i, c := range []byte(s) {
		if last, ok := seen[c]; ok && last >= start {
			start = last + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		seen[c] = i
	}
	return maxLen
}
