package ast

import (
	"github.com/tomanta/interpreteringo/token"
)

type Node interface {
	TokenLiteral() string
}

// Statements do not return a value
type Statement interface {
	Node
	statementNode()
}

// Expressions return a value
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST the parser produces
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // What to assign TO
	Value Expression  // The expression to assign
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
