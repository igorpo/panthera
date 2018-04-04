package repl

import (
	"bufio"
	"fmt"
	"io"
	"panthera/lexer"
	"panthera/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lx := lexer.New(line)
		for tok := lx.NextToken(); tok.Type != token.EOF; tok = lx.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
