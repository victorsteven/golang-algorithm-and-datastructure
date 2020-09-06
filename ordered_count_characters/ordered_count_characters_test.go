package ordered_count_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedCount(t *testing.T) {

	ans := OrderedCount("abracadabra")

	expected := []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}

	assert.Equal(t, expected, ans)



	ans2 := OrderedCount2("abracadabra")

	expected2 := []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}

	assert.Equal(t, expected2, ans2)

}


