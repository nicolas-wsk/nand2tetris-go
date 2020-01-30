package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Command types
const (
	CmdTypePush       = iota // push
	CmdTypePop               // pop
	CmdTypeArithmetic        // arithmetic
	CmdTypeComparator        // gt, lt, eq
	// CmdTypeBranching         // label, goto, if-goto
	// CmdTypeFunction          // function
	// CmdTypeCall              // call
	// CmdTypeReturn            // return
)

// Parser take a file and split it
type Parser struct {
	file    *os.File
	scanner *bufio.Scanner
	text    string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// New returns new Parser
func New(path string) *Parser {
	file, err := os.Open(path)
	check(err)

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	return &Parser{file, scanner, ""}
}

// HasMoreCommands check
func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

// Advance the scanner
func (p *Parser) Advance() string {
	p.text = p.scanner.Text()

	// normalization: remove comments
	// tokens := strings.SplitN(p.text, "//", 2)
	// if len(tokens) > 0 {
	// 	p.text = tokens[0]
	// }

	return strings.TrimSpace(p.text)
}

// CommandArgs returns command args
func CommandArgs(c string) (arg1 string, arg2 string, arg3 int) {
	args := strings.Split(c, " ")
	arg1 = args[0]
	if len(args) > 1 {
		arg2 = args[1]
	}
	if len(args) == 3 {
		arg3, _ = strconv.Atoi(args[2])
	}
	return
}

// CommandType returns command type constant
func CommandType(c string) int {
	arg1, _, _ := CommandArgs(c)
	switch {
	// case arg1 == "function":
	// 	return CmdTypeFunction
	// case arg1 == "call":
	// 	return CmdTypeCall
	// case arg1 == "return":
	// 	return CmdTypeReturn
	case arg1 == "push":
		return CmdTypePush
	case arg1 == "pop":
		return CmdTypePop
	case arg1 == "eq" || arg1 == "gt" || arg1 == "lt":
		return CmdTypeComparator
	// case arg1 == "label" || arg1 == "goto" || arg1 == "if-goto":
	// 	return CmdTypeBranching
	default:
		return CmdTypeArithmetic
	}
}
