package codewriter

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CodeWriter type
type CodeWriter struct {
	writer *bufio.Writer
}

// New creates new CodeWriter
func New(filePath string) *CodeWriter {
	hackfile := strings.TrimSuffix(filePath, ".asm") + ".hack"
	writerfp, err := os.OpenFile(hackfile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Err! .hack file couldn't create.")
		os.Exit(4)
	}
	writer := bufio.NewWriter(writerfp)

	return &CodeWriter{writer}
}

// WriteA translates a-instruction value to binary code
func (cw *CodeWriter) WriteA(v string) {
	addrInt, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(err)
	}

	s := fmt.Sprintf("%015b", addrInt) // int to binary string, padd with 0
	hack := "0" + s[len(s)-15:]        // trim and prefix with 0

	fmt.Fprintln(cw.writer, hack)
	cw.writer.Flush()
}

// WriteC translates c-instruction value to binary code
func (cw *CodeWriter) WriteC(d string, c string, j string) {
	if len(d) == 0 {
		d = "null"
	}
	if len(j) == 0 {
		j = "null"
	}

	hack := fmt.Sprintf("111%s%s%s", computations[c], destinations[d], jumps[j])
	fmt.Fprintln(cw.writer, hack)
	cw.writer.Flush()
}
