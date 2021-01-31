package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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

// ParseSource - parse SourceLines into their lexical components and populate Commands with the results
func (p *Parser) ParseSource() ([]Command, error) {
	commands := []Command{}
	for i, line := range p.SourceLines {
		tokens := strings.Split(line, " ")
		cmdType, err := lookupCommandType(tokens[0])
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Command type lookup failed at line %d (%s)", i+1, line))
		}
		command := Command{Type: cmdType, Source: line}

		// first arg - command itself if arithmetic, not set if return
		if command.Type == CmdArithmetic {
			command.Arg1 = tokens[0]
		} else if command.Type != CmdReturn {
			if len(tokens) < 2 {
				return nil, errors.New(fmt.Sprintf("Missing arg1 at line %d (%s)", i+1, line))
			}
			command.Arg1 = tokens[1]
		}

		// second arg - only set if push, pop, function, or call
		if command.Type == CmdPush || command.Type == CmdPop || command.Type == CmdFunction || command.Type == CmdCall {
			if len(tokens) < 3 {
				return nil, errors.New(fmt.Sprintf("Missing arg2 at line line %d (%s)", i+1, line))
			}
			command.Arg2, err = strconv.Atoi(tokens[2])
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("Invalid arg2 at line line %d (%s)", i+1, line))
			}
		}

		commands = append(commands, command)
	}
	return commands, nil
}
