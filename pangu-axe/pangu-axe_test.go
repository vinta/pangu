package main

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"testing"
)

type PanguAxeTestSuite struct {
	suite.Suite
	realStdout *os.File
	r, w       *os.File
}

// The SetupTest method will be run before every test in the suite.
func (suite *PanguAxeTestSuite) SetupTest() {
	suite.realStdout = os.Stdout
	suite.r, suite.w, _ = os.Pipe()
	os.Stdout = suite.w
}

// The TearDownTest method will be run after every test in the suite.
func (suite *PanguAxeTestSuite) TearDownTest() {
	os.Stdout = suite.realStdout
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPanguTestSuite(t *testing.T) {
	suite.Run(t, new(PanguAxeTestSuite))
}

func (suite *PanguAxeTestSuite) getOutput() string {
	suite.w.Close()
	out, _ := ioutil.ReadAll(suite.r)
	outs := fmt.Sprintf("%s", out)

	return outs
}

func (suite *PanguAxeTestSuite) TestTextCmd() {
	os.Args = []string{NAME, "text", "新八的構造成分有95%是眼鏡、3%是水、2%是垃圾"}
	main()

	suite.Equal("新八的構造成分有 95% 是眼鏡、3% 是水、2% 是垃圾\n", suite.getOutput())
}

func (suite *PanguAxeTestSuite) TestFileCmd() {
	os.Args = []string{NAME, "file", "../_fixtures/test_file.txt"}
	main()

	suite.Equal("", suite.getOutput())
}
