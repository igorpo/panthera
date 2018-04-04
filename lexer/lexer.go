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
	l.consumeWhitespace()
	var tok token.Token

	switch l.ch {
	case '=':
		if l.peek() == '=' {
			l.nextChar()
			tok = newToken(token.EQ, "="+string(l.ch))
		} else {
			tok = newToken(token.ASSIGN, string(l.ch))
		}
	case '+':
		tok = newToken(token.PLUS, string(l.ch))
	case '-':
		tok = newToken(token.MINUS, string(l.ch))
	case '!':
		if l.peek() == '=' {
			l.nextChar()
			tok = newToken(token.NEQ, "!"+string(l.ch))
		} else {
			tok = newToken(token.BANG, string(l.ch))
		}
	case '/':
		tok = newToken(token.SLASH, string(l.ch))
	case '*':
		tok = newToken(token.ASTERISK, string(l.ch))
	case '<':
		tok = newToken(token.LT, string(l.ch))
	case '>':
		tok = newToken(token.GT, string(l.ch))
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	case '(':
		tok = newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = newToken(token.RPAREN, string(l.ch))
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case '{':
		tok = newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = newToken(token.RBRACE, string(l.ch))
	case 0:
		tok = newToken(token.EOF, "")
	default:
		if isLetter(l.ch) {
			ident := l.parseIdentifier()

			// we do not need another call to nextChar() because of how our parseIdentifier() works
			return newToken(token.LookupIdentifier(ident), ident)
		} else if isNumber(l.ch) {
			return newToken(token.INT, l.parseNumber())
		} else {
			tok = newToken(token.ILLEGAL, string(l.ch))
		}
	}

	// update position in input
	l.nextChar()
	return tok
}

func (l *Lexer) parseIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.nextChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) parseNumber() string {
	start := l.position
	for isNumber(l.ch) {
		l.nextChar()
	}
	return l.input[start:l.position]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: ch,
	}
}

func (l *Lexer) consumeWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.nextChar()
	}
}

func (l *Lexer) peek() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}
