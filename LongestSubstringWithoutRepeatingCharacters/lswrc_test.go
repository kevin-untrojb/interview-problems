package LongestSubstringWithoutRepeatingCharacters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLswrc1(t *testing.T) {
	s := "abcabcbb"
	assert.Equal(t, lengthOfLongestSubstring(s), 3)
}
func TestLswrc2(t *testing.T) {
	s := "bbbbb"
	assert.Equal(t, lengthOfLongestSubstring(s), 1)
}
func TestLswrc3(t *testing.T) {
	s := "pwwkew"
	assert.Equal(t, lengthOfLongestSubstring(s), 3)
}
func TestLswrc4(t *testing.T) {
	s := ""
	assert.Equal(t, lengthOfLongestSubstring(s), 0)
}
func TestLswrc5(t *testing.T) {
	s := "asdfggegegegegegegegegegegegegegegegegegegegegegegegg"
	assert.Equal(t, lengthOfLongestSubstring(s), 5)
}
