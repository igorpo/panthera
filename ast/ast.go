package ast

import (
	"panthera/token"
)

type Node interface {
	PrintToken() string // debugging
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// root node of every AST our parser produces
type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // LET token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

type ReturnStatement struct {
	Token       token.Token // RETURN token
	ReturnValue Expression
}

func (p *Program) PrintToken() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].PrintToken()
	}
	return ""
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) PrintToken() string {
	return ls.Token.Literal
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) PrintToken() string {
	return i.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) PrintToken() string {
	return rs.Token.Literal
}
