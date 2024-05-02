package token

// type: string - ease of tokenization and ease of debugging - better perfomance would be int or byte
type TokenType string

type Token struct {
	Type    TokenType // type of token
	Literal string    // the token itself
}

const (
	ILLEGAL = "ILLEGAL" // tokens we don't know about
	EOF     = "EOF"     // stop at end of file

	// Keywords
	COMMAND   = "COMMAND"
	PREFIX    = "PREFIX"
	PARAMETER = "PARAMETER"

	// Operators
	COLON = ":"
)

// Table for mapping strings to token types
var commands = map[string]TokenType{
	"PASS": COMMAND,
	"NICK": COMMAND,
	"USER": COMMAND,
}

// takes a string to test and assigns it to a tokentype or returns a uesr-defined identifier
func LookupCmd(cmd string) TokenType {
	if tok, ok := commands[cmd]; ok {
		return tok
	}
	return COMMAND
}
