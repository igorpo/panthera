package repl

import (
	"bufio"
	"fmt"
	"io"
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
		// lx := lexer.New(line)
		// for token := lx.NextToken(); token.Type != token.EOF; token = lx.NextToken() {
		// 	fmt.Printf("%+v\n", token)
		// }
		fmt.Println(line)
	}
}
