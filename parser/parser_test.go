package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandStringer(t *testing.T) {
	c := Command{Type: CmdPush, Arg1: "constant", Arg2: 10, Source: "push constant 10"}
	assert.Equal(t, "1 constant 10 (push constant 10)", fmt.Sprint(c))
}

func TestSanitize(t *testing.T) {
	d := []byte("// comment\n\n// comment comment\npush constant 10 // comment\n\npop local 0\n\n\npush constant 21\n")
	p := NewParser()
	p.Sanitize(d)
	assert.Equal(t, "push constant 10\npop local 0\npush constant 21", strings.Join(p.SourceLines, "\n"))
}

func TestParseSource(t *testing.T) {
	type test struct {
		sourceLine     string
		expectedType   CommandType
		expectedArg1   string
		expectedArg2   int
		expectedErrMsg string
	}
	tests := []test{
		// Valid source lines
		{sourceLine: "add", expectedType: CmdArithmetic, expectedArg1: "add", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "sub", expectedType: CmdArithmetic, expectedArg1: "sub", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "neg", expectedType: CmdArithmetic, expectedArg1: "neg", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "eq", expectedType: CmdArithmetic, expectedArg1: "eq", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "gt", expectedType: CmdArithmetic, expectedArg1: "gt", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "lt", expectedType: CmdArithmetic, expectedArg1: "lt", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "and", expectedType: CmdArithmetic, expectedArg1: "and", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "or", expectedType: CmdArithmetic, expectedArg1: "or", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "not", expectedType: CmdArithmetic, expectedArg1: "not", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "push constant 10", expectedType: CmdPush, expectedArg1: "constant", expectedArg2: 10, expectedErrMsg: ""},
		{sourceLine: "pop local 0", expectedType: CmdPop, expectedArg1: "local", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "label l", expectedType: CmdLabel, expectedArg1: "l", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "goto g", expectedType: CmdGoto, expectedArg1: "g", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "if-goto i", expectedType: CmdIf, expectedArg1: "i", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "function file.f 5", expectedType: CmdFunction, expectedArg1: "file.f", expectedArg2: 5, expectedErrMsg: ""},
		{sourceLine: "return", expectedType: CmdReturn, expectedArg1: "", expectedArg2: 0, expectedErrMsg: ""},
		{sourceLine: "call file.c 5", expectedType: CmdCall, expectedArg1: "file.c", expectedArg2: 5, expectedErrMsg: ""},
		// Invalid source lines
		{sourceLine: "fail", expectedType: CmdUnknown, expectedArg1: "", expectedArg2: 0, expectedErrMsg: "Command type lookup failed at line 1 (fail): Invalid keyword: fail"},
		{sourceLine: "push", expectedType: CmdUnknown, expectedArg1: "", expectedArg2: 0, expectedErrMsg: "Missing arg1 at line 1 (push)"},
		{sourceLine: "pop local", expectedType: CmdUnknown, expectedArg1: "", expectedArg2: 0, expectedErrMsg: "Missing arg2 at line line 1 (pop local)"},
		{sourceLine: "pop local x", expectedType: CmdUnknown, expectedArg1: "", expectedArg2: 0, expectedErrMsg: `Invalid arg2 at line line 1 (pop local x): strconv.Atoi: parsing "x": invalid syntax`},
	}
	for _, tc := range tests {
		p := NewParser()
		p.SourceLines = append(p.SourceLines, tc.sourceLine)
		commands, err := p.ParseSource()
		if len(commands) > 0 {
			assert.Equal(t, tc.expectedType, commands[0].Type)
			assert.Equal(t, tc.expectedArg1, commands[0].Arg1)
			assert.Equal(t, tc.expectedArg2, commands[0].Arg2)
			assert.Equal(t, tc.sourceLine, commands[0].Source)
		}
		if len(tc.expectedErrMsg) > 0 {
			assert.EqualError(t, err, tc.expectedErrMsg)
		} else {
			assert.Nil(t, err)
		}
	}
}
