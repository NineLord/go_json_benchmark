package Vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVector_Pop(t *testing.T) {
	assertions := assert.New(t)

	vector := MakeVector[int](0)
	vector.Push(1)
	vector.Push(2)
	vector.Push(3)

	assertions.Equal([]int{1, 2, 3}, vector.GetAll())
}
