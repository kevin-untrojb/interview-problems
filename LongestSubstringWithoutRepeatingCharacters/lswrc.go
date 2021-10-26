package LongestSubstringWithoutRepeatingCharacters

import "github.com/kevin-untrojb/interview-problems/utils"

// website https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	m := make(map[uint8]struct{})
	max, i, j := 0, 0, 0

	for i < n && j < n {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = struct{}{}
			i++
			max = utils.Max(max, i-j)
			continue
		}
		delete(m, s[j])
		j++
	}
	return max
}
