package ast

import (
	"bytes"
	"panthera/token"
)

type Node interface {
	PrintToken() string // debugging
	String() string
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

// since we can have code in the form
//     $ let x = 5;
//     $ x + 10;
// we need an expression statement to parse the latter line
type ExpressionStatement struct {
	Token      token.Token // first token of expr
	Expression Expression
}

func (p *Program) PrintToken() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].PrintToken()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) PrintToken() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.PrintToken() + " ")
	out.WriteString(ls.Name.Value)
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) PrintToken() string {
	return i.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) PrintToken() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.PrintToken() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) PrintToken() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
