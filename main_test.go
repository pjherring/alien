package main

import (
	. "launchpad.net/gocheck"
	"testing"
)

type AlienSuite struct {
}

var (
	_ = Suite(&AlienSuite{})
)

func TestAlienLanguage(t *testing.T) {
	TestingT(t)
}

func (a *AlienSuite) TestParseInputFile(c *C) {
	language, testCases, err := parseInputFile("input.txt")
	c.Assert(err, IsNil)
	c.Assert(language.Dictionary, HasLen, 5)
	c.Assert(language.Dictionary[0], Equals, "abc")
	c.Assert(language.Dictionary[1], Equals, "bca")
	c.Assert(language.Dictionary[2], Equals, "dac")
	c.Assert(language.Dictionary[2], Equals, "dbc")
	c.Assert(language.Dictionary[2], Equals, "cba")

	c.Assert(testCases, HasLen, 4)
	c.Assert(testCases[0], Equals, "(ab)(bc)(ca)")
	c.Assert(testCases[1], Equals, "abc")
	c.Assert(testCases[2], Equals, "(abc)(abc)(abc)")
	c.Assert(testCases[3], Equals, "(zyx)bc")
}

func (a *AlienSuite) TestSolveTestCase(c *C) {
}
