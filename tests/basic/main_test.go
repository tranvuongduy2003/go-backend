package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input = 1
	// 	output = 2
	// )

	// actual := AddOne(input)
	// if actual != output {
	// 	t.Errorf("Expected %d, but got %d", output, actual)
	// }

	assert.Equal(t, 2, AddOne(1), "AddOne should return input + 1")

	assert.NotEqual(t, 2, 3)
	assert.Nil(t, nil, nil)
}