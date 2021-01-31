package generator

import (
	"testing"

	"github.com/TheInvader360/hack-vm-translator/parser"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAssembly(t *testing.T) {
	type test struct {
		commandType    parser.CommandType
		commandArg1    string
		commandArg2    int
		commandSource  string
		expectedAsm    string
		expectedErrMsg string
	}
	tests := []test{
		{commandType: parser.CmdPush, commandArg1: "constant", commandArg2: 10, commandSource: "push constant 10", expectedAsm: "// push constant 10\n@10\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", expectedErrMsg: ""},
		{commandType: parser.CmdArithmetic, commandArg1: "add", commandArg2: 0, commandSource: "add", expectedAsm: "// add\n@SP\nA=M-1\nD=M\nA=A-1\nM=M+D\n@SP\nM=M-1\n\n", expectedErrMsg: ""},
	}
	for _, tc := range tests {
		g := NewGenerator()
		c := parser.Command{Type: tc.commandType, Arg1: tc.commandArg1, Arg2: tc.commandArg2, Source: tc.commandSource}
		asm, err := g.GenerateAssembly([]parser.Command{c})
		assert.Equal(t, tc.expectedAsm, asm)
		if len(tc.expectedErrMsg) > 0 {
			assert.EqualError(t, err, tc.expectedErrMsg)
		} else {
			assert.Nil(t, err)
		}
	}
}
