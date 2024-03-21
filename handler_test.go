package lab2

import (
	"bytes"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestComputeHandler(c *C) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	
	handler := ComputeHandler{
		Input:  strings.NewReader("* 5 2"),
		Output: buffer,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(buffer.String(), Equals, "10")
}

func (s *MySuite) TestComputeHandlerError(c *C) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	
	handler := ComputeHandler{
		Input:  strings.NewReader("- 5 5 6"),
		Output: buffer,
	}
	err := handler.Compute()

	c.Assert(err, ErrorMatches, "Incorrect ratio of numbers and operators")
}
