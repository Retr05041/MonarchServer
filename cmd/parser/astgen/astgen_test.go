package astgen

import (
	"MonarchServer/cmd/parser/tokenizer"
	"testing"
)

func TestString(t *testing.T) {
	ast := &AST{
		Nodes: []Node{
            &Prefix{
                Token: tokenizer.Token{Type: tokenizer.PREFIX, Literal: "TestSource"},
                Value: "TestSource",
            },
            &Command{
                Token: tokenizer.Token{Type: tokenizer.CMD, Literal: "PASS"},
                Value: "PASS",
            },
            &Parameter{
                Token: tokenizer.Token{Type: tokenizer.PARAMETER, Literal: "supersecretpassword"},
                Value: "supersecretpassword",
            },
		},
	}

	if ast.String() != "TestSource PASS supersecretpassword" {
		t.Errorf("program.String() wrong. got=%q", ast.String())
	}
}
