package Isograms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsogram(t *testing.T) {

	assert.Equal(t, Isogram("Dermatoglyphics"), true)
	assert.Equal(t, Isogram("aba"), false)
	assert.Equal(t, Isogram("moOse"), false)
	assert.Equal(t, Isogram("isogram"), true)
	assert.Equal(t, Isogram("moose"), false)
	assert.Equal(t, Isogram("isIsogram"), false)
	assert.Equal(t, Isogram("thumbscrewjapingly"), true)
	assert.Equal(t, Isogram(""), true)

}
