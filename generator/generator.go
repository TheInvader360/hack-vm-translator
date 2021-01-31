package generator

import (
	"fmt"

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
func (g *Generator) GenerateAssembly(commands []parser.Command) ([]string, error) {
	lines := []string{}
	for i, command := range commands {
		line := fmt.Sprintf("TODO: %d %s", i, command) //TODO!
		lines = append(lines, line)
	}
	return lines, nil
}
