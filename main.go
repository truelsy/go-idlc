// Main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const (
	TAB2 = "\t\t"
	TAB3 = "\t\t\t"
	TAB4 = "\t\t\t\t"
)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func isEof(ch rune) bool {
	return (ch == eof)
}

func isOpenBrace(ch rune) bool {
	return (ch == '{')
}

func isEndBrace(ch rune) bool {
	return (ch == '}')
}

func isColon(ch rune) bool {
	return (ch == ':')
}

func isOpenBracket(ch rune) bool {
	return (ch == '[')
}

func isPrimitiveType(a string) bool {
	return (a == "string" || a == "int8" || a == "uint8" || a == "int16" || a == "uint16" || a == "int32" || a == "uint32" || a == "int64" || a == "uint64")
}

func isArrayType(a string) bool {
	m, _ := regexp.MatchString("\\[\\].+", a)
	return m
}

type TokenStmt struct {
	Elements []*TokenElement
}

type TokenElement struct {
	VarName []string
	VarType []string
	IsMsg   bool
	Name    string
	Id      int
}

func main() {
	lang := flag.String("l", "go", "target language [go|cs|cpp]")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stdout, "usage: %v -l=[go|cs|cpp] input\n", os.Args[0])
		return
	}

	switch *lang {
	case "go", "cs", "cpp": // do nothing
	default:
		fmt.Fprintf(os.Stdout, "usage: %v -l=[go|cs|cpp] input\n", os.Args[0])
		return
	}

	input := flag.Args()[0]
	rf, err := os.Open(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "input file open error(", err, ")")
		return
	}

	dir := filepath.Dir(input)
	output := filepath.Base(input)
	ext := filepath.Ext(output)
	fname := output[0 : len(output)-len(ext)]

	switch *lang {
	case "go":
		output = filepath.Join(dir, fname+".go")
	case "cs":
		output = filepath.Join(dir, fname+".cs")
	case "cpp":
		output = filepath.Join(dir, fname+".hpp")
	}

	fmt.Fprintf(os.Stderr, "compile (%v -> %v)\n", input, output)

	stmt, err := NewParser(bufio.NewReader(rf)).Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, "input file parsing error(", err, ")")
		return
	}

	wf, err := os.Create(output)
	if err != nil {
		fmt.Fprintln(os.Stderr, "output file create error(", err, ")")
		return
	}

	//	// redirect
	os.Stdout = wf

	fmt.Fprintln(os.Stdout, "//////////////////////////////////////////////////////////////////")
	fmt.Fprintln(os.Stdout, "// Automatically-generated file. Do not edit!")
	fmt.Fprintln(os.Stdout, "//////////////////////////////////////////////////////////////////\n")

	switch *lang {
	case "go":
		CompileGoCode(stmt)
	case "cs":
		CompileCsCode(stmt)
	case "cpp":
		CompileCppCode(stmt, fname)
	}

	rf.Close()
	wf.Close()
}
