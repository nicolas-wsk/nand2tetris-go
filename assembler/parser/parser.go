package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Command type
type Command int

const (
	N_COMMAND Command = 0 // empty command
	A_COMMAND Command = 1 // address ex)@123
	C_COMMAND Command = 2 // compute ex)A=M;JGT
	L_COMMAND Command = 3 // label   ex)(LOOP)
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
	tokens := strings.SplitN(p.text, "//", 2)
	if len(tokens) > 0 {
		p.text = tokens[0]
	}

	return strings.TrimSpace(p.text)
}

// CommandType getter
func (p *Parser) CommandType(s string) Command {
	if len(s) == 0 {
		// empty token
		return N_COMMAND
	}
	if s[0] == '@' {
		// @Xxx
		return A_COMMAND
	}
	if s[0] == '(' {
		// (LABEL)
		return L_COMMAND
	}

	return C_COMMAND
}

// CommandArgs returns command args
// (dest, comp, jump) in case of c-instruction
func (p *Parser) CommandArgs(s string) (a string, b string, c string) {
	a, b, c = "", "", ""

	switch p.CommandType(s) {
	case L_COMMAND:
		a = s[1 : len(s)-1]
	case A_COMMAND:
		a = s[1:]
	case C_COMMAND:
		splitEqual := strings.SplitN(s, "=", 2)
		splitComma := strings.SplitN(s, ";", 2)

		if len(splitEqual) > 1 {
			a = splitEqual[0]
			b = splitEqual[1]
		}

		if len(splitComma) > 1 {
			b = splitComma[0]
			c = splitComma[1]
		}

	}
	return
}

// IsVariable tests if string is variable or value
// returns true if variable
func IsVariable(s string) bool {
	_, err := strconv.ParseInt(s, 10, 16)
	return err != nil
}
