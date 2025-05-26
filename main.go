package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/sendelivery/seblang-interpreter/evaluator"
	"github.com/sendelivery/seblang-interpreter/lexer"
	"github.com/sendelivery/seblang-interpreter/object"
	"github.com/sendelivery/seblang-interpreter/parser"
	"github.com/sendelivery/seblang-interpreter/repl"
)

func main() {
	if len(os.Args) == 2 {
		os.Exit(runProgram(os.Args[1]))
	} else if len(os.Args) > 2 {
		fmt.Println("usage: seblang <optional filepath>")
		os.Exit(1)
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Seb programming language!\n",
		user.Name)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func runProgram(filepath string) int {
	b, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	l := lexer.New(string(b))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		repl.PrintParserErrors(os.Stdout, p.Errors())
		return 1
	}

	macroEnv := object.NewEnvironment()
	evaluator.DefineMacros(program, macroEnv)
	expanded := evaluator.ExpandMacros(program, macroEnv)

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(expanded, env)
	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	}

	return 0
}
