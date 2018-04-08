package parser

import (
	"panthera/ast"
	"panthera/lexer"
	"panthera/token"
)

type Parser struct {
	lxr       *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(lxr *lexer.Lexer) *Parser {
	p := &Parser{lxr: lxr}

	// set curr and peek
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lxr.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
