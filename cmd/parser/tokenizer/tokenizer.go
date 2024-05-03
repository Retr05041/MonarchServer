package tokenizer

type TokenType string 

type Token struct {
	Type    TokenType // type of token
	Literal string    // the token itself
}

const (
	COLON  = ":"
	BANG   = "!"
	AT     = "@"
	STAR   = "*"
	DASH   = "-"
	EQUALS = "="
    EOF = "EOF"
    WORD = "WORD"
    ILLEGAL = "ILLEGAL"
)
