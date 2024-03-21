package lab2

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("- 6 2"),
		Output: b,
	}
	err := handler.Compute()

	assert.Equal(t, err, nil)
	assert.Equal(t, b.String(), "4")
}

func TestComputeHandlerHard(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("+ 5 - 6 7"),
		Output: b,
	}
	err := handler.Compute()

	assert.Equal(t, err, nil)
	assert.Equal(t, b.String(), "4")
}

func TestComputeHandlerError(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("1 2 3"),
		Output: b,
	}
	err := handler.Compute()

	assert.Equal(t, err, fmt.Errorf("Error. Missing arguments or many operators"))
}
