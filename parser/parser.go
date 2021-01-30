package parser

import (
	"bufio"
	"bytes"
	"strings"
)

// Parser - struct
type Parser struct {
	SourceLines []string
}

// NewParser - returns a pointer to new parser
func NewParser() *Parser {
	return &Parser{}
}

// Sanitize - populate SourceLines with sanitized source data (comments and whitespace removed)
func (p *Parser) Sanitize(data []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line, "//")[0]
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			p.SourceLines = append(p.SourceLines, line)
		}
	}
}
