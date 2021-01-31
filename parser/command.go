package parser

import (
	"fmt"

	"github.com/pkg/errors"
)

// CommandType - enum
type CommandType int

// CommandType - enum values
const (
	CmdArithmetic CommandType = iota
	CmdPush
	CmdPop
	CmdLabel
	CmdGoto
	CmdIf
	CmdFunction
	CmdReturn
	CmdCall
	CmdUnknown
)

// Command - struct
type Command struct {
	Type CommandType
	Arg1 string
	Arg2 int
}

// String representation of Command struct
func (c Command) String() string {
	return fmt.Sprintf("%v %s %d", c.Type, c.Arg1, c.Arg2)
}

// Returns a constant representing the type of the current command
func lookupCommandType(keyword string) (CommandType, error) {
	switch keyword {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		return CmdArithmetic, nil
	case "push":
		return CmdPush, nil
	case "pop":
		return CmdPop, nil
	case "label":
		return CmdLabel, nil
	case "goto":
		return CmdGoto, nil
	case "if-goto":
		return CmdIf, nil
	case "function":
		return CmdFunction, nil
	case "return":
		return CmdReturn, nil
	case "call":
		return CmdCall, nil
	}
	return CmdUnknown, errors.New("Invalid keyword: " + keyword)
}
