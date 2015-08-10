package fake_test

import (
	"io"
	"testing"

	. "gopkg.in/check.v1"
	"fmt"
	"reflect"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHelloWorld(c *C) {
	c.Assert(42, Equals, 42)
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}

type MyNextSuite struct{}

var _ = Suite(&MyNextSuite{})

func (s *MyNextSuite) TestHelloWorld2(c *C) {
	c.Assert(42, Equals, 42)
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}

func (s *MyNextSuite) TestHelloWorld3(c *C) {
	c.Assert(42, Equals, 42)
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}


func (s *MyNextSuite) SetUpSuite(c *C) {
	fmt.Println("->>> SetUpSuite:",reflect.TypeOf(s).String())
}

func (s *MyNextSuite) SetUpTest(c *C) {
	fmt.Println("\t->>> SetUpTest:",c.TestName())
}

func (s *MyNextSuite) TearDownTest(c *C) {
	fmt.Println("\t-<<< TearDownTest:",c.TestName())
	// Use the data in s.dir in the test.
}

func (s *MyNextSuite) TearDownSuite(c *C) {
	fmt.Println("-<<< TearDownSuite:",reflect.TypeOf(s).String())
}

type MyNextNextSuite struct{
	MyNextSuite
}

var _ = Suite(&MyNextNextSuite{})

func (s *MyNextNextSuite) TestHelloWorlSuite(c *C) {
	c.Assert(42, Equals, 42)
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}