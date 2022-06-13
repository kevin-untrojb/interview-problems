package remove_all_duplicant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveAllDuplicant(t *testing.T) {
	assert.Equal(t, "ybth",removeDuplicates("yfttttfbbbbnnnnffbgffffgbbbbgssssgthyyyy",4))
}
