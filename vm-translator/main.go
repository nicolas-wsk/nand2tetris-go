package main

import (
	"os"

	"github.com/nicolas-wsk/nand2tetris-go/vm-translator/codewriter"
	"github.com/nicolas-wsk/nand2tetris-go/vm-translator/parser"
)

func main() {
	filePath := os.Args[1]

	p := parser.New(filePath)
	cw := codewriter.New(filePath)

	for p.HasMoreCommands() {
		c := p.Advance()
		arg1, arg2, arg3 := parser.CommandArgs(c)

		switch parser.CommandType(c) {
		case parser.CmdTypeArithmetic:
			cw.WriteArithmetic(arg1)
		case parser.CmdTypePush:
			cw.WritePush(arg2, arg3)
		case parser.CmdTypePop:
			cw.WritePop(arg2, arg3)
		case parser.CmdTypeComparator:
			cw.WriteComparator(arg1)
			// case parser.CmdTypeComparator:
			// 	cw.WriteComparator(arg1)
			// case parser.CmdTypeBranching:
			// 	cw.WriteBranching(arg1, arg2)
			// case parser.CmdTypeFunction:
			// 	cw.WriteFunction(arg2, arg3)
			// case parser.CmdTypeCall:
			// 	cw.WriteCall(arg2, arg3)
			// case parser.CmdTypeReturn:
			// 	cw.WriteReturn()
		}
	}
}
