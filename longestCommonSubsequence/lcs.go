package longestCommonSubsequence

import "github.com/kevin-untrojb/interview-problems/utils"

func longestCommonSubsequence(text1 string, text2 string) int {
	s1 := len(text1)
	s2 := len(text2)
	opt := make([][]int, s1+1)
	for i := range opt {
		opt[i] = make([]int, s2+1)
	}

	for i := 1; i <= s1; i++ {
		for j := 1; j <= s2; j++ {
			if text1[i-1] == text2[j-1] {
				opt[i][j] = opt[i-1][j-1] + 1
			} else {
				opt[i][j] = utils.Max(opt[i-1][j], opt[i][j-1])
			}

		}
	}

	return opt[s1][s2]
}
