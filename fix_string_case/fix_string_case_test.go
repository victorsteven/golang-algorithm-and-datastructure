package fix_string_case_test

import (
	"efficient-algorithms/fix_string_case"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {

	assert.Equal(t, fix_string_case.Solve("Hello"), "normal_http_call")

}
