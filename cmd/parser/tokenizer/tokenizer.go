package tokenizer

type TokenType string 

type Token struct {
	Type    TokenType // type of token
	Literal string    // the token itself
}

const (
	COLON  = ":"
    EOF = "EOF"
    PARAMETER = "PARAMETER"
    ILLEGAL = "ILLEGAL"
    CMD = "CMD"
    PREFIX = "PREFIX"
)
