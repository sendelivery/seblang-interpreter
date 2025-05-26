package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/sendelivery/seblang-interpreter/evaluator"
	"github.com/sendelivery/seblang-interpreter/lexer"
	"github.com/sendelivery/seblang-interpreter/object"
	"github.com/sendelivery/seblang-interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			PrintParserErrors(out, p.Errors())
			continue
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const HEART = `
    ....
  *      *     ...
 *        *   *   *
*         * *      *
 *         *       *
  *              *
   *            *
    *         *
      *     *
       *  *
        *
		
`

func PrintParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, HEART)
	io.WriteString(out, "Woops! We ran into a problem here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
