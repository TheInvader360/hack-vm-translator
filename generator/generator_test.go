package generator

import (
	"testing"

	"github.com/TheInvader360/hack-vm-translator/parser"
	"github.com/stretchr/testify/assert"
)

func TestBootstrap(t *testing.T) {
	g := NewGenerator()
	bootstrap := g.Bootstrap()
	assert.Equal(t, "\n// Bootstrap...\n@256\nD=A\n@SP\nM=D\n@SP\nD=M\n@R13\nM=D\n@ret0\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@LCL\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@ARG\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@THIS\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@THAT\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@R13\nD=M\n@0\nD=D-A\n@ARG\nM=D\n@SP\nD=M\n@LCL\nM=D\n@Sys.init\n0;JMP\n(ret0)\n\n", bootstrap)
}

func TestGenerateAssembly(t *testing.T) {
	type test struct {
		commandType   parser.CommandType
		commandArg1   string
		commandArg2   int
		commandSource string
		expectedAsm   string
	}
	tests := []test{
		{commandType: parser.CmdArithmetic, commandArg1: "add", commandArg2: 0, commandSource: "add", expectedAsm: "// add\n@SP\nA=M-1\nD=M\nA=A-1\nM=D+M\n@SP\nM=M-1\n\n"},
		{commandType: parser.CmdArithmetic, commandArg1: "sub", commandArg2: 0, commandSource: "sub", expectedAsm: "// sub\n@SP\nA=M-1\nD=M\nA=A-1\nM=M-D\n@SP\nM=M-1\n\n"},
		{commandType: parser.CmdArithmetic, commandArg1: "neg", commandArg2: 0, commandSource: "neg", expectedAsm: "// neg\n@SP\nA=M\nA=A-1\nM=-M\n\n"},
		{commandType: parser.CmdArithmetic, commandArg1: "eq", commandArg2: 0, commandSource: "eq", expectedAsm: "// eq\n@SP\nA=M-1\nD=M\nA=A-1\nD=M-D\n@TRUE0\nD;JEQ\n@SP\nA=M-1\nA=A-1\nM=0\n@CONT0\n0;JMP\n(TRUE0)\n@SP\nA=M-1\nA=A-1\nM=-1\n(CONT0)\n@SP\nM=M-1\n\n"},
		{commandType: parser.CmdArithmetic, commandArg1: "gt", commandArg2: 0, commandSource: "gt", expectedAsm: "// gt\n@SP\nA=M-1\nD=M\nA=A-1\nD=M-D\n@TRUE0\nD;JGT\n@SP\nA=M-1\nA=A-1\nM=0\n@CONT0\n0;JMP\n(TRUE0)\n@SP\nA=M-1\nA=A-1\nM=-1\n(CONT0)\n@SP\nM=M-1\n\n"},
		{commandType: parser.CmdArithmetic, commandArg1: "lt", commandArg2: 0, commandSource: "lt", expectedAsm: "// lt\n@SP\nA=M-1\nD=M\nA=A-1\nD=M-D\n@TRUE0\nD;JLT\n@SP\nA=M-1\nA=A-1\nM=0\n@CONT0\n0;JMP\n(TRUE0)\n@SP\nA=M-1\nA=A-1\nM=-1\n(CONT0)\n@SP\nM=M-1\n\n"},
		{commandType: parser.CmdArithmetic, commandArg1: "and", commandArg2: 0, commandSource: "and", expectedAsm: "// and\n@SP\nA=M-1\nD=M\nA=A-1\nM=D&M\n@SP\nM=M-1\n\n"},
		{commandType: parser.CmdArithmetic, commandArg1: "or", commandArg2: 0, commandSource: "or", expectedAsm: "// or\n@SP\nA=M-1\nD=M\nA=A-1\nM=D|M\n@SP\nM=M-1\n\n"},
		{commandType: parser.CmdArithmetic, commandArg1: "not", commandArg2: 0, commandSource: "not", expectedAsm: "// not\n@SP\nA=M\nA=A-1\nM=!M\n\n"},
		{commandType: parser.CmdPush, commandArg1: "constant", commandArg2: 10, commandSource: "push constant 10", expectedAsm: "// push constant 10\n@10\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n"},
		{commandType: parser.CmdPush, commandArg1: "local", commandArg2: 11, commandSource: "push local 11", expectedAsm: "// push local 11\n@LCL\nD=M\n@11\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n"},
		{commandType: parser.CmdPush, commandArg1: "argument", commandArg2: 12, commandSource: "push argument 12", expectedAsm: "// push argument 12\n@ARG\nD=M\n@12\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n"},
		{commandType: parser.CmdPush, commandArg1: "this", commandArg2: 13, commandSource: "push this 13", expectedAsm: "// push this 13\n@THIS\nD=M\n@13\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n"},
		{commandType: parser.CmdPush, commandArg1: "that", commandArg2: 14, commandSource: "push that 14", expectedAsm: "// push that 14\n@THAT\nD=M\n@14\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n"},
		{commandType: parser.CmdPush, commandArg1: "pointer", commandArg2: 0, commandSource: "push pointer 0", expectedAsm: "// push pointer 0\n@THIS\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n"},
		{commandType: parser.CmdPush, commandArg1: "pointer", commandArg2: 1, commandSource: "push pointer 1", expectedAsm: "// push pointer 1\n@THAT\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n"},
		{commandType: parser.CmdPush, commandArg1: "temp", commandArg2: 15, commandSource: "push temp 15", expectedAsm: "// push temp 15\n@R5\nD=A\n@15\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n"},
		{commandType: parser.CmdPush, commandArg1: "static", commandArg2: 16, commandSource: "push static 16", expectedAsm: "// push static 16\n@filename.16\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n"},
		{commandType: parser.CmdPop, commandArg1: "local", commandArg2: 21, commandSource: "pop local 21", expectedAsm: "// pop local 21\n@LCL\nD=M\n@21\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n"},
		{commandType: parser.CmdPop, commandArg1: "argument", commandArg2: 22, commandSource: "pop argument 22", expectedAsm: "// pop argument 22\n@ARG\nD=M\n@22\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n"},
		{commandType: parser.CmdPop, commandArg1: "this", commandArg2: 23, commandSource: "pop this 23", expectedAsm: "// pop this 23\n@THIS\nD=M\n@23\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n"},
		{commandType: parser.CmdPop, commandArg1: "that", commandArg2: 24, commandSource: "pop that 24", expectedAsm: "// pop that 24\n@THAT\nD=M\n@24\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n"},
		{commandType: parser.CmdPop, commandArg1: "pointer", commandArg2: 0, commandSource: "pop pointer 0", expectedAsm: "// pop pointer 0\n@THIS\nD=A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n"},
		{commandType: parser.CmdPop, commandArg1: "pointer", commandArg2: 1, commandSource: "pop pointer 1", expectedAsm: "// pop pointer 1\n@THAT\nD=A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n"},
		{commandType: parser.CmdPop, commandArg1: "temp", commandArg2: 25, commandSource: "pop temp 25", expectedAsm: "// pop temp 25\n@R5\nD=A\n@25\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n"},
		{commandType: parser.CmdPop, commandArg1: "static", commandArg2: 26, commandSource: "pop static 26", expectedAsm: "// pop static 26\n@filename.26\nD=A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n"},
		{commandType: parser.CmdLabel, commandArg1: "LABEL", commandArg2: 0, commandSource: "label LABEL", expectedAsm: "// label LABEL\n($LABEL)\n\n"},
		{commandType: parser.CmdGoto, commandArg1: "LABEL", commandArg2: 0, commandSource: "goto LABEL", expectedAsm: "// goto LABEL\n@$LABEL\n0;JMP\n\n"},
		{commandType: parser.CmdIf, commandArg1: "LABEL", commandArg2: 0, commandSource: "if-goto LABEL", expectedAsm: "// if-goto LABEL\n@SP\nAM=M-1\nD=M\n@$LABEL\nD;JNE\n\n"},
		{commandType: parser.CmdFunction, commandArg1: "fn", commandArg2: 3, commandSource: "function fn 3", expectedAsm: "// function fn 3\n(fn)\n@SP\nA=M\nM=0\nA=A+1\nM=0\nA=A+1\nM=0\nA=A+1\nD=A\n@SP\nM=D\n\n"},
		{commandType: parser.CmdReturn, commandArg1: "", commandArg2: 0, commandSource: "return", expectedAsm: "// return\n@LCL\nD=M\n@5\nA=D-A\nD=M\n@R13\nM=D\n@SP\nA=M-1\nD=M\n@ARG\nA=M\nM=D\nD=A+1\n@SP\nM=D\n@LCL\nAM=M-1\nD=M\n@THAT\nM=D\n@LCL\nAM=M-1\nD=M\n@THIS\nM=D\n@LCL\nAM=M-1\nD=M\n@ARG\nM=D\n@LCL\nA=M-1\nD=M\n@LCL\nM=D\n@R13\nA=M\n0;JMP\n\n"},
		{commandType: parser.CmdCall, commandArg1: "fn", commandArg2: 2, commandSource: "call fn 2", expectedAsm: "// call fn 2\n@SP\nD=M\n@R13\nM=D\n@ret0\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@LCL\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@ARG\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@THIS\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@THAT\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n@R13\nD=M\n@2\nD=D-A\n@ARG\nM=D\n@SP\nD=M\n@LCL\nM=D\n@fn\n0;JMP\n(ret0)\n\n"},
	}
	for _, tc := range tests {
		g := NewGenerator()
		c := parser.Command{Type: tc.commandType, Arg1: tc.commandArg1, Arg2: tc.commandArg2, Source: tc.commandSource}
		asm := g.GenerateAssembly("filename", []parser.Command{c})
		assert.Equal(t, tc.expectedAsm, asm)
	}
}
