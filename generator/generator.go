package generator

import (
	"fmt"
	"strings"

	"github.com/TheInvader360/hack-vm-translator/parser"
)

// Generator - struct
type Generator struct {
	filename   string
	labelCount int
}

// NewGenerator - returns a pointer to new generator
func NewGenerator() *Generator {
	return &Generator{}
}

// GenerateAssembly - generates assembly code lines from parsed VM commands
func (g *Generator) GenerateAssembly(filename string, commands []parser.Command) string {
	g.filename = filename
	fmt.Println(fmt.Sprintf("Translating %s:", filename))
	fmt.Println("------------------------------")
	for _, command := range commands {
		fmt.Println(command.Source)
	}
	fmt.Println("------------------------------")

	var asm strings.Builder
	for _, command := range commands {
		asm.WriteString(fmt.Sprintf("// %s\n", command.Source))
		switch command.Type {
		case parser.CmdArithmetic:
			g.handleArithmetic(command.Arg1, &asm)
		case parser.CmdPush:
			g.handlePush(command.Arg1, command.Arg2, &asm)
		case parser.CmdPop:
			g.handlePop(command.Arg1, command.Arg2, &asm)
		}
	}
	return asm.String()
}

func (g *Generator) handleArithmetic(operation string, asm *strings.Builder) {
	switch operation {
	case "add":
		asm.WriteString("@SP\nA=M-1\nD=M\nA=A-1\nM=D+M\n@SP\nM=M-1\n\n")
	case "sub":
		asm.WriteString("@SP\nA=M-1\nD=M\nA=A-1\nM=M-D\n@SP\nM=M-1\n\n")
	case "neg":
		asm.WriteString("@SP\nA=M\nA=A-1\nM=-M\n\n")
	case "eq":
		asm.WriteString(fmt.Sprintf("@SP\nA=M-1\nD=M\nA=A-1\nD=M-D\n@TRUE%[1]d\nD;JEQ\n@SP\nA=M-1\nA=A-1\nM=0\n@CONT%[1]d\n0;JMP\n(TRUE%[1]d)\n@SP\nA=M-1\nA=A-1\nM=-1\n(CONT%[1]d)\n@SP\nM=M-1\n\n", g.labelCount))
		g.labelCount++
	case "gt":
		asm.WriteString(fmt.Sprintf("@SP\nA=M-1\nD=M\nA=A-1\nD=M-D\n@TRUE%[1]d\nD;JGT\n@SP\nA=M-1\nA=A-1\nM=0\n@CONT%[1]d\n0;JMP\n(TRUE%[1]d)\n@SP\nA=M-1\nA=A-1\nM=-1\n(CONT%[1]d)\n@SP\nM=M-1\n\n", g.labelCount))
		g.labelCount++
	case "lt":
		asm.WriteString(fmt.Sprintf("@SP\nA=M-1\nD=M\nA=A-1\nD=M-D\n@TRUE%[1]d\nD;JLT\n@SP\nA=M-1\nA=A-1\nM=0\n@CONT%[1]d\n0;JMP\n(TRUE%[1]d)\n@SP\nA=M-1\nA=A-1\nM=-1\n(CONT%[1]d)\n@SP\nM=M-1\n\n", g.labelCount))
		g.labelCount++
	case "and":
		asm.WriteString("@SP\nA=M-1\nD=M\nA=A-1\nM=D&M\n@SP\nM=M-1\n\n")
	case "or":
		asm.WriteString("@SP\nA=M-1\nD=M\nA=A-1\nM=D|M\n@SP\nM=M-1\n\n")
	case "not":
		asm.WriteString("@SP\nA=M\nA=A-1\nM=!M\n\n")
	}
}

func (g *Generator) handlePush(segment string, i int, asm *strings.Builder) {
	switch segment {
	case "constant":
		// *SP=i, SP++
		asm.WriteString(fmt.Sprintf("@%d\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", i))
	case "local":
		// addr=LCL+i, *SP=*addr, SP++
		asm.WriteString(fmt.Sprintf("@LCL\nD=M\n@%d\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", i))
	case "argument":
		// addr=ARG+i, *SP=*addr, SP++
		asm.WriteString(fmt.Sprintf("@ARG\nD=M\n@%d\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", i))
	case "this":
		// addr=THIS+i, *SP=*addr, SP++
		asm.WriteString(fmt.Sprintf("@THIS\nD=M\n@%d\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", i))
	case "that":
		// addr=THAT+i, *SP=*addr, SP++
		asm.WriteString(fmt.Sprintf("@THAT\nD=M\n@%d\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", i))
	case "pointer":
		if i == 0 {
			// *SP=THIS, SP++
			asm.WriteString("@THIS\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n")
		} else {
			// *SP=THAT, SP++
			asm.WriteString("@THAT\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n")
		}
	case "temp":
		// addr=5+i, *SP=*addr, SP++
		asm.WriteString(fmt.Sprintf("@R5\nD=A\n@%d\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", i))
	case "static":
		// static i in file.vm -> file.i in file.asm
		asm.WriteString(fmt.Sprintf("@%s.%d\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", g.filename, i))
	}
}

func (g *Generator) handlePop(segment string, i int, asm *strings.Builder) {
	switch segment {
	case "local":
		// addr=LCL+i, SP--, *addr=*SP
		asm.WriteString(fmt.Sprintf("@LCL\nD=M\n@%d\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n", i))
	case "argument":
		// addr=ARG+i, SP--, *addr=*SP
		asm.WriteString(fmt.Sprintf("@ARG\nD=M\n@%d\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n", i))
	case "this":
		// addr=THIS+i, SP--, *addr=*SP
		asm.WriteString(fmt.Sprintf("@THIS\nD=M\n@%d\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n", i))
	case "that":
		// addr=THAT+i, SP--, *addr=*SP
		asm.WriteString(fmt.Sprintf("@THAT\nD=M\n@%d\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n", i))
	case "pointer":
		if i == 0 {
			// SP--, THIS=*SP
			asm.WriteString("@THIS\nD=A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n")
		} else {
			// SP--, THAT=*SP
			asm.WriteString("@THAT\nD=A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n")
		}
	case "temp":
		// addr=5+i, SP--, *addr=*SP
		asm.WriteString(fmt.Sprintf("@R5\nD=A\n@%d\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n", i))
	case "static":
		// static i in file.vm -> file.i in file.asm
		asm.WriteString(fmt.Sprintf("@%s.%d\nD=A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n\n", g.filename, i))
	}
}
