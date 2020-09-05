package invertbinarytrees

import (
	"efficient-algorithms/datastructures/trees"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvertTree(t *testing.T) {

	node2 := &trees.BinaryNode{Data: 2}
	node3 := &trees.BinaryNode{Data: 3}
	node := &trees.BinaryNode{Data: 1, Left: node2, Right: node3}

	node = InvertTree(node)

	assert.Equal(t, node.Left.Data, 3)
	assert.Equal(t, node.Right.Data, 2)
}

