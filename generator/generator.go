package generator

import (
	"fmt"
	"strings"

	"github.com/TheInvader360/hack-vm-translator/parser"
)

// Generator - struct
type Generator struct {
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
			handleArithmetic(command.Arg1, &asm)
		case parser.CmdPush:
			handlePush(command.Arg1, command.Arg2, &asm)
		case parser.CmdPop:
			handlePop(command.Arg1, command.Arg2, &asm)
		}
	}
	return asm.String(), nil
}

func handleArithmetic(operation string, asm *strings.Builder) {
	comp := ""
	switch operation {
	case "add":
		comp = "M=M+D"
	case "sub":
		// TODO
	case "neg":
		// TODO
	case "eq":
		// TODO
	case "gt":
		// TODO
	case "lt":
		// TODO
	case "and":
		// TODO
	case "or":
		// TODO
	case "not":
		// TODO
	}
	asm.WriteString(fmt.Sprintf("@SP\nA=M-1\nD=M\nA=A-1\n%s\n@SP\nM=M-1\n\n", comp))
}

func handlePush(segment string, i int, asm *strings.Builder) {
	switch segment {
	case "local", "argument", "this", "that":
		asm.WriteString(fmt.Sprintf("// TODO: push %s %d\n\n", segment, i)) // TODO
	case "constant":
		asm.WriteString(fmt.Sprintf("@%d\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n\n", i))
	case "static":
		asm.WriteString(fmt.Sprintf("// TODO: push %s %d\n\n", segment, i)) // TODO
	case "temp":
		asm.WriteString(fmt.Sprintf("// TODO: push %s %d\n\n", segment, i)) // TODO
	case "pointer":
		asm.WriteString(fmt.Sprintf("// TODO: push %s %d\n\n", segment, i)) // TODO
	}
}

func handlePop(segment string, i int, asm *strings.Builder) {
	switch segment {
	case "local", "argument", "this", "that":
		asm.WriteString(fmt.Sprintf("// TODO: pop %s %d\n\n", segment, i)) // TODO
	case "static":
		asm.WriteString(fmt.Sprintf("// TODO: pop %s %d\n\n", segment, i)) // TODO
	case "temp":
		asm.WriteString(fmt.Sprintf("// TODO: pop %s %d\n\n", segment, i)) // TODO
	case "pointer":
		asm.WriteString(fmt.Sprintf("// TODO: pop %s %d\n\n", segment, i)) // TODO
	}
}
