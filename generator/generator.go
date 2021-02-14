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
		case parser.CmdLabel:
			g.handleLabel(command.Arg1, &asm)
		case parser.CmdGoto:
			g.handleGoto(command.Arg1, &asm)
		case parser.CmdIf:
			g.handleIf(command.Arg1, &asm)
		case parser.CmdFunction:
			g.handleFunction(command.Arg1, command.Arg2, &asm)
		case parser.CmdReturn:
			g.handleReturn(&asm)
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

func (g *Generator) handleLabel(label string, asm *strings.Builder) {
	asm.WriteString(fmt.Sprintf("(%s)\n\n", label))
}

func (g *Generator) handleGoto(label string, asm *strings.Builder) {
	asm.WriteString(fmt.Sprintf("@%s\n0;JMP\n\n", label))
}

func (g *Generator) handleIf(label string, asm *strings.Builder) {
	asm.WriteString(fmt.Sprintf("@SP\nAM=M-1\nD=M\n@%s\nD;JNE\n\n", label))
}

func (g *Generator) handleFunction(fn string, nVars int, asm *strings.Builder) {
	asm.WriteString(fmt.Sprintf("(%s)\n@SP\nA=M\n", fn))
	for i := 0; i < nVars; i++ {
		asm.WriteString("M=0\nA=A+1\n")
	}
	asm.WriteString("D=A\n@SP\nM=D\n\n")
}

func (g *Generator) handleReturn(asm *strings.Builder) {
	asm.WriteString("@LCL\nD=M\n@5\nA=D-A\nD=M\n@R13\nM=D\n@SP\nA=M-1\nD=M\n@ARG\nA=M\nM=D\nD=A+1\n@SP\nM=D\n@LCL\nAM=M-1\nD=M\n@THAT\nM=D\n@LCL\nAM=M-1\nD=M\n@THIS\nM=D\n@LCL\nAM=M-1\nD=M\n@ARG\nM=D\n@LCL\nA=M-1\nD=M\n@LCL\nM=D\n@R13\nA=M\n0;JMP\n\n")
}
