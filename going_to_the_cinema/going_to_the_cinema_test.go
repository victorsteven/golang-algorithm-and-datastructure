package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovie(t *testing.T) {

	assert.Equal(t, Movie(500, 15, 0.9), 43)

	assert.Equal(t, Movie(100, 10, 0.95), 24)

}
