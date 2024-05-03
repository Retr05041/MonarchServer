package parser

import (
	"MonarchServer/cmd/parser/astgen"
	"MonarchServer/cmd/parser/lexer"
	"MonarchServer/cmd/parser/tokenizer"
)

type (
	commandParseFn func() astgen.Node
)

// ##############################################################################
type Parser struct {
	l *lexer.Lexer // Instance of the lexer

	// like position and readPosition but for tokens instead of characters
	curToken  tokenizer.Token
	peekToken tokenizer.Token

	// Error handleing
	errors []string
}

// === NEW PARSER ===
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set at the beginning
	p.nextToken()
	p.nextToken()

	return p
}

// == Program Parsing ==
func (p *Parser) Parse() *astgen.AST {
	currAST := &astgen.AST{}
	currAST.Nodes = []astgen.Node{}

	// Goes through every token, creating a new node tree and adding it to the statements list
	for p.curToken.Type != tokenizer.EOF {
		stmt := p.parseTokens()
		if stmt != nil {
			currAST.Nodes = append(currAST.Nodes, stmt)
		}
		p.nextToken()
	}

	p.runCommand()
	return currAST
}

func (p *Parser) parseTokens() astgen.Node {
	switch p.curToken.Type {
	case tokenizer.PREFIX:
		return p.parsePrefix()
	case tokenizer.CMD:
		return p.parseCommand()
	default:
		return p.parseParameters() // Everything but a "let" or "return" is an expression
	}
}

// ##############################################################################
// == Specific Statement Parsing ==
// Constructs a *ast.LetStatement node with the token its currently on
func (p *Parser) parsePrefix() *astgen.Prefix {
	return &astgen.Prefix{Token: p.curToken}
}

func (p *Parser) parseCommand() *astgen.Command {
	return &astgen.Command{Token: p.curToken}
}

func (p *Parser) parseParameters() *astgen.Parameter {
    return &astgen.Parameter{Token: p.curToken}
}

func (p *Parser) runCommand() {
    return
}

// ##############################################################################
// == Utils ==

// Slide down the tokens from the lexer
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.ReadToken()
}
