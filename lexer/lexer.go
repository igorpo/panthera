package lexer

import (
	"panthera/token"
)

type Lexer struct {
	input        string
	position     int  // points to current char
	nextPosition int  // points to next char
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	lxr := &Lexer{input: input}
	lxr.nextChar() // start the lexer with the first char of input
	return lxr
}

// update the next position in the input
func (l *Lexer) nextChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0 // ASCII nul
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, string(l.ch))
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	case '(':
		tok = newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = newToken(token.RPAREN, string(l.ch))
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case '+':
		tok = newToken(token.PLUS, string(l.ch))
	case '{':
		tok = newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = newToken(token.RBRACE, string(l.ch))
	case 0:
		tok = newToken(token.EOF, "")
	}

	// update position in input
	l.nextChar()
	return tok
}

func newToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: ch,
	}
}
