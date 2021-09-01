package missingNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissingNumber(t *testing.T) {

	arr := []int{0,1,4,2,5,6,7,9,8,3,11}

	assert.Equal(t, TrivialMissingNumber(arr), 10)
}
func TestMissingNumber2(t *testing.T) {

	arr := []int{0,1,4,2,5,6,3}

	assert.Equal(t, TrivialMissingNumber(arr), 7)
}

func TestMissingNumber3(t *testing.T) {

	arr := []int{0,1,4,2,5,6,7,9,8,3,11}

	assert.Equal(t, SortedMissingNumber(arr), 10)
}
func TestMissingNumber4(t *testing.T) {

	arr := []int{0,1,4,2,5,6,3}

	assert.Equal(t, SortedMissingNumber(arr), 7)
}
func TestMissingNumber5(t *testing.T) {

	arr := []int{0,1,4,2,5,6,7,9,8,3,11}

	assert.Equal(t, HashMissingNumber(arr), 10)
}
func TestMissingNumber6(t *testing.T) {

	arr := []int{0,1,4,2,5,6,3}

	assert.Equal(t, HashMissingNumber(arr), 7)
}
func TestMissingNumber7(t *testing.T) {

	arr := []int{0,1,4,2,5,6,7,9,8,3,11}

	assert.Equal(t, OptMissingNumber(arr), 10)
}
func TestMissingNumber8(t *testing.T) {

	arr := []int{0,1,4,2,5,6,3}

	assert.Equal(t, OptMissingNumber(arr), 7)
}