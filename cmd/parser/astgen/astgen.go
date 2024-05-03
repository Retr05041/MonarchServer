package astgen

import (
	"MonarchServer/cmd/parser/tokenizer"
	"bytes"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Prefix struct {
	Token tokenizer.Token
	Value string
}

func (p *Prefix) TokenLiteral() string { return p.Token.Literal }
func (p *Prefix) String() string       { return p.Value }

type Command struct {
	Token  tokenizer.Token
	Value  string
}

func (c *Command) TokenLiteral() string { return c.Token.Literal }
func (c *Command) String() string {
	var out bytes.Buffer
	out.WriteString(" " + c.Value)
	return out.String()
}

type Parameter struct {
    Token tokenizer.Token
    Value string
}

func (par *Parameter) TokenLiteral() string { return par.Token.Literal }
func (par *Parameter) String() string {
    var out bytes.Buffer
    out.WriteString(" " + par.Value)
    return out.String()
}

// Root Node for every AST
type AST struct {
	Nodes []Node
}

func (a *AST) TokenLiteral() string {
	if len(a.Nodes) > 0 {
		return a.Nodes[0].TokenLiteral()
	} else {
		return ""
	}
}

func (a *AST) String() string {
	var out bytes.Buffer

	for _, s := range a.Nodes{
		out.WriteString(s.String())
	}

	return out.String()
}
