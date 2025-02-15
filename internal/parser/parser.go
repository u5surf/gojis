package parser

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

//go:generate antlr -Dlanguage=Go -visitor ECMAScript.g4

type Parser struct {
	ast interface{}
}

func New() *Parser {
	p := new(Parser)
	// p.ast = NewEmptyAst()
	return p
}

func (p *Parser) ParseFiles(paths ...string) (errs []error) {
	for _, path := range paths {
		err := p.ParseFile(path)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return
}

func (p *Parser) ParseFile(path string) error {
	input, err := antlr.NewFileStream(path)
	if err != nil {
		return fmt.Errorf("Error while loading file: %v", err)
	}

	_ = input

	// lexer := NewECMAScriptLexer(input)
	// lexer.RemoveErrorListeners()

	// stream := antlr.NewCommonTokenStream(lexer, 0)
	// par := NewECMAScriptParser(stream)
	// par.RemoveErrorListeners()

	// errorCollector := NewCollectingErrorListener()
	// par.AddErrorListener(errorCollector)
	// par.BuildParseTrees = true

	// tree := par.Program()
	// if errs, hasErrors := errorCollector.Errors(); hasErrors {
	// 	return NewParserError(path, errs...)
	// }

	// // only append root if no errors occurred while parsing
	// p.ast.AddRoot(path, tree)

	panic("TODO")
}

func (p *Parser) Ast() interface{} {
	return p.ast
}

type ParserError struct {
	file string
	errs []error
}

func NewParserError(file string, errs ...error) ParserError {
	return ParserError{
		file: file,
		errs: errs,
	}
}

func (e ParserError) Error() string {
	if len(e.errs) == 0 {
		return ""
	}

	var buf bytes.Buffer
	_, err := buf.WriteString(fmt.Sprintf("Errors while parsing '%v':", e.file))
	if err != nil {
		// this cannot happen as we use a bytes.Buffer, which never returns an error
		panic(err)
	}

	for i, err := range e.errs {
		_, err := buf.WriteString("\n\t" + strconv.Itoa(i+1) + ") " + err.Error())
		if err != nil {
			// this cannot happen as we use a bytes.Buffer, which never returns an error
			panic(err)
		}
	}

	return buf.String()
}
