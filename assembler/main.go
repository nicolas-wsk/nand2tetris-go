package main

import (
	"os"
	"strconv"

	"github.com/nicolas-wsk/nand2tetris-go/assembler/codewriter"
	"github.com/nicolas-wsk/nand2tetris-go/assembler/parser"
	"github.com/nicolas-wsk/nand2tetris-go/assembler/symboltable"
)

func main() {
	filePath := os.Args[1]
	var instructions []string

	p := parser.New(filePath)
	w := codewriter.New(filePath)
	s := symboltable.New()

	// 1st : convert the Loop symbols and collect instructions
	for romAddr := 0; p.HasMoreCommands(); {
		instruction := p.Advance()
		a, _, _ := p.CommandArgs(instruction)

		switch p.CommandType(instruction) {
		case parser.A_COMMAND, parser.C_COMMAND:
			romAddr++
			instructions = append(instructions, instruction)
		case parser.L_COMMAND:
			s[a] = romAddr
		}

	}

	// 2nd: Translate all the instructions previously collected
	romAddr := 16
	for _, instruction := range instructions {
		a, b, c := p.CommandArgs(instruction)

		switch p.CommandType(instruction) {
		case parser.A_COMMAND:
			if parser.IsVariable(a) {
				// if variable, use from table
				_, found := s[a]
				if !found { // if not declared yet, put to table
					s[a] = romAddr
					romAddr++
				}
				w.WriteA(strconv.Itoa(s[a]))
			} else {
				w.WriteA(a)
			}
		case parser.C_COMMAND:
			w.WriteC(a, b, c)
		}
	}

}
