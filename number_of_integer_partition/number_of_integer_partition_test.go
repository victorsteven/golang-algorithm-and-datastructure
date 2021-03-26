package number_of_integer_partition_test

import (
	"efficient-algorithms/number_of_integer_partition"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartitions(t *testing.T) {

	assert.Equal(t, number_of_integer_partition.Partitions(5), 7)

}
