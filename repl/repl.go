package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/juliaogris/monkey/evaluator"
	"github.com/juliaogris/monkey/lexer"
	"github.com/juliaogris/monkey/object"
	"github.com/juliaogris/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT) //nolint:errcheck
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect()) //nolint:errcheck
			io.WriteString(out, "\n")                //nolint:errcheck
		}
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)                                       //nolint:errcheck
	io.WriteString(out, "Woops! We ran into some monkey business here!\n") //nolint:errcheck
	io.WriteString(out, " parser errors:\n")                               //nolint:errcheck
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n") //nolint:errcheck
	}
}
