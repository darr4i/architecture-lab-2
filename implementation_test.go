package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})



func (s *MySuite) TestCalculatePrefix(c *C) {
	result, err := CalculatePrefix("+ 5 * - 4 2 3")
	c.Assert(result, Equals, "11")
  
	result, err = CalculatePrefix("- * + 4 6 - 3 + * 4 6 * 3 5 - * 2 4 * 3 7")
	c.Assert(result, Equals, "-347")
  
	result, err = CalculatePrefix("")
	c.Assert(err, ErrorMatches, "Unexpected symbols or Extra spaces")
  
	result, err = CalculatePrefix("+ a b")
	c.Assert(err, ErrorMatches, "Unexpected symbols or Extra spaces")
  
	result, err = CalculatePrefix("- + 5 5")
	c.Assert(err, ErrorMatches, "Incorrect ratio of numbers and operators")
  
	result, err = CalculatePrefix("+ 5 0 2 4 + -")
	c.Assert(err, ErrorMatches, "Not a prefix notation order")
  }

  func ExampleCalculatePrefix() {
	res, err := CalculatePrefix("+ 5 * - 4 2 3")
	if err != nil {
	  panic(err)
	} else {
	  fmt.Println(res)
	}
	// Output: 11
  }
