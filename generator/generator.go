package generator

import (
	"fmt"
	"strings"

	"github.com/TheInvader360/hack-vm-translator/parser"
)

// Generator - struct
type Generator struct {
	labelCount int
}

// NewGenerator - returns a pointer to new generator
func NewGenerator() *Generator {
	return &Generator{}
}

// GenerateAssembly - generates assembly code lines from parsed VM commands
func (g *Generator) GenerateAssembly(commands []parser.Command) (string, error) {
	var asm strings.Builder
	for _, command := range commands {
		asm.WriteString(fmt.Sprintf("// %s\n", command.Source))
		switch command.Type {
		case parser.CmdArithmetic:
			g.handleArithmetic(command.Arg1, &asm)
		case parser.CmdPush:
			handlePush(command.Arg1, command.Arg2, &asm)
		case parser.CmdPop:
			handlePop(command.Arg1, command.Arg2, &asm)
		}
	}
	return asm.String(), nil
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

func handlePush(segment string, i int, asm *strings.Builder) {
	switch segment {
	case "constant":
		asm.WriteString(fmt.Sprintf("@%d\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", i))
	case "local", "argument", "this", "that":
		asm.WriteString(fmt.Sprintf("// TODO: push %s %d\n\n", segment, i)) // TODO
	case "pointer":
		asm.WriteString(fmt.Sprintf("// TODO: push %s %d\n\n", segment, i)) // TODO
	case "temp":
		asm.WriteString(fmt.Sprintf("// TODO: push %s %d\n\n", segment, i)) // TODO
	case "static":
		asm.WriteString(fmt.Sprintf("// TODO: push %s %d\n\n", segment, i)) // TODO
	}
}

func handlePop(segment string, i int, asm *strings.Builder) {
	switch segment {
	case "local", "argument", "this", "that":
		asm.WriteString(fmt.Sprintf("// TODO: pop %s %d\n\n", segment, i)) // TODO
	case "pointer":
		asm.WriteString(fmt.Sprintf("// TODO: pop %s %d\n\n", segment, i)) // TODO
	case "temp":
		asm.WriteString(fmt.Sprintf("// TODO: pop %s %d\n\n", segment, i)) // TODO
	case "static":
		asm.WriteString(fmt.Sprintf("// TODO: pop %s %d\n\n", segment, i)) // TODO
	}
}
