package paperfold_sequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaperFold(t *testing.T) {

	arr := []int{}

	ch := make(chan int, 100)
	go PaperFold(ch)

	for i := 0; i < 20; i++ {
		arr = append(arr, <-ch)
	}
	assert.Equal(t, arr, []int{1,1,0,1,1,0,0,1,1,1,0,0,1,0,0,1,1,1,0,1})
}
