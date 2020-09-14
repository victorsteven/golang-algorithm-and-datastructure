package vowel_count_test

import (
	"efficient-algorithms/vowel_count"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCount(t *testing.T) {

	assert.Equal(t, vowel_count.GetCount("hellomme"), 3)

	assert.Equal(t, vowel_count.GetCount4(" uethhztm sggi rprrxl risbs hsclup"), 5)
}


