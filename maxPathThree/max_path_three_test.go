package maxPathThree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolutionWithOneElement(t *testing.T) {
	var tree = Tree{X: 1}
	assert.Equal(t, 1, maxPathThree(&tree))
}

func TestSolutionWithEmptyMap(t *testing.T) {
	assert.Equal(t, 0, maxPathThree(nil))
}

func TestSolutionWithTwoSons(t *testing.T) {
	var root = &Tree{
		X: 1,
		R: &Tree{
			X: 2,
			R: &Tree{
				X: 2,
				R: nil,
				L: nil,
			},
			L: &Tree{
				X: 1,
				R: nil,
				L: nil,
			},
		},
		L: &Tree{
			X: 3,
			R: nil,
			L: nil,
		},
	}
	assert.Equal(t, 2, maxPathThree(root))
}

func TestSolutionWithRepeteadSons(t *testing.T) {
	var root = &Tree{
		X: 1,
		R: &Tree{
			X: 2,
			R: &Tree{
				X: 2,
				R: nil,
				L: nil,
			},
			L: &Tree{
				X: 3,
				R: nil,
				L: &Tree{
					X: 2,
					R: nil,
					L: nil,
				},
			},
		},
		L: &Tree{
			X: 3,
			R: &Tree{
				X: 4,
				R: nil,
				L: &Tree{
					X: 5,
					R: nil,
					L: nil,
				},
			},
			L: nil,
		},
	}
	assert.Equal(t, 4, maxPathThree(root))
}
