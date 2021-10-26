package majorityElementfunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedMajorityElement1(t *testing.T) {

	arr := []int{0, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 2, 5, 6, 3}

	assert.Equal(t, SortedMajorityElement(arr), 5)
}
func TestMeOptimal(t *testing.T) {

	arr := []int{0, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 2, 5, 6, 3}

	assert.Equal(t, OptimalMajorityElement(arr), 5)
}
